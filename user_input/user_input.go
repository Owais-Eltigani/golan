package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	fmt.Println("enter your name: ")

	reader := bufio.NewReader(os.Stdin)

	// comma ok , err ok {it's act as try and catch}
	name, _ := reader.ReadString('\n')

	println("your name is ", name)
}