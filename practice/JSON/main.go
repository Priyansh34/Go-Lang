package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"Name"` // tag is used to specify the key name in JSON
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("Let's learn JSON in Go Lang")
	// create a struct
	person := Person{Name: "Priyansh", Age: 24, IsAdult: true}
	fmt.Println("Person details:", person)

	// JSON is a plain text string file where everything is in double quotes

	// convert person struct into JSON encoding (Marshalling)
	jsonData, err := json.Marshal(person) // it will convert the struct into JSON encoding
	if err != nil {
		fmt.Println("Error while marshalling JSON:", err)
		return
	}
	fmt.Println("JSON Data:", string(jsonData)) // convert byte slice to string

	// decoding JSON (Unmarshalling)
	var personData Person
	error := json.Unmarshal(jsonData, &personData) // it will convert JSON encoding into struct
	if error != nil {
		fmt.Println("Error while unmarshalling JSON:", error)
		return
	}
	fmt.Println("Person Data:", personData) // unmarshelling automtically converts byte slice to string

}
