package multix

import (
	"log"
	"sync"

	"../feed"
	client "../feed/clients"
)

type Result struct {
	FeedName string
	Response feed.Retriever
}

func RunRequestsAndProcess(feeds []feed.Feed) {
	// Make an unbuferred channel, so we don't enqueue any result
	resultsChan := make(chan *Result)

	// Setup a wait group so we can process all the feeds.
	var waitGroup sync.WaitGroup

	// Set the number of goroutines we need to wait for while
	// they process the individual feeds.
	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		log.Print("Processing feed", feed)
		go runRequest(feed, &waitGroup, resultsChan)
	}

	// Launch a goroutine to monitor when all the work is done.
	go waitAndCloseChan(&waitGroup, resultsChan)

	// Will dynamically display the results and process them while receiving
	processResults(resultsChan)
}

func runRequest(feed feed.Feed, waitGroup *sync.WaitGroup, channel chan<- *Result) {
	httpResp := &client.HTTPResp{}
	httpResp.Retrieve(feed, nil)

	// debugResp := client.DebugResp{}
	// debugResp.Retrieve(feed, nil)

	channel <- &Result{
		FeedName: feed.Action,
		Response: httpResp,
	}
	waitGroup.Done()
}

func waitAndCloseChan(waitGroup *sync.WaitGroup, channel chan *Result) {
	// Wait for everything to be processed.
	log.Print("Waiting for all go routines to finish...")
	waitGroup.Wait()

	// Close the channel to signal to the Display
	// function that we can exit the program.
	close(channel)
}

func processResults(results <-chan *Result) {
	for result := range results {
		log.Printf("%v result: %+v", result.FeedName, result.Response)
	}
}
