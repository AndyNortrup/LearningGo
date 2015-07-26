package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"penn-state/system-integration/student"
)

const path string = "AddStudent"

//AddStudentService is an instance of http.Handler
type AddStudentService struct{}

//AddStudent processes a request to add a student to the student registration
// system
func (service AddStudentService) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	err := isValidRequest(request)
	if err != nil {

		response := StudentAdded{
			CountAdded: 0,
			Error:      "Invalid request.  Requires POST, application/json data."}
		response.sendResponse(writer)
		return
	}

	students, err := unmarshallRequest(request.Body)
	if err != nil {

		response := StudentAdded{
			CountAdded: 0,
			Error:      "Unable to unmarshal JSON request"}
		response.sendResponse(writer)
		return
	}

	//Nominal: Validate business rules
	for _, student := range students {
		if !isValidStudent(student) {
			response := StudentAdded{
				CountAdded: 0,
				Error:      "Invalid Submitted Data"}
			response.sendResponse(writer)
			return
		}
	}

	//Nominal: Send data to database via messaging system
	//nsq.Send(Student)

	response := StudentAdded{
		CountAdded: len(students),
		Error:      ""}
	response.sendResponse(writer)
}

//isValidRequest validates that the request meets the method and data type formats
//requests should only be of the type POST and be of MIME type application/json
func isValidRequest(request *http.Request) error {

	//verify the request method
	if request.Method != postMethod {
		return errors.New("Invalid method. Expected: " + postMethod + " Recieved: " + request.Method)
	}

	//verify the content type
	requestContentType := request.Header.Get(contentType)
	if requestContentType != mimeType {
		return errors.New("Invalid MIME type expected: " + mimeType + " recieved " + contentType)
	}
	return nil
}

//unmarshallRequest takes a reader (likely derived from a http.Request) and
//converts it to a StudentList and any errors that occured
func unmarshallRequest(request io.Reader) ([]integrate.StudentInfo, error) {
	var student []integrate.StudentInfo
	err := json.NewDecoder(request).Decode(&student)
	return student, err
}

func isValidStudent(students integrate.StudentInfo) bool {
	if students.StudID <= 0 {
		return false
	}
	return true
}
