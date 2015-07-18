package main

import (
	"bytes"
	"encoding/json"
	"log"
	"penn-state/system-integration/student"
	"sync"

	"github.com/bitly/go-nsq"
)

var topic = "AddToCourseInfo"
var channel = "AddStudent"

type addStudentHandler struct{}

func main() {

	//Create a wait group to keep the server running.
	wg := &sync.WaitGroup{}
	wg.Add(1)

	//Create a new NSQ configuration and a new consumer
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, channel, config)

	//Check for errors
	if err != nil {
		log.Fatal("Failed to create consumer.")
	}

	//Create a new handler to handle requests coming from the producer
	handler := new(addStudentHandler)
	consumer.AddHandler(handler)

	//Connect to the NSQ server
	err = consumer.ConnectToNSQD("127.0.0.1:4150")

	wg.Wait()
}

//HandleMessage recieves the message from NSQ and in this case print them to the
//standard output.  In the real world this method would write the information
//into a database.
func (handler addStudentHandler) HandleMessage(message *nsq.Message) error {
	var student integrate.StudentInfo
	decoder := json.NewDecoder(bytes.NewReader(message.Body))
	err := decoder.Decode(student)

	if err != nil {
		return err
	}

	//Print to log
	log.Printf("Name: %v\n Student ID: %v", student.Name, student.StudID)
	return nil
}
