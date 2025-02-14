package main

import "fmt"

func main() {

	fmt.Println("this is a GO web server.")
}

// creating the type for the courses
type Course struct {
	Price      float32 `json:"price"`
	CourseId   string  `json:"courseid"`
	CourseName string  `json:"coursename"`
	Aurhor     *Author
}

// the author type
type Author struct {
	Name  string `json:"authorname"`
	Books []string
}

// temporary DB
var course []Course

// middleware to check there is a course name & id

func (course *Course) isValid() bool {
	return course.CourseId != "" && course.CourseName != ""
}
