package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("dealing files: ")

	content := "\nthis data need to be fed to a txt file.\nconst thx."
	path := "./mytext.txt"

	createFile(content, path)

}

func createFile(content string, path string) {

	// creating a file
	myFile, err := os.Create(path)
	if err != nil {
		fmt.Println("not good.")
	}

	// writing to the file.
	length, err := io.WriteString(myFile, content)
	if err != nil {
		fmt.Println("not good.")
	}

	fmt.Println("length is: ", length)

	defer myFile.Close()

	readFile(path)

}

func readFile(path string) {

	dataByte, err := ioutil.ReadFile(path)

	if err != nil {
		panic("something went wrong.")
	}

	fmt.Printf("raw data is: %v,\nparsed data is: %v\n", dataByte, string(dataByte))

}
