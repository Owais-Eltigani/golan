package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	println("your name is ", name)
}
