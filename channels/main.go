package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("channels on go.")

	mychan := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	// only writing to channel
	go func(mychan <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()

		value, isValid := <-mychan
		fmt.Println("channel value: ", value, isValid)
		value2, isValid := <-mychan
		fmt.Println("channel value: ", value2, isValid)
	}(mychan, wg)

	// only to for the channel
	go func(mychan chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()

		mychan <- 3
		mychan <- 12
		close(mychan)
	}(mychan, wg)

	wg.Wait()
}
