package models

import "errors"

var students = map[uint64]Student{
	1: {
		Name:  "John",
		Age:   20,
		Grade: "F",
	},
	2: {
		Name:  "Jain",
		Age:   20,
		Grade: "A+",
	},
	3: {
		Name:  "June",
		Age:   20,
		Grade: "C-",
	},
}

var ErrStudentNotFound = errors.New("Student record not found")

func FetchStudentData(studentId uint64) (Student, error) {
	s, ok := students[studentId]
	if !ok {
		return Student{}, ErrStudentNotFound
	}
	return s, nil
}

func FetchAllStudentData() map[uint64]Student{
	return students
}
