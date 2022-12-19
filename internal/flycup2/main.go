package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/compagnon/go-clients/pkg/googlecloud/bqclient"
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

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "../credentials/application_default_credentials.json")
	os.Setenv("GOOGLE_CLOUD_PROJECT", GOOGLE_CLOUD_PROJECT)

	fmt.Println(os.LookupEnv("GOOGLE_APPLICATION_CREDENTIALS"))

	TEMP_ROOT_DIR := "C:/Users/TEMP/"
	TEMP_PATH := "gsp394/tables/"
	GOOGLE_DATASET := "drl"
	GOOGLE_REGION := "US"
	var csvdatafiles []fs.FileInfo
	var err error
	var sql, city, event, pilot string
	//var scanner *bufio.Scanner
	var line []byte

	//MENU
MENU:
	//Google Fly Cup Challenge: Pilot
	fmt.Printf("MENU:\n*Task 1: Create Google Cloud Storage Bucket \n*Task 2: Import DRL Pilot Data to BigQuery\n*Task 3: Download Starter Code to Vertex AI Notebook\n")
	fmt.Printf("*Task 4: Train a Model\n*Task 5: Upload Model into Vertex AI Model Registry\n*Task 6: Deploy Model to an Endpoint\n")
	fmt.Printf("* Task 7: Query your Endpoint \n")
	reader := bufio.NewReader(os.Stdin)
	//car, _ := reader.ReadByte()1
	line, _, _ = reader.ReadLine()

	switch string(line) {
	case "0":
		goto TASK0
	case "1":
		goto TASK1
	case "2":
		goto TASK2
	case "3":
		goto TASK3
	case "4":
		goto TASK4
	case "5":
		goto TASK5
	case "6":
		goto TASK6
	case "7":
		goto TASK7
	case "8":
		goto TASK8
	case "9":
		goto TASK9
	case "10":
		goto TASK10
	case "11":
		goto TASK11
	default:
		goto MENU
	}

	// once a time : get dataset files in local
TASK0:
	//csclient.DownloadDirectory(os.Stdout, "spls", nil, TEMP_PATH, TEMP_ROOT_DIR)

	fmt.Printf("Please enter the city:")
	reader = bufio.NewReader(os.Stdin)
	line, _, err = reader.ReadLine()
	if err != nil {
		log.Fatalln(err)
	}
	city = string(line)

	fmt.Printf("Please enter the event:")
	reader = bufio.NewReader(os.Stdin)
	line, _, err = reader.ReadLine()
	if err != nil {
		log.Fatalln(err)
	}
	event = string(line)

	fmt.Printf("Please enter the pilot:")
	reader = bufio.NewReader(os.Stdin)
	line, _, err = reader.ReadLine()
	if err != nil {
		log.Fatalln(err)
	}
	pilot = string(line)

	goto MENU
	// Task1:
TASK1:
	err = bqclient.CreateDataset(os.Stdout, GOOGLE_CLOUD_PROJECT, GOOGLE_DATASET, GOOGLE_REGION)
	if err != nil {
		log.Println(err)
	}

	// loop on all csv files in the directory
	csvdatafiles, err = ioutil.ReadDir(TEMP_ROOT_DIR + TEMP_PATH)
	if err != nil {
		log.Fatal(err)
	}
	// import csv into the BigQuery dataset
	for _, csvfile := range csvdatafiles {
		if !csvfile.IsDir() {
			if strings.ToUpper(csvfile.Name()[len(csvfile.Name())-3:]) == "CSV" {
				log.Println(csvfile.Name())
				err := bqclient.ImportCSVFromFile(GOOGLE_CLOUD_PROJECT, GOOGLE_DATASET, csvfile.Name()[:len(csvfile.Name())-4], TEMP_ROOT_DIR+TEMP_PATH+csvfile.Name())
				if err != nil {
					log.Fatal(err)
				}
			}

		}
	}
	goto MENU
	// Task 2: Events in a Certain City
TASK2:

	sql = "SELECT name FROM `drl.events` " +
		"WHERE city = \"" + city + "\""

	err = bqclient.QueryBasic(os.Stdout, GOOGLE_CLOUD_PROJECT, sql, GOOGLE_REGION)
	if err != nil {
		log.Fatal(err)
	}
	goto MENU
TASK3:
	//Task 3: Event Pilot Names
	sql = "SELECT p.name, ep.id AS event_pilot_id " +
		"FROM `drl.event_pilots` AS ep " +
		"LEFT JOIN `drl.pilots` AS p ON p.id = ep.pilot_id"
	err = bqclient.QueryBasic(os.Stdout, GOOGLE_CLOUD_PROJECT, sql, GOOGLE_REGION)
	if err != nil {
		log.Fatal(err)
	}
	goto MENU
	//Task 4: Pilots Who Flew in an Event
TASK4:

	sql = "SELECT p.name, e.name AS event_name " +
		"FROM `drl.event_pilots` AS ep " +
		"LEFT OUTER JOIN `drl.pilots` AS p ON p.id = ep.pilot_id " +
		"LEFT OUTER JOIN `drl.events` AS e ON e.id = ep.event_id " +
		"WHERE e.name = '" + event + "'"

	err = bqclient.QueryBasic(os.Stdout, GOOGLE_CLOUD_PROJECT, sql, GOOGLE_REGION)
	if err != nil {
		log.Fatal(err)
	}
	goto MENU

TASK5:
	//Task 5:Average Time of Rank 1 Round Finish
	sql = "WITH cte AS (SELECT rs.minimum_time FROM `drl.round_standings` AS rs WHERE `rank` = 1) " +
		"SELECT time" +
		"(timestamp_seconds" +
		"(CAST" +
		"  (AVG" +
		"    (UNIX_SECONDS" +
		"      (PARSE_TIMESTAMP('%H:%M.%S', minimum_time))" +
		"    )" +
		"AS INT64)" +
		") " +
		") AS avg FROM cte"
	err = bqclient.QueryBasic(os.Stdout, GOOGLE_CLOUD_PROJECT, sql, GOOGLE_REGION)
	if err != nil {
		log.Fatal(err)
	}
	goto MENU

TASK6:
	//Task 6:Clean and Combine Time Trial Data
	sql = "CREATE TABLE drl.time_trial_cleaned AS ( " +
		"SELECT " +

		"ttgpt.id AS time_trial_group_pilot_times_id, " +

		"ttgp.id AS time_trial_group_pilot_id, " +

		"ttg.id AS time_trial_group_id, " +

		"round_id, " +

		"CASE " +

		"WHEN ttgpt.time_adjusted IS NOT null THEN ttgpt.time_adjusted " +

		"WHEN ttg.racestack_scoring = 0 THEN ttgpt.time " +

		"ELSE ttgpt.racestack_time " +

		"END AS time " +

		"FROM `drl.time_trial_group_pilot_times` 		AS ttgpt " +
		"LEFT OUTER JOIN `drl.time_trial_group_pilots` 	AS ttgp ON ttgpt.time_trial_group_pilot_id = ttgp.id  " +
		"LEFT OUTER JOIN `drl.time_trial_groups` 		AS ttg 	ON ttgp.time_trial_group_id = ttg.id " +

		")"
	err = bqclient.QueryBasic(os.Stdout, GOOGLE_CLOUD_PROJECT, sql, GOOGLE_REGION)
	if err != nil {
		log.Fatal(err)
	}
	goto MENU
TASK7:

	//Task 7 : Fastest Time Trial at Event
	sql = "WITH cte AS " +

		"(SELECT " +

		"r.event_id, " +

		"r.name, " +

		"e.name AS event_name, " +

		"time " +

		"FROM `drl.time_trial_cleaned` AS ttc " +

		"LEFT OUTER JOIN `drl.rounds` AS r ON ttc.round_id = r.id " +

		"LEFT OUTER JOIN `drl.events` AS e ON e.id = r.event_id) " +

		"SELECT MIN(time) as fastest_time FROM cte WHERE event_name = '" + event + "'  " +
		"AND name = 'Time Trials' "
	err = bqclient.QueryBasic(os.Stdout, GOOGLE_CLOUD_PROJECT, sql, GOOGLE_REGION)
	if err != nil {
		log.Fatal(err)
	}
	goto MENU

TASK8:
	//Task 8:Pilot Heat Statistics

	sql = "SELECT " +

		"p.name AS pilot_name, " +

		"hs.heat_id AS heat_id, " +

		"hs.minimum_time, " +

		"hs.points " +

		"FROM `drl.heat_standings` AS hs " +

		"LEFT JOIN `drl.event_pilots` AS ep ON ep.id = hs.event_pilot_id " +

		"LEFT JOIN `drl.pilots` 		AS p ON p.id = ep.pilot_id " +

		"WHERE " +

		"name = '" + pilot + "' " +

		"AND " +

		"minimum_time != '" + pilot + "' " +

		"AND " +

		"minimum_time != ''"
	err = bqclient.QueryBasic(os.Stdout, GOOGLE_CLOUD_PROJECT, sql, GOOGLE_REGION)
	if err != nil {
		log.Fatal(err)
	}
	goto MENU

TASK9:
	//Task 9: Pilot Running Average Heat Time
	sql = "WITH cte AS " +

		"(SELECT p.name AS pilot_name, hs.heat_id, hs.points, hs.minimum_time " +

		"FROM `drl.heat_standings` AS hs " +

		"LEFT JOIN `drl.event_pilots` AS ep ON ep.id = hs.event_pilot_id " +

		"LEFT JOIN `drl.pilots` AS p ON p.id = ep.pilot_id " +

		"WHERE name = '" + pilot + "'  AND minimum_time != 'DNF' AND minimum_time != '') " +

		"SELECT " +

		"pilot_name, " +

		"heat_id " +

		"minimum_time, " +

		"points, " +

		"time " +

		"(timestamp_seconds " +

		"  (CAST " +

		"    (AVG " +

		"      (UNIX_SECONDS " +

		"        (PARSE_TIMESTAMP('%H:%M.%S', minimum_time)) " +

		"      ) " +

		"    OVER (ORDER BY heat_id ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) " +

		"  AS INT64) " +

		"  ) " +

		") " +

		"AS running_avg " +

		"FROM cte"
	err = bqclient.QueryBasic(os.Stdout, GOOGLE_CLOUD_PROJECT, sql, GOOGLE_REGION)
	if err != nil {
		log.Fatal(err)
	}
	goto MENU

TASK10:
	//Task 10:Pilot Time Improvements
	sql = "WITH cte AS " +

		"(SELECT " +

		"p.name AS pilot_name, hs.heat_id, hs.points, hs.minimum_time " +

		"FROM `drl.heat_standings` AS hs " +

		"LEFT JOIN `drl.event_pilots` AS ep ON ep.id = hs.event_pilot_id " +

		"LEFT JOIN `drl.pilots` AS p ON p.id = ep.pilot_id " +

		"WHERE name = '" + pilot + "'  AND minimum_time != 'DNF' AND minimum_time != ''), " +

		"cte2 AS " +

		"(SELECT " +

		"pilot_name, " +

		"heat_id, " +

		"minimum_time, " +

		"points, " +

		"time " +

		"(timestamp_seconds " +

		" (CAST " +

		"   (AVG " +

		"     (UNIX_SECONDS " +

		"       (PARSE_TIMESTAMP('%H:%M.%S', minimum_time)) " +

		"     ) " +

		"   OVER (ORDER BY heat_id ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) " +

		" AS INT64) " +

		" ) " +

		") " +

		"AS running_avg FROM cte) " +

		"SELECT *, " +

		"TIME_DIFF(PARSE_TIME('%H:%M.%S', minimum_time), running_avg, SECOND) as time_diff_from_avg FROM cte2"
	err = bqclient.QueryBasic(os.Stdout, GOOGLE_CLOUD_PROJECT, sql, GOOGLE_REGION)
	if err != nil {
		log.Fatal(err)
	}
	goto MENU

TASK11:
	//Task 11:Visualize Pilot Heat Statistics

	sql = "WITH cte AS " +

		"(SELECT p.name AS pilot_name, hs.heat_id, hs.points, hs.minimum_time " +
		"FROM `drl.heat_standings` AS hs " +

		"LEFT JOIN `drl.event_pilots` AS ep ON ep.id = hs.event_pilot_id " +

		"LEFT JOIN `drl.pilots` AS p ON p.id = ep.pilot_id " +

		"WHERE points != 0 AND minimum_time != 'DNF' AND minimum_time != ''), " +

		"cte2 AS " +

		"(SELECT " +

		"pilot_name, " +

		"heat_id, " +

		"minimum_time, " +

		"points, " +

		"time " +

		"(timestamp_seconds " +

		"(CAST " +

		"  (AVG " +

		"	(UNIX_SECONDS " +

		"	  (PARSE_TIMESTAMP('%H:%M.%S', minimum_time)) " +

		"	) " +

		" OVER (ORDER BY heat_id ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) " +

		"AS INT64) " +

		") " +

		") " +

		"AS running_avg FROM cte) " +

		"SELECT *, " +

		"TIME_DIFF(PARSE_TIME('%H:%M.%S', minimum_time), running_avg, SECOND) as time_diff_from_avg FROM cte2"
	err = bqclient.QueryBasic(os.Stdout, GOOGLE_CLOUD_PROJECT, sql, GOOGLE_REGION)
	if err != nil {
		log.Fatal(err)
	}
	goto MENU
}
