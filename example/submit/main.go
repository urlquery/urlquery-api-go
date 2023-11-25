package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urlquery/urlquery-api-go"
)

// Command-line arguments
var key = flag.String("apikey", "", "urlquery API key")
var url = flag.String("submit", "", "submit url")

func main() {
	flag.Parse()

	if *url == "" || *key == "" {
		flag.PrintDefaults()
		fmt.Println("\nPlease provide a APIKEY and URL to submit")
		os.Exit(0)
	}
	urlquery.SetDefaultKey(*key)

	submission := urlquery.SubmitJob{
		Url:    *url,
		Access: "public", // public, restricted, private
	}

	queue, err := urlquery.Submit(submission)
	if err != nil {
		log.Fatal("submission failed:", err)
	}

	fmt.Println(queue)

	// status -> queued, processing, done, failed
	for queue.Status != "done" && queue.Status != "failed" {
		time.Sleep(3 * time.Second)
		queue, err = urlquery.GetQueueStatus(queue.QueueID)
		if err != nil {
			panic(err)
		}

		fmt.Printf("\rStatus: %s", queue.Status)
	}

	if queue.Status == "done" {
		report, _ := urlquery.GetReport(queue.ReportID)
		fmt.Println(report.String())
		fmt.Println("\nReport at: https://urlquery.net/report/" + report.ID)

	} else if queue.Status == "failed" {
		fmt.Println("Processing URL failed!")
	}

}
