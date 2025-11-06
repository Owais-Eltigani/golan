package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type BOOK struct {
	Id          string
	Name        string
	AuthorName  string
	Price       string
	PublishDate string
}

//! =================================================

var DB *sql.DB

// homepage get all books
func GetAllBooks(wtr http.ResponseWriter, req *http.Request) {

	// deal with edge cases.
	if req.Method != "GET" {
		fmt.Println("bad gateway {books}")
		return
	}

	// getting all books
	fmt.Println("getting all books")

	// query
	query, err := DB.Query("SELECT B.book_id, B.book_name, A.name, B.price, B.publish_date FROM BOOKS B JOIN AUTHORS A USING(AUTHOR_ID)")

	if err != nil {
		fmt.Println("error while retreiving books")
		http.Error(wtr, "error while retreiving books", http.StatusNotFound)
		return
	}
	defer query.Close()

	var queryBooks []BOOK
	for query.Next() {

		var bookData BOOK
		if err := query.Scan(&bookData.Id, &bookData.Name, &bookData.AuthorName, &bookData.Price, &bookData.PublishDate); err != nil {
			fmt.Println("error while parsing books details")
			http.Error(wtr, "error while parsingg books details", http.StatusNotFound)
			return

		}

		queryBooks = append(queryBooks, bookData)
		fmt.Println("books retreived successfully.")
	}

	json.NewEncoder(wtr).Encode(queryBooks)

}

// add a new book
func AddBook(wtr http.ResponseWriter, req *http.Request) {

	// using post method
	if req.Method != "POST" {
		fmt.Println("wrong method type")
		wtr.Write([]byte("wrong method type"))
		return
	}

	// check body is not nil
	if req.Body == nil {

		fmt.Println("request body is nill", req.Body)
		wtr.Write([]byte("body is empty.\n"))
		return
	}

	// decoding the body
	var book BOOK
	json.NewDecoder(req.Body).Decode(&book)

	// check if the book is already in the DB
	if book.isEmpty(&book) {

		fmt.Println("one of book content is empty", book)
		wtr.Write([]byte("one of book attribute are empty"))
		return
	}

	// check is valid id [exist in DB]

	query, err := DB.Query("SELECT book_name FROM BOOKS WHERE book_id = " + book.Id)

	if err != nil {
		fmt.Println("error while getting one book.")
		wtr.Write([]byte("error while getting one book."))
		return
	}
	defer query.Close()

	for query.Next() {
		var isValid sql.NullString
		err := query.Scan(&isValid)

		if err != nil {
			fmt.Println("error is: ", err)
			return

		}
		if val, _ := isValid.Value(); val != "" {
			fmt.Println("book with id already exist.")
			wtr.Write([]byte("book with id already exist."))
			return
		}
	}
	myBook := struct {
		bookID      string
		bookName    string
		price       string
		publishDate string
		authorId    string
	}{
		bookID:      book.Id,
		bookName:    book.Name,
		price:       book.Price,
		publishDate: book.PublishDate,
		authorId:    book.AuthorName,
	}

	stmt, err := DB.Prepare("INSERT INTO BOOKS (book_id, book_name, price, publish_date, author_id) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println("error preparing statement:", err)
		http.Error(wtr, "error preparing statement", http.StatusInternalServerError)
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(myBook.bookID, myBook.bookName, myBook.price, myBook.publishDate, myBook.authorId)
	if err != nil {
		fmt.Println("error executing insert:", err)
		http.Error(wtr, "error inserting book", http.StatusInternalServerError)
		return
	}

	fmt.Println("book added", book)

	json.NewEncoder(wtr).Encode(book)

}

func DeleteById(wtr http.ResponseWriter, req *http.Request) {

	// check the url is correct
	if req.URL.Path != "/book/delete" {

		fmt.Println("wrong url")
		wtr.Write([]byte("wrong url"))
		return
	}

	// check method
	if req.Method != "DELETE" {
		fmt.Println("wrong method")
		wtr.Write([]byte("wrong method"))
		return
	}

	// check id is not empty
	id := req.URL.Query().Get("id")
	if id == "" {
		fmt.Println("no id")
		wtr.Write([]byte("no id provided."))
		return
	}

	// check is valid id [exist in DB]
	query, err := DB.Query("SELECT book_name FROM BOOKS where book_id = " + id)

	if err != nil {
		fmt.Println("error while getting book info.")
		wtr.Write([]byte("error while getting book info."))
		return
	}
	defer query.Close()

	var isValid sql.NullString

	if query.Next() {
		err := query.Scan(&isValid)

		if err != nil {
			fmt.Println("error is: ", err)
			return

		}
	}

	if !isValid.Valid {
		fmt.Println("book with doesn't exist.", isValid)
		wtr.Write([]byte("book with id already exist."))
		return

	}

	deleteQuery, _ := DB.Query("DELETE FROM BOOKS WHERE book_id = " + id)
	defer deleteQuery.Close()

	fmt.Println("book is deleted", isValid.Valid)
	wtr.Write([]byte("book is deleted"))

}

func UpdateById(wtr http.ResponseWriter, req *http.Request) {

	// check the url is correct
	if req.URL.Path != "/book/update" {

		fmt.Println("wrong url")
		wtr.Write([]byte("bad gate way"))
		return
	}

	// check method
	if req.Method != "PUT" {
		fmt.Println("wrong method")
		wtr.Write([]byte("wrong method"))
		return
	}

	// check id is not empty
	id := req.URL.Query().Get("id")
	if id == "" {
		fmt.Println("no id")
		wtr.Write([]byte("no id provided."))
		return
	}

	// check body is not nil
	if req.Body == nil {

		fmt.Println("request body is nill", req.Body)
		wtr.Write([]byte("body is empty.\n"))
		return
	}

	// decoding the body
	var book BOOK
	json.NewDecoder(req.Body).Decode(&book)

	// check if the book is already in the DB
	if book.isEmpty(&book) {

		fmt.Println("one of book content is empty", book)
		wtr.Write([]byte("one of book attribute are empty"))
		return
	}

	// check is valid id [exist in DB]
	query, err := DB.Query("SELECT book_name FROM BOOKS where book_id = " + id)

	if err != nil {
		fmt.Println("error while getting book info.")
		wtr.Write([]byte("error while getting book info."))
		return
	}
	defer query.Close()

	var isValid sql.NullString

	if query.Next() {
		err := query.Scan(&isValid)

		if err != nil {
			fmt.Println("error is: ", err)
			return

		}
	}

	if !isValid.Valid {
		fmt.Println("book with id doesn't exist.", isValid)
		wtr.Write([]byte("book with id doesn't exist."))
		return
	}

	// Prepare update statement
	stmt, err := DB.Prepare("UPDATE BOOKS SET book_name=?, price=?, publish_date=?, author_id=? ,book_id=? WHERE book_id=?")
	if err != nil {
		fmt.Println("error preparing update statement:", err)
		http.Error(wtr, "error preparing update statement", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	// Execute the update
	_, err = stmt.Exec(book.Name, book.Price, book.PublishDate, book.AuthorName, book.Id, id)
	if err != nil {
		fmt.Println("error executing update:", err)
		http.Error(wtr, "error updating book", http.StatusInternalServerError)
		return
	}

	fmt.Println("book updated successfully")
	json.NewEncoder(wtr).Encode(book)

}

func GetBookById(wtr http.ResponseWriter, req *http.Request) {

	if req.Method != "GET" {
		fmt.Println("bad gateway {books}")
		return
	}

	// check method
	if req.Method != "GET" {
		fmt.Println("wrong method")
		wtr.Write([]byte("wrong method"))
		return
	}

	// check id is not empty
	id := req.URL.Query().Get("id")
	if id == "" {
		fmt.Println("no id")
		wtr.Write([]byte("no id provided."))
		return
	}

	// check is valid id [exist in DB]
	query, err := DB.Query("SELECT B.book_id, B.book_name, B.price, B.publish_date, A.name FROM BOOKS B JOIN AUTHORS A ON B.book_id = A.author_id  WHERE book_id = " + id)

	if err != nil {
		fmt.Println("error while getting book info.\n", err)
		wtr.Write([]byte("error while getting book info."))
		return
	}
	defer query.Close()

	var book BOOK

	for query.Next() {
		err := query.Scan(&book.Id, &book.Name, &book.Price, &book.PublishDate, &book.AuthorName)

		if err != nil {
			fmt.Println("error while parsing the file: ", err)
			return

		}

	}

	fmt.Println("book was fetched deleted")
	json.NewEncoder(wtr).Encode(book)

}

//? ============== helper functions =====================

func (b *BOOK) isEmpty(book *BOOK) bool {

	return book.Id == "" || book.AuthorName == ""
}

func DBconnect() *sql.DB {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/library")

	if err != nil {
		fmt.Println("couldn't connect to db.")
		log.Fatal("couldn't connect to db.")
	}

	fmt.Println("db connected successfully. ")
	return db
}
