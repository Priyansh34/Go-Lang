package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		fmt.Println("Let's learn File handling in Go Lang")
		fileName, err := os.Create("example.txt") // it will create a file
		if err != nil {
			fmt.Println("Error while creating file:", err)
			return
		}
		defer fileName.Close() // it will close the file
		// fmt.Println("File created successfully")
		// fmt.Println("File name is:", fileName.Name())

		// writing data into the file
		content := "Hello World by Priyansh Vashistha"
		_, error := io.WriteString(fileName, content+"\n") // it will write the content into the file
		// io.WriteString() function takes 2 arguments, first is file name and second is content which we want to write into the file
		// it returns 2 values, first is number of bytes written and second is error
		// if there is no error then it will return nil
		// io is used for input and output operations
		// here we are not using the number of bytes written so we have used _ to ignore it

		// byte, error := io.WriteString(fileName, content+"\n")
		// fmt.Println("Number of bytes written:", byte)
		if error != nil {
			fmt.Println("Error while writing into the file:", error)
			return
		}

		fmt.Println("File created and data written successfully")
	*/

	/*
		file, err := os.Open("example.txt") // it will open the file in read mode
		if err != nil {
			fmt.Println("Error while opening file:", err)
			return
		}
		defer file.Close()

		// create a buffer to read the file
		buffer := make([]byte, 1024) // it will create a buffer of 1024 bytes

		// read the file content into the buffer
		for {
			n, error := file.Read(buffer) // it will read the file content into the buffer
			if error == io.EOF {
				break // if we reach the end of the file then we will break the loop
			}
			if error != nil {
				fmt.Println("Error while reading file:", error)
				return
			}

			// process the read content
			fmt.Println(string(buffer[:n])) // it will print the file content
		}

	*/

	// read entire file content at once into a byte slice
	//content, err := ioutil.ReadFile("example.txt") // it will read the entire file content at once
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error while reading file:", err)
		return
	}
	fmt.Println(string(content))

}
