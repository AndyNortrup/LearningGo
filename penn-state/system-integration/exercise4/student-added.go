package main

import (
	"encoding/json"
	"net/http"
)

//StudentAdded will be used to send a response back to the user
type StudentAdded struct {
	CountAdded int    `json:"countAdded"`
	Error      string `json:"errors"`
}

//sendResponse sends a response back to the user.  If the error string is
func (response StudentAdded) sendResponse(writer http.ResponseWriter) {
	writer.Header().Set(contentType, mimeType)

	if response.CountAdded > 0 {
		writer.WriteHeader(http.StatusOK)
	} else {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	responseBody, _ := json.Marshal(response)
	writer.Write(responseBody)
}
