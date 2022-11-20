package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/bigquery"
	"github.com/compagnon/go-clients/pkg/googlecloud/bqclient"
	"google.golang.org/api/iterator"
)

const GOOGLE_CLOUD_PROJECT = "idyllic-pact-205716"

func main() {

	fmt.Println(bqclient.QueryBasic(os.Stdout, GOOGLE_CLOUD_PROJECT, "SELECT name FROM `bigquery-public-data.usa_names.usa_1910_2013` "+
		"WHERE state = \"TX\" "+
		"LIMIT 100", "US"))

	/*
		projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
		if projectID == "" {
			fmt.Println("GOOGLE_CLOUD_PROJECT environment variable must be set.")
			projectID = "idyllic-pact-205716"
		}

		// [START bigquery_simple_app_client]
		ctx := context.Background()

		client, err := bigquery.NewClient(ctx, projectID)
		if err != nil {
			log.Fatalf("bigquery.NewClient: %v", err)
		}
		defer client.Close()
		// [END bigquery_simple_app_client]

		rows, err := query(ctx, client)
		if err != nil {
			log.Fatal(err)
		}
		if err := printResults(os.Stdout, rows); err != nil {
			log.Fatal(err)
		}
	*/
}

// query returns a row iterator suitable for reading query results.
func query(ctx context.Context, client *bigquery.Client) (*bigquery.RowIterator, error) {

	// [START bigquery_simple_app_query]
	query := client.Query(
		`SELECT
			CONCAT(
				'https://stackoverflow.com/questions/',
				CAST(id as STRING)) as url,
			view_count
		FROM ` + "`bigquery-public-data.stackoverflow.posts_questions`" + `
		WHERE tags like '%google-bigquery%'
		ORDER BY view_count DESC
		LIMIT 10;`)
	return query.Read(ctx)
	// [END bigquery_simple_app_query]
}

// [START bigquery_simple_app_print]
type StackOverflowRow struct {
	URL       string `bigquery:"url"`
	ViewCount int64  `bigquery:"view_count"`
}

// printResults prints results from a query to the Stack Overflow public dataset.
func printResults(w io.Writer, iter *bigquery.RowIterator) error {
	for {
		var row StackOverflowRow
		err := iter.Next(&row)
		if err == iterator.Done {
			return nil
		}
		if err != nil {
			return fmt.Errorf("error iterating through results: %w", err)
		}

		fmt.Fprintf(w, "url: %s views: %d\n", row.URL, row.ViewCount)
	}
}

// [END bigquery_simple_app_print]
// [END bigquery_simple_app_all]
