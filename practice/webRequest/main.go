package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Let's learn Web Request in Go Lang")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	// it will make a GET request to the URL
	// http.Get() function takes the URL as an argument and returns the response and error
	if err != nil {
		fmt.Println("Error while making GET request:", err)
		return
	}
	defer res.Body.Close() // it will close the response body
	fmt.Printf("Response type: %T\n", res)
	// fmt.Println("Response:", res)

	// read the response body
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while reading response body:", err)
		return
	}
	fmt.Println("Response body:", string(data)) // convert byte slice to string
}
