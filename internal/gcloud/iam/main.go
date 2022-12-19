package main

// code from https://cloud.google.com/iam/docs/creating-managing-service-account-keys#iam-service-account-keys-create-go
import (
	"context"
	"os"

	// "encoding/base64"
	"fmt"
	"io"

	iam "google.golang.org/api/iam/v1"
)

func main() {

	SA_NAME := "golang2"
	PROJECT_ID := "idyllic-pact-205716"
	var err error
	var sakey *iam.ServiceAccountKey
	var sakeys []*iam.ServiceAccountKey

	sakey, err = createKey(os.Stdout, SA_NAME+"@"+PROJECT_ID+".iam.gserviceaccount.com")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sakey.KeyType)
	}

	sakeys, err = listKeys(os.Stdout, SA_NAME+"@"+PROJECT_ID+".iam.gserviceaccount.com")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(len(sakeys))
	}
}

// createKey creates a service account key.
func createKey(w io.Writer, serviceAccountEmail string) (*iam.ServiceAccountKey, error) {
	ctx := context.Background()
	service, err := iam.NewService(ctx)
	if err != nil {
		return nil, fmt.Errorf("iam.NewService: %v", err)
	}

	resource := "projects/-/serviceAccounts/" + serviceAccountEmail
	request := &iam.CreateServiceAccountKeyRequest{}
	key, err := service.Projects.ServiceAccounts.Keys.Create(resource, request).Do()
	if err != nil {
		return nil, fmt.Errorf("Projects.ServiceAccounts.Keys.Create: %v", err)
	}
	// The PrivateKeyData field contains the base64-encoded service account key
	// in JSON format.
	// TODO(Developer): Save the below key (jsonKeyFile) to a secure location.
	// You cannot download it later.
	// jsonKeyFile, _ := base64.StdEncoding.DecodeString(key.PrivateKeyData)
	fmt.Fprintf(w, "Key created successfully")
	return key, nil
}

// listKey lists a service account's keys.
func listKeys(w io.Writer, serviceAccountEmail string) ([]*iam.ServiceAccountKey, error) {
	ctx := context.Background()
	service, err := iam.NewService(ctx)
	if err != nil {
		return nil, fmt.Errorf("iam.NewService: %v", err)
	}

	resource := "projects/-/serviceAccounts/" + serviceAccountEmail
	response, err := service.Projects.ServiceAccounts.Keys.List(resource).Do()
	if err != nil {
		return nil, fmt.Errorf("Projects.ServiceAccounts.Keys.List: %v", err)
	}
	for _, key := range response.Keys {
		fmt.Fprintf(w, "Listing key: %v", key.Name)
	}
	return response.Keys, nil
}
