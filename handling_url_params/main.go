package main

import (
	"fmt"
	"net/url"
)

const URL = "https://example.go.com:3000/learn?lang=go&user=owais"

func main() {

	fmt.Println("welcome to handling go params.")

	// resolving the url
	resolve, _ := url.Parse(URL)

	fmt.Printf("the main thing is: %v\n", resolve)

	// meta data

	fmt.Println("resolve.Path: ", resolve.Path)
	fmt.Println("resolve.Scheme: ", resolve.Scheme)
	fmt.Println("resolve.RawPath: ", resolve.RawPath)
	fmt.Println("resolve.RawQuery: ", resolve.RawQuery)

	// extracting queries.

	params := resolve.Query()

	fmt.Printf("the language used to teach this course is: %v\n", params["lang"])

	// constructing the url

	url_infos := &url.URL{
		Scheme:   "https",
		Path:     "/css",
		RawQuery: "name=owais&age=23",
		Host:     "test.dev",
	}

	constructed_url := url_infos.String()

	fmt.Printf("\nurl is: %v\n", constructed_url)
}
