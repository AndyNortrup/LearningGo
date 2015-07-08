package exercise

import (
	"encoding/json"
	"io"
	"log"
)

//StudentInfo is a struct that holds data about a student.
type StudentInfo struct {
	StudID           int    `json:"StudentID"`
	Name             string `json:"Name"`
	SSN              string `json:"SSN"`
	EmailAddress     string `json:"email"`
	HomePhone        string `json:"HomePhone"`
	HomeAddress      string `json:"HomeAddress"`
	LocalAddress     string `json:"LocalAddress"`
	EmergencyContact string `json:"EmergencyContact"`
	ProgramID        int    `json:"ProgramId"`
	PaymentID        string `json:"PaymentID"`
	AcademicStatus   int    `json:"AcademicStatus"`
}

//Encode encodes an instance of StudentInfo into JSON.
func (info StudentInfo) Encode(writer io.Writer) {
	marsheler := json.NewEncoder(writer)
	marsheler.Encode(info)
}

//Decode converts JSON from a reader into an instance of StudentInfo
func Decode(reader io.Reader) StudentInfo {
	var info StudentInfo
	Decoder := json.NewDecoder(reader)
	err := Decoder.Decode(&info)
	if err != nil {
		log.Fatalf("Error decoding Student Info.\n%v", err)
	}

	return info
}
