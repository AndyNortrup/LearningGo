package main

import (
	"os"
	"log"
	"penn-state/system-integration/common"
	"net"
)

var host string = "localhost:8080"
var protocol string = "tcp"

func main() {
	var info exercise.StudentInfo
	info.StudID = 1111
	info.Name = "Bob Smith"
	info.SSN = "222-333-1111"
	info.EmailAddress = "bsmith@yahoo.com"
	info.HomePhone = "215-777-8888"
	info.HomeAddress = "123 Tulip Road, Ambler, PA 19002"
	info.LocalAddress = "321 Maple Avenue, Lion Town, PA 16800"
	info.EmergencyContact = "John Smith (215-222-6666)"
	info.ProgramID = 206
	info.PaymentID = "1111-206"
	info.AcademicStatus = 1
	
	conn, err := net.Dial(protocol, host)
	if err != nil {
		log.Fatal("Failed to connect to server.")
		os.Exit(1)
	}
	
	info.Encode(conn)
	
}
