package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup //	pointer in usual.
var mutex sync.Mutex  //	pointer in usual.
var signals = []string{"test"}

func main() {

	// go print("foo")
	// print("hi")

	websites := []string{
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, website := range websites {
		go pingWebsite(website)
		// pingWebsite(website)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Printf("signals is: %v", signals)
}

func print(str string) {

	for i := 0; i < len(str); i++ {
		fmt.Println("message is: ", str)
		time.Sleep(3 * time.Second)
	}
}

func pingWebsite(url string) {

	defer wg.Done()
	result, err := http.Get(url)

	if err != nil {
		fmt.Printf("not available %v\n", url)
	} else {

		mutex.Lock()
		signals = append(signals, url)
		mutex.Unlock()

		fmt.Printf("%v success: %v\n", url, result.StatusCode)
	}
}
