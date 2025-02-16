package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// API controllers
func ServHome(wtr http.ResponseWriter, req *http.Request) { /* cspell:disable-line */

	fmt.Printf("welcoming new users.\n")
	// main page msg:
	fmt.Printf("Welcome to Golang server.\nthe Courses are:\n%#v", Courses)
	// Set content type header for JSON response
	wtr.Header().Set("Content-Type", "application/json")
	// wtr.Write([]byte(MSG))   //? for basic messages or raw data.

	json.NewEncoder(wtr).Encode(Courses)
}

// middleware to check there is a course name & id
func (course *Course) isEmpty() bool {
	return string(course.CourseId) == "" && course.CourseName == ""
}

func GetAllCourses(wtr http.ResponseWriter, req *http.Request) {

	fmt.Printf("retrieving all Courses.\n")
	// retrieving all Courses
	content := fmt.Sprintf("<h1>all Courses are: %v.<h1/>\n", Courses)

	wtr.Write([]byte(content))

	// //! real tutorial code:

	// // sending header
	// wtr.Header().Set("Content-Type", "application/json")

	// // encoding the data, weird didn't use json.marshal
	// json.NewEncoder(wtr).Encode(course)

	// //!

}

func GetCourseById(wtr http.ResponseWriter, req *http.Request) {

	fmt.Printf("asking for one course.\n")

	params := mux.Vars(req)
	courseid := params["courseid"] /* cspell:disable-line */

	courseIndex := linearSearch(courseid, Courses...) /* cspell:disable-line */

	if courseIndex == -1 {

		wtr.Write([]byte("no there.\n"))
		// json.NewEncoder(wtr).Encode("not found")
		return
	}

	json.NewEncoder(wtr).Encode(Courses[courseIndex])

}

func linearSearch(courseid string, Courses ...Course) int {

	for index, course := range Courses {

		if course.CourseId == courseid { /* cspell:disable-line */
			return index
		}
	}

	return -1

}

func AddNewCourse(wtr http.ResponseWriter, req *http.Request) {

	fmt.Printf("adding new course to the db.\n")
	wtr.Header().Set("Content-Type", "application/json")

	// check if the body is nill

	if req.Body == nil {
		json.NewEncoder(wtr).Encode("something wrong with body.\n")
		return
	}

	// receiving data from request.
	var course Course
	json.NewDecoder(req.Body).Decode(&course)

	// check if the course is empty
	if course.isEmpty() {

		json.NewEncoder(wtr).Encode("the course is empty\n")
		return
	}

	Courses = append(Courses, course)

	json.NewEncoder(wtr).Encode(Courses)

	//TODO don't depend on the user to send valid unique id, just generate id from
	//TODO  random.seed and check if the id isn't taken.

}

func UpdateCourseById(wtr http.ResponseWriter, req *http.Request) {

	fmt.Printf("updating one course by id.\n")

	params := mux.Vars(req)
	id := params["courseid"] /* cspell:disable-line */

	// searching for the course by id
	courseIndex := linearSearch(id, Courses...)

	if courseIndex == -1 {
		fmt.Printf("course not found.\n")
		json.NewEncoder(wtr).Encode("course by id not found.\n")
		return
	}

	// removing the course found by id.
	Courses = append(Courses[:courseIndex], Courses[courseIndex+1:]...)

	// parse new course data
	var updatedCourse Course
	json.NewDecoder(req.Body).Decode(&updatedCourse)

	// adding new course
	Courses = append(Courses, updatedCourse)

	content := fmt.Sprintf("the old course with id: %v was been removed and new course: %#v was add.\n", id, updatedCourse)
	json.NewEncoder(wtr).Encode(content)

}

func DeleteById(wtr http.ResponseWriter, req *http.Request) {

	wtr.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	courseToDel := params["courseid"] /* cspell:disable-line */

	if courseToDel == "" {
		msg := fmt.Sprintf("id: %v not found in DB.\n", courseToDel)
		fmt.Printf(msg)
		json.NewEncoder(wtr).Encode(msg)
		return

	}

	fmt.Printf("removing course by id: %v.\n", courseToDel)

	courseIndex := linearSearch(courseToDel, Courses...)
	if courseIndex == -1 {
		fmt.Printf("course id: %v not found in DB.\n", courseIndex)

		msg := "course by id " + string(courseIndex) + " not found"
		json.NewEncoder(wtr).Encode(msg)
		return

	}

	Courses = append(Courses[:courseIndex], Courses[courseIndex+1:]...)
	fmt.Printf("Courses deletion done.\nCourses: %#v\n", Courses)

	msg := fmt.Sprintf("course deleted, Courses now: %#v\n", Courses)
	json.NewEncoder(wtr).Encode(msg)

}
