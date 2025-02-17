package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("this go server is running:")

	// generating the
	router := mux.NewRouter()
	PORT := ":3000"

	// w is http response writer || r is request from http
	router.HandleFunc("/", landingPage)

	log.Fatal(http.ListenAndServe(PORT, router))
}

func landingPage(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read body", http.StatusBadRequest)
		return
	}
	user_msg := string(body)
	// content := "<h1>welcome to golang<h1/> <br>your message is " + user_msg
	w.Write([]byte(user_msg))

}
