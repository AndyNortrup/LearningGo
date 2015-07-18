package main 

import (
	"os"
	"log"
	"net/rpc"
	"penn-state/system-integration/exercise2a/common"
)

var protocol = "tcp"
var host = "localhost:8080"

func main() {
	client, err := rpc.DialHTTP(protocol, host)
	if err != nil {
		log.Fatalf("Failed to connect to server.\n%v", err.Error())
	}
	
	var result bool
	args := createStudentInfo()
	err = client.Call("CourseInfo.WriteStudentInfo", args, &result)
	if err != nil {
		log.Printf("Failed to call proceadure. \n %v", err.Error())
		os.Exit(1)
	}
	
	log.Printf("Item added to Course System.")
}

func createStudentInfo() student.StudentInfo {
	var info student.StudentInfo
	
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
	
	return info
}