package client

import (
	"fmt"
	"io"

	"../../feed"
)

type DebugResp struct{}

func (DebugResp) Retrieve(feedData feed.Feed, body io.Reader) error {
	fmt.Println("--------------------------------")
	fmt.Printf("Action -> %v\n", feedData.Action)
	fmt.Printf("URI    -> %v\n", feedData.URI)
	fmt.Printf("Method -> %v\n", feedData.Method)
	fmt.Println("--------------------------------")

	return nil
}
