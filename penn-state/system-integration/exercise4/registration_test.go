package main

//TestUnmarshalRequest runs tests on unmarshallRequest
import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var oneStudentJSONBlob = []byte(`[
  {
    "StudentID": 10,
    "Name": "John Smith",
    "SSN": "123-45-6789",
    "email": "jsmith@yahoo.com",
    "HomePhone": "555-123-4321",
    "HomeAddress": "18 Longwood Dr, Kennebunk, ME 04043",
    "LocalAddress": "2305 S I St, Tacoma, WA 98405",
    "EmergencyContact": "Jen",
    "ProgramID": 1111,
    "PaymentID": "123-11",
    "AcadmenicStatus": 1
  }]
  `)

var twoStudentJSONBlob = []byte(`[
  {
    "StudentID": 10,
    "Name": "John Smith",
    "SSN": "123-45-6789",
    "email": "jsmith@yahoo.com",
    "HomePhone": "555-123-4321",
    "HomeAddress": "18 Longwood Dr, Kennebunk, ME 04043",
    "LocalAddress": "2305 S I St, Tacoma, WA 98405",
    "EmergencyContact": "Jen",
    "ProgramID": 1111,
    "PaymentID": "123-11",
    "AcadmenicStatus": 1
  },
  {
    "StudentID": 11,
    "Name": "Jane Smith",
    "SSN": "123-45-6789",
    "email": "jane.smith@yahoo.com",
    "HomePhone": "555-123-4321",
    "HomeAddress": "18 Longwood Dr, Kennebunk, ME 04043",
    "LocalAddress": "2305 S I St, Tacoma, WA 98405",
    "EmergencyContact": "Jen",
    "ProgramID": 1111,
    "PaymentID": "123-11",
    "AcadmenicStatus": 1
  }]`)

//inisValidRequest is invalid because the first studentID is a negative numbers
var inisValidRequest = []byte(`[
    {
      "StudentID": -1,
      "Name": "John Smith",
      "SSN": "123-45-6789",
      "email": "jsmith@yahoo.com",
      "HomePhone": "555-123-4321",
      "HomeAddress": "18 Longwood Dr, Kennebunk, ME 04043",
      "LocalAddress": "2305 S I St, Tacoma, WA 98405",
      "EmergencyContact": "Jen",
      "ProgramID": 1111,
      "PaymentID": "123-11",
      "AcadmenicStatus": 1
    },
    {
      "StudentID": 11,
      "Name": "Jane Smith",
      "SSN": "123-45-6789",
      "email": "jane.smith@yahoo.com",
      "HomePhone": "555-123-4321",
      "HomeAddress": "18 Longwood Dr, Kennebunk, ME 04043",
      "LocalAddress": "2305 S I St, Tacoma, WA 98405",
      "EmergencyContact": "Jen",
      "ProgramID": 1111,
      "PaymentID": "123-11",
      "AcadmenicStatus": 1
    }]`)

var emptyRequest = []byte(``)

//Test 1: Inputs with different size StudentLists return the right number of
//        variables
//Test 2: Errors in the unmarshalling should return an error.  Test with an
//        overloade in the int type
func TestUnmarshalRequest(testing *testing.T) {

	reader := bytes.NewReader(oneStudentJSONBlob)
	studentList, err := unmarshallRequest(reader)
	if err != nil {
		testing.Fatalf("Failed to unmarshal json in UnmarshalRequest. \n\t%v", err.Error())
	} else if len(studentList) != 1 {
		testing.Fatalf("Incorrect number of students returned.  Expected: %v\tReceived: %v", 1, len(studentList))
	}

	reader = bytes.NewReader(twoStudentJSONBlob)
	studentList, err = unmarshallRequest(reader)
	if err != nil {
		testing.Fatalf("Failed to unmarshal json in UnmarshalRequest. \n\t%v", err.Error())
	} else if len(studentList) != 2 {
		testing.Fatalf("Incorrect number of students returned.  Expected: %v\tReceived: %v", 2, len(studentList))
	}
}

//TestValidateRequest tests to ensure that validateRequests is not permitting
//unauthorized values.
func TestValidateRequest(t *testing.T) {
	getMethod := "GET"
	invalidMIME := "text/plain"

	request, _ := http.NewRequest(postMethod, "", nil)
	request.Header.Set(contentType, mimeType)

	if err0 := isValidRequest(request); err0 != nil {
		t.Errorf("Failed to validate good request.\n\t%v", err0.Error())
	}

	//Change the method to be invalid
	request.Method = getMethod
	if err1 := isValidRequest(request); err1 == nil {
		t.Errorf("Failed to disqualify wrong http.Method value.\n\t%v", err1.Error())
	}

	//reset method and change content type to be invalid
	request.Method = postMethod
	request.Header.Set(contentType, invalidMIME)
	if err2 := isValidRequest(request); err2 == nil {
		t.Errorf("Failed to disqualify wrong mime type data. \n\t%v", err2.Error())
	}
}

//TestService creates an instance of the server, and sends requests to validate
// the responses.
func TestService(t *testing.T) {
	app := &AddStudentService{}
	s := httptest.NewServer(app)
	defer s.Close()

	//Send a request to the server with one student
	t.Log("Sending Request for one student")
	response := sendRequest(t, oneStudentJSONBlob, s.URL)
	if response.CountAdded != 1 {
		t.Errorf("Incorrect student count returned. Expected: 1 \t Received: %v",
			response)
	}

	//Send a request to the server with one student
	t.Log("Sending Request for two Students")
	response = sendRequest(t, twoStudentJSONBlob, s.URL)
	if response.CountAdded != 2 {
		t.Errorf("Incorrect student count returned. Expected: 2 \t Received: %v",
			response)
	}

	sendBadGetRequest(t, s.URL)

	sendEmptyRequest(t, s.URL)

	sendInvliadRequest(t, s.URL)

}

//Send request manages the process of sending a request to the server
func sendRequest(t *testing.T, requestBody []byte, url string) *StudentAdded {

	//Post the request information to the server
	response, err := http.Post(url, mimeType, bytes.NewReader(requestBody))
	if err != nil {
		t.Fatalf("Failed to post request. \n\t%v", err.Error())
	}

	//convert the response into a StudentsAdded struct
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal("Unable to read response body.")
	}
	studentsAdded := &StudentAdded{}
	json.Unmarshal(responseBody, studentsAdded)

	t.Logf("Sent valid request.  HTTP Status: %v\n\t",
		response.Status)
	return studentsAdded
}

//Send request manages the process of sending a request to the server
func sendBadGetRequest(t *testing.T, url string) {
	response, _ := http.Get(url)
	t.Logf("Sent get request.  HTTP Status: %v", response.Status)
	if response.StatusCode != http.StatusInternalServerError {
		t.Error("Server failed to reject GET request")
	}
}

func sendEmptyRequest(t *testing.T, url string) {
	response, _ := http.Post(url, mimeType, bytes.NewReader(emptyRequest))
	t.Logf("Sent empty request.  HTTP Status: %v", response.Status)
	if response.StatusCode != http.StatusInternalServerError {
		t.Error("Failed to reject empty request.")
	}
}

func sendInvliadRequest(t *testing.T, url string) {
	response, _ := http.Post(url, mimeType, bytes.NewReader(inisValidRequest))
	t.Logf("Sent invalid request.  HTTP Status: %v", response.Status)
	if response.StatusCode != http.StatusInternalServerError {
		t.Error("Failed to reject request with negative Student ID Number")
	}
}
