package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("CRUD interface using std lib")

	// create the server
	// server := http.FileServer(http.Dir("."))
	// http.Handle("/", server)

	// if err := godotenv.Load(".env"); err != nil {

	// 	fmt.Println("couldn't load the the env file", err)
	// 	log.Fatal("load env failed.\n")
	// 	return

	// }

	// http.HandleFunc("/books", GetAllBooks)
	// http.HandleFunc("/book", GetBookById)
	// http.HandleFunc("/book/add-book", AddBook)
	// http.HandleFunc("/book/delete", DeleteById) // Access ID using r.URL.Query().Get("id")
	// http.HandleFunc("/book/update", UpdateById) // Access ID using r.URL.Query().Get("id")

	// if err := http.ListenAndServe(os.Getenv("PORT"), nil); err != nil {

	// 	// check if everything is ok
	// 	fmt.Println("server is not working", err)
	// 	log.Fatal("server is not working")
	// 	return
	// }
	// fmt.Println("server up and running.")

	db := DBconnect()
	query, err := db.Query("SELECT customer_id, first_name, last_name, birth_date, phone FROM CUSTOMERs")

	if err != nil {
		log.Fatal("invalid query:", err)
	}
	defer query.Close() // Close the rows when done

	for query.Next() {
		var customerID int
		var firstName, lastName, birth_date, phone sql.NullString

		err = query.Scan(&customerID, &firstName, &lastName, &birth_date, &phone)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}
		fmt.Printf("ID: %d, Name: %s %s, Email: %s, phone: %s\n",
			customerID,
			firstName.String,
			lastName.String,
			birth_date.String,
			phone.String)
	}

	if err = query.Err(); err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}

func DBconnect() *sql.DB {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/store")

	if err != nil {
		fmt.Println("couldn't connect to db.")
		log.Fatal("couldn't connect to db.")
	}

	fmt.Println("db connected successfully. ")
	return db
}
