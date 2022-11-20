# Golang Clients

 for Bot , Automation , Stress / unit test, MicroService Achitecture  

some code snippets , research about how to automatize and perf /integration  tests 
keywords:  https rest graphql E2Etest

readings and ramblings

- https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091 (No or less global variables)
- http://peter.bourgon.org/go-kit/
- [How to write Go code](https://golang.org/doc/code.html)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Wiki](https://github.com/golang/go/wiki)
- [Using Go Modules](https://go.dev/blog/using-go-modules)
- [project layout Overkill mustread](https://github.com/golang-standards/project-layout)

- https://goreportcard.com/

## GoogleCloud

Collection of automation with GCP
###  BigQuery

please refer to
  https://cloud.google.com/bigquery/docs/reference/libraries#client-libraries-usage-go
  
#### service account and JSON key:
https://cloud.google.com/bigquery/docs/quickstarts/quickstart-client-libraries

* Enable BQ API
https://console.cloud.google.com/flows/enableapi?apiid=bigquery&_ga=2.97460124.2018023344.1668447428-1515625999.1614013342
* Create a service account:
https://console.cloud.google.com/projectselector/iam-admin/serviceaccounts/create?supportedpurview=project&_ga=2.97460124.2018023344.1668447428-1515625999.1614013342


CREDENTIALs for the CLient:


$env:GOOGLE_APPLICATION_CREDENTIALS="C:\Users\username\Downloads\service-account-file.json"
set GOOGLE_APPLICATION_CREDENTIALS=KEY_PATH

or
%APPDATA%\gcloud\application_default_credentials.json



go mod init YOUR_MODULE_NAME

go get cloud.google.com/go/bigquery
go get google.golang.org/api/iterator