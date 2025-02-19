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

	server := http.NewServeMux()

	// routes
	server.HandleFunc("GET /books", GetAllBooks)
	server.HandleFunc("POST /book/add-book", AddBook)
	server.HandleFunc("DELETE /book/delete", DeleteById) // Access ID using r.URL.Query().Get("id")
	server.HandleFunc("PUT /book/update", UpdateById)    // Access ID using r.URL.Query().Get("id")
	server.HandleFunc("GET /book", GetBookById)

	if err := http.ListenAndServe(os.Getenv("PORT"), server); err != nil {

		// check if everything is ok
		fmt.Println("server is not working", err)
		log.Fatal("server is not working")
		return
	}
	fmt.Println("server up and running.")

}
