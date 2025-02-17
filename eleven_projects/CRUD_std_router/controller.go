package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Book struct {
	Title         string `json:"title"`
	Id            string `json:"id"`
	NumberOfPages int32  `json:"numberOfPages, omitempty"`
	Author        *Author
}

type Author struct {
	Name     string `json:"name"`
	AuthorId string `json:"authorid"`
	Books    []Book
}

var books []Book

// homepage get all books
func GetAllBooks(wtr http.ResponseWriter, req *http.Request) {

	// deal with edge cases.

	if req.Method != "GET" {
		fmt.Println("bad gateway {books}")
		return
	}

	// getting all books
	fmt.Println("getting all books")

	json.NewEncoder(wtr).Encode(books)

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
	var book Book
	json.NewDecoder(req.Body).Decode(&book)

	// check if the book is already in the DB
	if book.isEmpty(&book) {

		fmt.Println("one of book content is empty", book)
		wtr.Write([]byte("one of book attribute are empty"))
		return
	}

	// check is valid id [exist in DB]
	index := linearSearch(book.Id)
	if index != -1 {
		fmt.Println("book already exist")
		wtr.Write([]byte("book already exist"))
		return
	}

	books = append(books, book)
	fmt.Println("book added", book)

	json.NewEncoder(wtr).Encode(book)

}

func DeleteById(wtr http.ResponseWriter, req *http.Request) {

	// check the url is correct
	if req.URL.Path != "/book/delete" {

		fmt.Println("wrong url")
		wtr.Write([]byte("bad gate way"))
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

	index := linearSearch(id)
	if index == -1 {
		fmt.Println("not valid id")
		wtr.Write([]byte("not valid id"))
		return
	}

	books = append(books[:index], books[index+1:]...)
	fmt.Println("book deleted")
	json.NewEncoder(wtr).Encode(books)

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
	var book Book
	json.NewDecoder(req.Body).Decode(&book)

	// check if the book is already in the DB
	if book.isEmpty(&book) {

		fmt.Println("one of book content is empty", book)
		wtr.Write([]byte("one of book attribute are empty"))
		return
	}

	// check is valid id [exist in DB]

	index := linearSearch(id)
	if index == -1 {
		fmt.Println("not valid id")
		wtr.Write([]byte("not valid id"))
		return
	}

	books[index] = book
	fmt.Println("book updated")
	json.NewEncoder(wtr).Encode(books)

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

	index := linearSearch(id)
	if index == -1 {
		fmt.Println("not valid id")
		wtr.Write([]byte("not valid id"))
		return
	}

	json.NewEncoder(wtr).Encode(books[index])
	fmt.Println("book retrieved.")

}

//? ============== helper functions =====================

// auto initialize the the slice of
func init() {
	books = []Book{
		{Title: "The Great Gatsby", Id: "1", NumberOfPages: 180, Author: &Author{Name: "F. Scott Fitzgerald", AuthorId: "A1"}},
		{Title: "1984", Id: "2", NumberOfPages: 328, Author: &Author{Name: "George Orwell", AuthorId: "A2"}},
		{Title: "To Kill a Mockingbird", Id: "3", NumberOfPages: 281, Author: &Author{Name: "Harper Lee", AuthorId: "A3"}},
		{Title: "Pride and Prejudice", Id: "4", NumberOfPages: 432, Author: &Author{Name: "Jane Austen", AuthorId: "A4"}},
		{Title: "The Catcher in the Rye", Id: "5", NumberOfPages: 234, Author: &Author{Name: "J.D. Salinger", AuthorId: "A5"}},
		{Title: "Lord of the Flies", Id: "6", NumberOfPages: 224, Author: &Author{Name: "William Golding", AuthorId: "A6"}},
		{Title: "The Hobbit", Id: "7", NumberOfPages: 310, Author: &Author{Name: "J.R.R. Tolkien", AuthorId: "A7"}},
		{Title: "Fahrenheit 451", Id: "8", NumberOfPages: 249, Author: &Author{Name: "Ray Bradbury", AuthorId: "A8"}},
		{Title: "Animal Farm", Id: "9", NumberOfPages: 141, Author: &Author{Name: "George Orwell", AuthorId: "A2"}},
		{Title: "Brave New World", Id: "10", NumberOfPages: 311, Author: &Author{Name: "Aldous Huxley", AuthorId: "A9"}},
	}
}

func (b *Book) isEmpty(book *Book) bool {

	return book.Id == "" || book.Author.Name == ""
}

func linearSearch(id string) int {

	for index, book := range books {

		if book.Id == id {
			return index
		}
	}

	return -1

}
