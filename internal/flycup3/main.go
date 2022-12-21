package main

import (
	"fmt"
	"os"

	"github.com/compagnon/go-clients/pkg/googlecloud/csclient"
)

const GOOGLE_CLOUD_PROJECT = "qwiklabs-gcp-03-33563a0ccaec"

// et
// * Enable BQ API
// https://console.cloud.google.com/flows/enableapi?apiid=bigquery&_ga=2.97460124.2018023344.1668447428-1515625999.1614013342
// * Create a service account:
// https://console.cloud.google.com/projectselector/iam-admin/serviceaccounts/create?supportedpurview=project&_ga=2.97460124.2018023344.1668447428-1515625999.1614013342

// CREDENTIALs for the CLient:
//%APPDATA%\gcloud\application_default_credentials.json
// or set GOOGLE_APPLICATION_CREDENTIALS = ....json

func main() {

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "../credentials/application_credentials_flycup.json")
	os.Setenv("GOOGLE_CLOUD_PROJECT", GOOGLE_CLOUD_PROJECT)

	fmt.Println(os.LookupEnv("GOOGLE_APPLICATION_CREDENTIALS"))

	TEMP_ROOT_DIR := "C:/Users/TEMP/"
	TEMP_PATH := "gsp397/myracermodel/"
	// GOOGLE_DATASET := "drl"
	// GOOGLE_REGION := "US"
	// var csvdatafiles []fs.FileInfo
	// var err error
	// var sql, city, event, pilot string
	// //var scanner *bufio.Scanner
	// var line []byte

	csclient.UploadDirectory(os.Stdout, GOOGLE_CLOUD_PROJECT+string("-bucket"), nil, TEMP_ROOT_DIR+TEMP_PATH, "model")
	// gs: //$PROJECT_ID-bucket/model
}
