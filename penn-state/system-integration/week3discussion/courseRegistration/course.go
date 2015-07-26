package registration

//Course represents an individual course taught at the University
type Course struct {
	CourseID    int
	CourseName  string
	SectionNo   int
	Prerquisite []Course
}
