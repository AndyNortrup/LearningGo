package registration

//CourseSchedule is a representation of when a course is scheduled
//at the university
type CourseSchedule struct {
	CourseScheduleID int
	Couse            int
	Semester         string
	ScheduleTime     string
	Location         string
	Capacity         int
	Faculty
	SpecialNeeds string
	Availability string
}
