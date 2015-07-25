package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"penn-state/system-integration/student"
)

const path string = "AddStudent"

//StudentAdded will be used to send a response back to the user
type StudentAdded struct {
	CountAdded int `json:"countAdded"`
	Error      int `json:"errors"`
}

//AddStudentService is an instance of http.Handler
type AddStudentService struct{}

//AddStudent processes a request to add a student to the student registration
// system
func (service AddStudentService) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	err := validRequest(request)
	if err != nil {
		sendResponse(0,
			"Invalid request.  Requires POST, application/json data.",
			writer)
		return
	}

	students, err := unmarshallRequest(request.Body)
	if err != nil {
		sendResponse(0,
			"Unable to unmarshal JSON request",
			writer)
		return
	}

	//Nominal: Validate business rules
	for _, student := range students {
		if !isValidStudent(student) {
			sendResponse(0, "Invaid Submitted Data", writer)
			return
		}
	}

	//Nominal: Send data to database via messaging system

	sendResponse(len(students), "", writer)
}

//validRequest validates that the request meets the method and data type formats
//requests should only be of the type POST and be of MIME type application/json
func validRequest(request *http.Request) error {

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

//sendResponse sends a response back to the user.  If the error string is
func sendResponse(studentsAdded int, error string, writer http.ResponseWriter) {
	success := new(StudentAdded)
	writer.Header().Set(contentType, mimeType)

	if error == "" {
		success.CountAdded = studentsAdded
		writer.WriteHeader(http.StatusOK)
	} else {
		success.CountAdded = 0
		writer.WriteHeader(http.StatusInternalServerError)
	}

	responseBody, _ := json.Marshal(success)
	writer.Write(responseBody)
}
