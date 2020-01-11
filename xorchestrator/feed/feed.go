package feed

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Retriever interface {
	Retrieve(Feed, io.Reader) error
}

type Feed struct {
	Action string `json:"action"`
	URI    string `json:"url"`
	Method string `json:"method"`
}

const feedDataFile = "feed/data/feeds.json"

func Load() map[string]Feed {
	var feeds map[string]Feed

	file, err := os.Open(feedDataFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	json.NewDecoder(file).Decode(&feeds)

	return feeds
}

func Select(feeds map[string]Feed, feedNames []string) []Feed {
	var selectedFeeds []Feed

	for _, feedName := range feedNames {
		if feed, ok := feeds[feedName]; ok {
			selectedFeeds = append(selectedFeeds, feed)
		} else {
			log.Printf(`Feed "%v" does not exists... skipping it!`, feedName)
		}
	}

	return selectedFeeds
}
