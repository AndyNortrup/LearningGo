package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	//
)

const requestUsername string = "username"
const requestPassword string = "password"

type User struct {
	username, password string
}

type AuthData struct {
	domain string `json:"domain"`
	users  []struct {
		username string `json:"username"`
		password string `json:"password"`
	}
}

func main() {
	log.Printf("Application Started\n")
	startServer()
}

func startServer() {
	log.Printf("Sever Started\n")

	http.HandleFunc("/api/2/domain/", handleRequest)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	log.Printf("Handling Request\n")
	requestDomain := domainFrom(r)
	username, password := userFrom(r)

	fmt.Fprintf(w, "Domain: %v\t Auth Data:%v\n", requestDomain, username, password
	)
}

func userFrom(r *http.Request) (string, string) {
	var requestVals = new(User)

	username = r.FormValue(requestUsername)
	password = r.FormValue(requestPassword)

	return username, password
}

func domainFrom(r *http.Request) string {
	return strings.Split(r.RequestURI, "/")[4]
}
