package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("CRUD interface using std lib")

	// create the server
	// server := http.FileServer(http.Dir("."))
	// http.Handle("/", server)
	// reading .env to load creds.

	if err := godotenv.Load(".env"); err != nil {

		fmt.Println("couldn't load the the env file", err)
		log.Fatal("load env failed.\n")
		return

	}

	//
	DB = DBconnect()

	defer DB.Close()

	// routes
	http.HandleFunc("/books", GetAllBooks)
	http.HandleFunc("/book/add-book", AddBook)
	http.HandleFunc("/book/delete", DeleteById) // Access ID using r.URL.Query().Get("id")
	http.HandleFunc("/book/update", UpdateById) // Access ID using r.URL.Query().Get("id")
	http.HandleFunc("/book", GetBookById)

	if err := http.ListenAndServe(os.Getenv("PORT"), nil); err != nil {

		// check if everything is ok
		fmt.Println("server is not working", err)
		log.Fatal("server is not working")
		return
	}
	fmt.Println("server up and running.")

}
