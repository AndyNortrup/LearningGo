package roster

//Program represents a program of study in the university
type Program struct {
	ProgramID  int
	Name       string
	Department string
	College    string
	Students   []Student
}
