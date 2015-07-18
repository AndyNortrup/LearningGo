package main

import (
	"log"
	"penn-state/system-integration/exercise2a/common"
)

type CourseInfo struct {
	entries []student.StudentInfo
}

func NewCourseInfo() *CourseInfo {
	return &CourseInfo {
		entries: make([]student.StudentInfo, 0),
	}
}

func (ci *CourseInfo) WriteStudentInfo(info *student.StudentInfo, result *bool) error {
	ci.entries = append(ci.entries, *info)
	*result = true
	log.Printf("Student Entry Count: %v", len(ci.entries))
	return nil
}
