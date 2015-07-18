package exercise

import "encoding/json"

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

	json.Encoder
	json.Decoder
}
