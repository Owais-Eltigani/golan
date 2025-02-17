package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("welcome to simple web server using std lib")

	// creating a server and connecting it to the main route.
	server := http.FileServer(http.Dir(".")) //? change the name of doc of anything happened.
	http.Handle("/", server)
	PORT := ":3000"

	// server routes.
	http.HandleFunc("/hello", func(wtr http.ResponseWriter, req *http.Request) {

		// dealing with edge cases.

		// if req.URL.Path != "/hello" {
		// 	fmt.Println("wrong route {hello}")

		// 	http.Error(wtr, "page not found 404", http.StatusNotFound)
		// 	return
		// }

		if req.Method != "GET" {

			fmt.Println("bad method {hello, GET}")
			http.Error(wtr, "unsupported method", http.StatusNotFound)
			return
		}

		fmt.Println("visitor number one")

		content := "<h1>welcome to go server<h1/>"
		json.NewEncoder(wtr).Encode(content)

	})

	// form route

	http.HandleFunc("/form", func(wtr http.ResponseWriter, req *http.Request) {

		// form data.

		if err := req.ParseForm(); err != nil {

			fmt.Println("wrong while parsing form data.")
			http.Error(wtr, "error while parsing form.", http.StatusBadRequest)

			return
		}

		fmt.Println("parsing form.")

		name := req.FormValue("name")
		password := req.FormValue("password")

		fmt.Printf("name: %v\npassword: %v\n\n", name, password)

		myStruct := struct {
			UsrName string
			Pass    string
		}{
			UsrName: name,
			Pass:    password,
		}

		json.NewEncoder(wtr).Encode(myStruct)

	})

	// serving the content of the server.
	if err := http.ListenAndServe(PORT, nil); err != nil { //? if the server not worked try replacing server with nil.

		log.Fatal("something went wrong while starting the server")
	}
	fmt.Printf("server connected.\n")
}
