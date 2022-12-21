package main

import (
	"fmt"
	"os"

	"github.com/compagnon/go-clients/pkg/googlecloud/csclient"
)

const GOOGLE_CLOUD_PROJECT = "idyllic-pact-205716"

func main() {
	//https://storage.googleapis.com/gcp-public-data-landsat/LC08/01/001/003/LC08_L1GT_001003_20140812_20170420_01_T2/LC08_L1GT_001003_20140812_20170420_01_T2_B3.TIF

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./internal/credentials/application_credentials_jyp.json")
	os.Setenv("GOOGLE_CLOUD_PROJECT", GOOGLE_CLOUD_PROJECT)

	fmt.Println(os.LookupEnv("GOOGLE_APPLICATION_CREDENTIALS"))
	//csclient.ListFiles(os.Stdout, "gcp-public-data-landsat", 1)
	//csclient.ListFilesWithPrefix(os.Stdout, "gcp-public-data-landsat", "LC08/01/001/003/LC08_L1GT_001003_20140812_20170420_01_T2/", "", 1)
	csclient.ListFilesWithPrefix(os.Stdout, "spls", "gsp394/tables", "", 10)

	// download files from a google cloud storage
	//csclient.DownloadFile(os.Stdout, "spls", nil, "gsp394/tables/time_trial_groups.csv", "C:/Users/TEMP/time_trial_groups.csv")
	//csclient.DownloadDirectory(os.Stdout, "spls", nil, "gsp394/tables/", "C:/Users/TEMP/")

	csclient.DownloadDirectory(os.Stdout, "spls", nil, "gsp397/", "C:/Users/TEMP/")
}
