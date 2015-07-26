package main

import (
	"net/http"
	"log"
	"net"
	"net/rpc"
	"os"
)

var protocol = "tcp"
var host = ":8080"

func main() {
	
	//Check for command line arguments
	if len(os.Args) > 1 {
		host = getArgs()
	}
	
	courseInfo := NewCourseInfo()
	rpc.Register(courseInfo)
	rpc.HandleHTTP()
	ln, err := net.Listen(protocol, host)
	if err != nil {
		log.Fatalf("Failed to open network listener.\n%v", err.Error())
	}
	http.Serve(ln, nil)
}

//retrieve the remote host argument
func getArgs() string {
	return os.Args[1]	
}

