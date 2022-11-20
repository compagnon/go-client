package csclient

import (
	"bytes"
	"reflect"
	"testing"

	"cloud.google.com/go/storage"
)

func Test_createBucket(t *testing.T) {
	type args struct {
		projectID  string
		bucketName string
		location   string
	}
	tests := []struct {
		name    string
		args    args
		want    *storage.BucketHandle
		wantErr bool
	}{
		// TODO: Add test cases.
		{"simpleTest",
			args{
				projectID:  "xxxxxxxxxxxxxxxxx",
				bucketName: "20221120-testgo-clients",
				location:   "US",
			},
			nil,
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateBucket(tt.args.projectID, tt.args.bucketName, tt.args.location)
			if (err != nil) != tt.wantErr {
				t.Errorf("createBucket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createBucket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_downloadFile(t *testing.T) {
	type args struct {
		bucketName   string
		bucket       *storage.BucketHandle
		object       string
		destFileName string
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"download Csv",
			args{bucketName: "spls",
				bucket:       nil,
				object:       "/gsp394/tables/*.csv",
				destFileName: "C:/Users/TEMP",
			},
			"Blob /gsp394/tables/*.csv downloaded to local file C:/Users/TEMP\n",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := DownloadFile(w, tt.args.bucketName, tt.args.bucket, tt.args.object, tt.args.destFileName); (err != nil) != tt.wantErr {
				t.Errorf("downloadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("downloadFile() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func Test_listFilesWithPrefix(t *testing.T) {
	type args struct {
		bucketName string
		prefix     string
		delim      string
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"list Csv in a bucket",
			args{bucketName: "spls",
				prefix: "/gsp394/tables",
				delim:  ",",
			},
			"Blob /gsp394/tables/*.csv downloaded to local file C:/Users/TEMP\n",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := ListFilesWithPrefix(w, tt.args.bucketName, tt.args.prefix, tt.args.delim, 1); (err != nil) != tt.wantErr {
				t.Errorf("listFilesWithPrefix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("listFilesWithPrefix() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
