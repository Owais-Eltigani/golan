package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("this is a GO web server running... ")

	//
	router := mux.NewRouter()
	PORT := ":5000"

	router.HandleFunc("/", ServHome).Methods("GET")                          /* cspell:disable-line */
	router.HandleFunc("/courses", GetAllCourses).Methods("GET")              /* cspell:disable-line */
	router.HandleFunc("/course/{courseid}", GetCourseById).Methods("GET")    /* cspell:disable-line */
	router.HandleFunc("/course", AddNewCourse).Methods("POST")               /* cspell:disable-line */
	router.HandleFunc("/course/{courseid}", UpdateCourseById).Methods("PUT") /* cspell:disable-line */
	router.HandleFunc("/course/{courseid}", DeleteById).Methods("DELETE")    /* cspell:disable-line */

	// server Listener
	log.Fatal(http.ListenAndServe(PORT, router))
}
