package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("CRUD interface using std lib")

	// create the server
	// server := http.FileServer(http.Dir("."))
	// http.Handle("/", server)

	if err := godotenv.Load(".env"); err != nil {

		fmt.Println("couldn't load the the env file", err)
		log.Fatal("load env failed.\n")
		return

	}

	http.HandleFunc("/books", GetAllBooks)
	http.HandleFunc("/book", GetBookById)
	http.HandleFunc("/book/add-book", AddBook)
	http.HandleFunc("/book/delete", DeleteById) // Access ID using r.URL.Query().Get("id")
	http.HandleFunc("/book/update", UpdateById) // Access ID using r.URL.Query().Get("id")

	if err := http.ListenAndServe(os.Getenv("PORT"), nil); err != nil {

		// check if everything is ok
		fmt.Println("server is not working", err)
		log.Fatal("server is not working")
		return
	}
	fmt.Println("server up and running.")

}
