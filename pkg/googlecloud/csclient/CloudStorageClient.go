package csclient

/*
code from https://github.com/GoogleCloudPlatform/golang-samples/tree/main/storage
*/

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func init() {
	fmt.Println("Init csclient")
}

// createBucket create a new bucket with given name
func CreateBucket(projectID string, bucketName string, location string) (*storage.BucketHandle, error) {
	ctx := context.Background()

	// Creates a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil, err
	}
	defer client.Close()

	// Creates a Bucket instance.
	bucket := client.Bucket(bucketName)

	// Creates the new bucket.
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	if err := bucket.Create(ctx, projectID, nil); err != nil {
		log.Fatalf("Failed to create bucket: %v", err)
		return nil, err
	}

	return bucket, nil
}

// downloadFile downloads an object to a file.
func DownloadFile(w io.Writer, bucketName string, bucket *storage.BucketHandle, object string, destFileName string) error {
	// bucket := "bucket-name"
	// object := "object-name"
	// destFileName := "file.txt"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("storage.NewClient: %v", err)
		return err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	f, err := os.Create(destFileName)
	if err != nil {
		log.Fatalf("os.Create: %v", err)
		return err
	}

	if bucket == nil {
		bucket = client.Bucket(bucketName)
	}

	rc, err := bucket.Object(object).NewReader(ctx)
	if err != nil {
		log.Fatalf("Object(%q).NewReader: %v", object, err)
		return err
	}
	defer rc.Close()

	if _, err := io.Copy(f, rc); err != nil {
		log.Fatalf("io.Copy: %v", err)
		return err
	}

	if err = f.Close(); err != nil {
		log.Fatalf("f.Close: %v", err)
		return err
	}

	fmt.Fprintf(w, "Blob %v downloaded to local file %v\n", object, destFileName)

	return nil
}

// downloadFileIntoMemory downloads an object.
func DownloadFileIntoMemory(w io.Writer, bucketName, object string) ([]byte, error) {
	// bucket := "bucket-name"
	// object := "object-name"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	rc, err := client.Bucket(bucketName).Object(object).NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("Object(%q).NewReader: %v", object, err)
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %v", err)
	}
	fmt.Fprintf(w, "Blob %v downloaded.\n", object)
	return data, nil
}

// listFiles lists objects within specified bucket.
func ListFiles(w io.Writer, bucketName string, timeOut time.Duration) error {
	// bucket := "bucket-name"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	it := client.Bucket(bucketName).Objects(ctx, nil)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("Bucket(%q).Objects: %v", bucketName, err)
		}
		fmt.Fprintln(w, attrs.Name)
	}
	return nil
}

// listFilesWithPrefix lists objects using prefix and delimeter.
func ListFilesWithPrefix(w io.Writer, bucketName, prefix, delim string, timeOut time.Duration) error {
	// bucket := "bucket-name"
	// prefix := "/foo"
	// delim := "_"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	// Prefixes and delimiters can be used to emulate directory listings.
	// Prefixes can be used to filter objects starting with prefix.
	// The delimiter argument can be used to restrict the results to only the
	// objects in the given "directory". Without the delimiter, the entire tree
	// under the prefix is returned.
	//
	// For example, given these blobs:
	//   /a/1.txt
	//   /a/b/2.txt
	//
	// If you just specify prefix="a/", you'll get back:
	//   /a/1.txt
	//   /a/b/2.txt
	//
	// However, if you specify prefix="a/" and delim="/", you'll get back:
	//   /a/1.txt
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeOut))
	defer cancel()

	it := client.Bucket(bucketName).Objects(ctx, &storage.Query{
		Prefix:    prefix,
		Delimiter: delim,
	})
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("Bucket(%q).Objects(): %v", bucketName, err)
		}
		fmt.Fprintln(w, attrs.Name)
	}
	return nil
}

// downloadFile downloads an object to a file.
func DownloadDirectory(w io.Writer, bucketName string, bucket *storage.BucketHandle, object string, destDirName string) error {

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("storage.NewClient: %v", err)
		return err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	if bucket == nil {
		bucket = client.Bucket(bucketName)
	}
	it := bucket.Objects(ctx, &storage.Query{
		Prefix: object,
	})
	for {
		obj, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("Bucket(%q).Objects(): %v", bucketName, err)
		}
		if obj.Name[len(obj.Name)-1:] == "/" {
			continue
		}

		//one by one file
		f, err := create(destDirName + obj.Name)
		if err != nil {
			log.Fatalf("os.Create: %v", err)
			return err
		}

		rc, err := bucket.Object(obj.Name).NewReader(ctx)
		if err != nil {
			log.Fatalf("Object(%q).NewReader: %v", obj.Name, err)
			return err
		}
		defer rc.Close()

		if _, err := io.Copy(f, rc); err != nil {
			log.Fatalf("io.Copy: %v", err)
			return err
		}

		if err = f.Close(); err != nil {
			log.Fatalf("f.Close: %v", err)
			return err
		}

	}

	fmt.Fprintf(w, "Blob %v downloaded to local file %v\n", object, destDirName)

	return nil
}

// uploadFile uploads an object.
func UploadFile(w io.Writer, bucketName string, bucket *storage.BucketHandle, fromFileName string, object string) error {
	// bucket := "bucket-name"
	// object := "object-name"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	if bucket == nil {
		bucket = client.Bucket(bucketName)
	}

	err = uploadFile(w, ctx, bucket, fromFileName, object)
	return err
}

// uploadFile uploads an object.
func UploadDirectory(w io.Writer, bucketName string, bucket *storage.BucketHandle, fromDirName string, object string) error {

	// bucket := "bucket-name"
	// object := "object-name"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	if bucket == nil {
		bucket = client.Bucket(bucketName)
	}

	return uploadDir(w, ctx, bucket, fromDirName, object)
}
func uploadDir(w io.Writer, ctx context.Context, bucket *storage.BucketHandle, fromDirName string, object string) error {

	files, err := ioutil.ReadDir(fromDirName)
	if err != nil {
		log.Fatalf("os.readDir: %v", err)
		return err
	}

	for _, fi := range files {

		if fi.IsDir() {
			// recursive
			err = uploadDir(w, ctx, bucket, fromDirName+string(os.PathSeparator)+fi.Name(), object+string(os.PathSeparator)+fi.Name())
		} else { // file to upload
			// Upload an object with storage.Writer.
			err = uploadFile(w, ctx, bucket, fromDirName+string(os.PathSeparator)+fi.Name(), object+string(os.PathSeparator)+fi.Name())
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func uploadFile(w io.Writer, ctx context.Context, bucket *storage.BucketHandle, fromFileName string, object string) error {
	// Open local file.
	f, err := os.Open(fromFileName)
	if err != nil {
		return fmt.Errorf("os.Open: %v", err)
	}
	defer f.Close()

	o := bucket.Object(object)

	// Optional: set a generation-match precondition to avoid potential race
	// conditions and data corruptions. The request to upload is aborted if the
	// object's generation number does not match your precondition.
	// For an object that does not yet exist, set the DoesNotExist precondition.
	o = o.If(storage.Conditions{DoesNotExist: true})
	// If the live object already exists in your bucket, set instead a
	// generation-match precondition using the live object's generation number.
	// attrs, err := o.Attrs(ctx)
	// if err != nil {
	// 	return fmt.Errorf("object.Attrs: %v", err)
	// }
	// o = o.If(storage.Conditions{GenerationMatch: attrs.Generation})

	// Upload an object with storage.Writer.
	wc := o.NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}
	fmt.Fprintf(w, "Blob %v uploaded from local file %v.\n", object, fromFileName)
	return nil
}

func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}
