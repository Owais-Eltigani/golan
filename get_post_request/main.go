package main

import (
	"fmt"
	"io"
	"net/http"
)

const URL = "http://localhost:3000"

func main() {

	fmt.Println("handling get requests: ")

	// making the request

	response, _ := http.Get(URL)
	defer response.Body.Close()

	data, _ := io.ReadAll(response.Body)

	fmt.Printf("data received is: %v\n", string(data), response.Header)

}
