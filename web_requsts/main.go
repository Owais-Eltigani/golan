package main

import (
	"fmt"
	"io"
	"net/http"
)

const URL = "http://example.com/"

func main() {

	fmt.Println("these are go web requests: ")

	// starting the request.
	response, err := http.Get(URL)
	errCheck(err)

	fmt.Printf("the response type is %T,", response)
	// always close the connection.
	defer response.Body.Close()

	// reading the actual data.
	dataByte, err := io.ReadAll(response.Body) // ioutil.ReadAll(response.Body)
	errCheck(err)

	fmt.Printf("\n\n\n\the actual data is:\n %v", string(dataByte))
}

func errCheck(err error) {

	if err != nil {
		fmt.Println("something wnt wrong")
	}
}
