package main

import (
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
	multix.RunRequestAndProcess(selectedFeeds)
}
