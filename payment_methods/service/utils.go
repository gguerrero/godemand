package service

import (
	"encoding/json"
	"net/http"
)

type data interface{}
type okResp struct {
	Data data `json:"data"`
}

type errorResp struct {
	Message string `json:"message"`
	Code    int    `json:"statusCode"`
}

func jsonResponse(writer http.ResponseWriter, data data) {
	initializeHeader(writer, http.StatusOK)
	writeResponse(writer, okResp{data})
}

// Render an error response with the given error status and error message
func jsonErrorResponse(writer http.ResponseWriter, statusCode int, message string) {
	initializeHeader(writer, statusCode)
	writeResponse(writer, errorResp{message, statusCode})
}

func initializeHeader(writer http.ResponseWriter, statusCode int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
}

func writeResponse(writer http.ResponseWriter, response interface{}) {
	err := json.NewEncoder(writer).Encode(&response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
