package client

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"../../feed"
)

var httpClient = &http.Client{}

type HTTPResp struct {
	Data   data   `json:"data"`
	Errors errors `json:"errors"`
}

type data interface{}
type errors interface{}

func (httpResp *HTTPResp) Retrieve(feedData feed.Feed, body io.Reader) error {
	start := time.Now()

	req, err := http.NewRequest(feedData.Method, feedData.URI, body)
	if err != nil {
		return err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(httpResp)
	if err != nil {
		return err
	}

	log.Printf("%v: %v %v [%v]", feedData.Action, feedData.Method, feedData.URI, time.Since(start))
	return nil
}
