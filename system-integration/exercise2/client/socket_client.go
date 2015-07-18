package main

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"penn-state/system-integration/student"
)

//Default network connections
var host = "localhost:8080"
var protocol = "tcp"

func main() {

	//Check for command line arguments
	if len(os.Args) > 1 {
		host = getArgs()
	}

	//Create an instance of the information to be sent
	info := integrate.StudentInfo{
		StudID:           1111,
		Name:             "Bob Smith",
		SSN:              "222-333-1111",
		EmailAddress:     "bsmith@yahoo.com",
		HomePhone:        "215-777-8888",
		HomeAddress:      "123 Tulip Road, Ambler, PA 19002",
		LocalAddress:     "321 Maple Avenue, Lion Town, PA 16800",
		EmergencyContact: "John Smith (215-222-6666)",
		ProgramID:        206,
		PaymentID:        "1111-206",
		AcademicStatus:   1,
	}

	client := NewTCPStudentClient(host)
	client.SendData(info)

}

//retrieve the remote host argument
func getArgs() string {
	return os.Args[1]
}

//StudentClient is used to send instances of StudentInfo over the network
type StudentClient struct {
	address, network string
	net.Conn
}

//NewTCPStudentClient constructs a new StudentClient instance using TCP
func NewTCPStudentClient(address string) StudentClient {
	result := StudentClient{network: protocol}
	if address != "" {
		result.address = address
	} else {
		result.address = "localhost:8080"
	}
	result.network = "tcp"

	return result
}

//SendData sends provided StudentInfo over the client's net.Conn
func (client StudentClient) SendData(student integrate.StudentInfo) {
	var err error
	client.Conn, err = net.Dial(client.network, client.address)
	defer client.Close()

	if err != nil {
		log.Fatal("Failed to connect to remote host")
	}

	encoder := json.NewEncoder(client)
	encoder.Encode(student)
}
