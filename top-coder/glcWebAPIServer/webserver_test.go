package main

import (
	"net/http"
	"testing"
)

func TestUserFromPost(t *testing.T) {

	testRequest, err := http.NewRequest("POST", ":8080", nil)

	if err != nil {
		t.Fatal("Failed to create http request.")
	}

	testRequest.PostForm.Add(requestUsername, requestUsername)
	testRequest.PostForm.Add(requestPassword, requestPassword)

	testResult := userFrom(testRequest)

	if testResult.username == requestUsername &&
		testResult.password == requestPassword {
		t.Fatalf("Failed to extract username and password from HTTP Post Request")
	}
}
