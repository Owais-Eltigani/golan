package main

// creating the type for the Courses
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var Courses []Course

func init() {
	Courses = []Course{
		{"1", "ReactJS", 299, &Author{"John Doe", "http://johndoe.com"}},
		{"2", "NodeJS", 199, &Author{"Jane Smith", "http://janesmith.com"}},
		{"3", "Angular", 299, &Author{"Bob Wilson", "http://bobwilson.com"}},
	}

}
