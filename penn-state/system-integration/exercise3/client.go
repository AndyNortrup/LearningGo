package main

import (
	"encoding/json"
	"log"
	"penn-state/system-integration/exercise2/common"

	"github.com/bitly/go-nsq"
)

var topic = "AddToCourseInfo"
var channel = "AddStudent"

func main() {

	//Create a new producer which points to an instance of nsqd
	p, err := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())
	if err != nil {
		log.Fatal("Failed to create producer.")
	}

	//Create an instance of StudentInfo and convert it into a JSON byte
	//array
	bytes, _ := json.Marshal(createStudentInfo())

	//Publish the message to the nsqd instance
	p.Publish(topic, bytes)
}

func createStudentInfo() exercise.StudentInfo {
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

	return info
}
