package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("this go server:")

	// generating the
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>hi<h1/>"))
	})

	log.Fatal(http.ListenAndServe(":3000", router))
}
