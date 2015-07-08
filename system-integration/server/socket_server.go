package main 

import (
	"log"
	"net"
	"penn-state/system-integration/common"
)

var host string = ":8080"
var protocol string = "tcp"

func main() {
	ln, err := net.Listen(protocol, host)
	
	if err != nil {
		log.Fatal("Failed to open server.")
	}
	
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("Failed to accept connection.")
		}
		
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	student := exercise.Decode(conn)
	log.Print("Data Recieved: %+v", student)
}