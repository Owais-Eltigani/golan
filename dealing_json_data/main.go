package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	fmt.Println("dealing with json data")

	// creating list of objs
	items := []Product{
		{"p1", 1},
		{"p2", 2},
	}

	// fmt.Printf("stringified data is: %s", encodeJson(items...))
	fmt.Printf("encoded data is: %s", decodeJson(items...))

	// sometimes you just want to save data as key-value pair
	// for that we use a map.

	// myMap := make(map[string]interface{}) because we dont know
	// the type of data we use interface.
	// json.Unmarshal(data_from_web,&myMap)

}

func encodeJson(items ...Product) string {
	stringified_data, err := json.Marshal(items)

	if err != nil {
		fmt.Printf("ops")
	}
	// fmt.Printf("data is: %s", (stringified_data))

	return string(stringified_data)
}

func decodeJson(items ...Product) string {

	// calling the encode method to gives some encoded data to
	// decoded.

	encoded_data := encodeJson(items...)

	temp := []Product{} //? a vassal to store in resolved data

	data := []byte(encoded_data) //? because data comes as byte from the web.

	isValidJson := json.Valid(data)

	if !isValidJson {
		fmt.Printf("something went wrong\n")
	}

	fmt.Printf("the json is valid and your data is:\n%v\n%#v", encoded_data, json.Unmarshal(data, &temp))
	return string(data)

}

type Product struct {
	Name string
	Id   int
}
