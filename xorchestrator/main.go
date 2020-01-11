package main

import (
	"fmt"
	"log"
	"os"

	"./feed"
	"./multix"
)

var feeds map[string]feed.Feed

func init() {
	feeds = feed.Load()
}

func main() {
	selectedFeeds := feed.Select(feeds, os.Args[1:])
	results := multix.RunRequestsAndProcess(selectedFeeds)
	printResults(results)
}

func printResults(results map[string][]feed.Retriever) {
	log.Print("-----------------------")
	for k, values := range results {
		fmt.Printf("%v:\n", k)
		for _, v := range values {
			fmt.Printf("%+v\n", v)
		}
		fmt.Printf("\n")
	}
}
