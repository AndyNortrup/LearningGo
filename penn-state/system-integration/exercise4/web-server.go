package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

const contentType string = "Content-Type"
const mimeType string = "application/json"
const postMethod string = "POST"
const listenPort string = ":8080"

func main() {
	service := &AddStudentService{}

	//Create a request router to send http requests to the right function
	router := mux.NewRouter()

	//Add a route for our Add Student service
	router.HandleFunc(path, service.ServeHTTP)

	//Only Accept POST requests
	router.Methods(postMethod)

	//Oonly Accept json Content-Types
	router.Headers(contentType, mimeType)

	//Listen on port 8080
	http.ListenAndServe(listenPort, router)
}
