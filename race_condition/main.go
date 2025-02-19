package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("race condition: ")

	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	score := []int{0}
	wg.Add(3)

	go func(wg *sync.WaitGroup, mutex *sync.Mutex) {
		defer wg.Done()

		fmt.Printf("func1 \n")
		mutex.Lock()
		score = append(score, 1)
		mutex.Unlock()
	}(wg, mutex)

	go func(wg *sync.WaitGroup, mutex *sync.Mutex) {
		defer wg.Done()

		fmt.Printf("func2 \n")
		mutex.Lock()
		score = append(score, 2)
		mutex.Unlock()
	}(wg, mutex)

	go func(wg *sync.WaitGroup, mutex *sync.Mutex) {

		defer wg.Done()

		fmt.Printf("func3 \n")
		mutex.Lock()
		score = append(score, 3)
		mutex.Unlock()
	}(wg, mutex)

	wg.Wait()
	fmt.Printf("score: %v", score)
}

/* val, err := something()

if !err {
	fmt.Println("err is", err)
}

Printf("val is", val)
*/
