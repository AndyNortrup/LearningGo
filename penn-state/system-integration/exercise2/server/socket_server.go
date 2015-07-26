package main

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"penn-state/system-integration/student"
)

func main() {

	host := ":8080"

	//See if there are arguments in the command line to set host and port
	//otherwise listen on localhost:8080
	if len(os.Args) > 1 {
		host = getArgs()
	}

	server := NewTCPStudentServer(host)
	server.listen()
}

//getArgs retrieves the value of the first argument on the command line
func getArgs() string {
	return os.Args[1]
}

//StudentServer is a server that listens on a network port and receives
//instances of StudentInfo
type StudentServer struct {
	Host     string
	Protocol string
	net.Listener
}

//NewTCPStudentServer creates an instance of StudentServer
func NewTCPStudentServer(host string) StudentServer {
	server := new(StudentServer)
	if host != "" {
		server.Host = host
	}
	server.Protocol = "tcp"

	return *server
}

func (server StudentServer) listen() {
	var err error
	server.Listener, err = net.Listen(server.Protocol, server.Host)

	//Check for errors
	if err != nil {
		log.Fatal("Failed to open server.")
	}

	//as long as the program is running, accept connection requests.
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal("Failed to accept connection.")
		}

		//open a new goroutine for each incoming connection
		go server.handleConnection(conn)
	}
}

//Handle connection takes a connection, decodes the contents and prints them to the
//log.
func (server StudentServer) handleConnection(conn net.Conn) {
	student := new(integrate.StudentInfo)
	decoder := json.NewDecoder(conn)
	decoder.Decode(student)
	log.Printf("Recieved Data: " + student.String())
}
