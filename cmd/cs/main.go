package main

import (
	"os"

	"github.com/compagnon/go-clients/pkg/googlecloud/csclient"
)

const GOOGLE_CLOUD_PROJECT = "idyllic-pact-205716"

func main() {
	//https://storage.googleapis.com/gcp-public-data-landsat/LC08/01/001/003/LC08_L1GT_001003_20140812_20170420_01_T2/LC08_L1GT_001003_20140812_20170420_01_T2_B3.TIF

	//csclient.ListFiles(os.Stdout, "gcp-public-data-landsat", 1)
	//csclient.ListFilesWithPrefix(os.Stdout, "gcp-public-data-landsat", "LC08/01/001/003/LC08_L1GT_001003_20140812_20170420_01_T2/", "", 1)
	csclient.ListFilesWithPrefix(os.Stdout, "spls", "gsp394/tables", "", 10)
	csclient.DownloadFile(os.Stdout, "spls", nil, "gsp394/tables/time_trial_groups.csv", "C:/Users/TEMP/time_trial_groups.csv")
	csclient.DownloadDirectory(os.Stdout, "spls", nil, "gsp394/tables/", "C:/Users/TEMP/")
}
