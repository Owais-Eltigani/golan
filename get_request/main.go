package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const URL = "http://localhost:3000/ingredients"

func main() {

	fmt.Println("handling get requests: ")

	// making the request and reading the data.
	// getData(URL)
	postData(URL)

}

func getData(url string) {

	response, err := http.Get(url)
	errHandling(err)
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	errHandling(err)

	fmt.Printf("the data is: %v\n\n", string(data))

}

func postData(url string) {

	//TODO more than one value to post
	// if len(values) > 1 {

	// 	for _, val := range(values) {

	// 	}

	// }

	// stringifying the data.
	content := strings.NewReader(`
	{
	"id":"110",
	"text":"test2"
	}
	`)

	// sending the request.
	response, err := http.Post(url, "application/json", content)
	errHandling(err)
	defer response.Body.Close()

	// reading the returned values.
	data, err := io.ReadAll(response.Body)
	errHandling(err)

	fmt.Printf("data is: %v", string(data))
}

func postFormData(target_url string) {

	// creating a map to store the values of the form.
	// if use formpost or normal post they do the same
	// job so far but there maybe difference .

	data := url.Values{}

	data.Add("owais", "23")

	response, err := http.PostForm(target_url, data)
	errHandling(err)
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	errHandling(err)

	fmt.Printf("content is: %v", string(content))
}

func errHandling(err error) {

	if err != nil {
		fmt.Printf("something went bad\n")
	}

}
