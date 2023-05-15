package main

import (
	statscollector "beanbot/internal/stats-collector"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var help = `beancounter <outputPath> <intervalSeconds>

This tool will collect system statistics and create/update a report at <outputPath> every <intervalSeconds>.`

// beancounter pulls data from the system at a set interval and writes a report
// to a file. This file can then be exposed to the interactive bot's Docker container
// so the data can be reported there.
//
// This expects a call signature: <executable> <path> <interval>
//
// A file is created or updated at <path> to container the latest report data and
// the collection interval is <interval> seconds.
//
// If any of the input parameters are not as expected (i.e. fewer than 2, interval not a number, ...)
// the program fails critically before starting the iteration. Errors during the iteration are
// handled and logged.
func main() {
	if len(os.Args) != 3 {
		fmt.Println(help)
		log.Fatal("beancounter requires one argument representing the path to write reports to.")
	}

	outputPath := os.Args[1]
	intervalSeconds, intErr := strconv.Atoi(os.Args[2])

	if intErr != nil {
		log.Fatal("The provided interval is not a valid number.")
	}

	log.Println(fmt.Sprintf("😸 Counting beans (report path=%s, interval=%ds)...", outputPath, intervalSeconds))
	for {
		log.Println("Collecting new datapoints.")

		summary, collectErr := statscollector.CollectStatistics()

		if collectErr != nil {
			log.Println(fmt.Sprintf("Failed to collect system statistics: %s", collectErr))
		}
		summaryJson, jsonErr := json.Marshal(&summary)

		if jsonErr != nil {
			log.Println("Failed to convert report to json.")
		}

		writeErr := os.WriteFile(outputPath, summaryJson, 0666)

		if writeErr != nil {
			log.Println("Failed to write report to disk.")
		}

		log.Println(fmt.Sprintf("Updated report at %s.", outputPath))

		time.Sleep(time.Duration(intervalSeconds) * time.Second)
	}
}
