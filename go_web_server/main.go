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

	router.HandleFunc("/", servHome).Methods("GET")                          /* cspell:disable-line */
	router.HandleFunc("/courses", getAllCourses).Methods("GET")              /* cspell:disable-line */
	router.HandleFunc("/course/{courseid}", getCourseById).Methods("GET")    /* cspell:disable-line */
	router.HandleFunc("/course", addNewCourse).Methods("POST")               /* cspell:disable-line */
	router.HandleFunc("/course/{courseid}", updateCourseById).Methods("PUT") /* cspell:disable-line */
	router.HandleFunc("/course/{courseid}", deleteById).Methods("DELETE")    /* cspell:disable-line */

	// server Listener
	log.Fatal(http.ListenAndServe(PORT, router))
}
