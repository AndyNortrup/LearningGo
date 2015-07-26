package roster

//Student is a struct that holds data about a student.
type Student struct {
	StudID           int    `json:"StudentID"`
	Name             string `json:"Name"`
	SSN              string `json:"SSN"`
	EmailAddress     string `json:"email"`
	HomePhone        string `json:"HomePhone"`
	HomeAddress      string `json:"HomeAddress"`
	LocalAddress     string `json:"LocalAddress"`
	EmergencyContact string `json:"EmergencyContact"`
	Program          `json:"ProgramId"`
	Payment          `json:"PaymentID"`
	AcademicStatus   int `json:"AcademicStatus"`
}

//String implements the Stringer interface for printing as a String
func (student Student) String() string {
	return "Student ID: " + string(student.StudID) + " Name: " + student.Name
}
