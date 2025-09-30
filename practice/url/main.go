package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Let's learn URL handling in Go Lang")
	myURL := "https://example.com:8080/path/to/resource?key1=value1&key2=value2"
	fmt.Println("URL is:", myURL)
	fmt.Printf("Tupe of myURL is: %T\n", myURL)

	// parsing the URL
	parsedUrl, err := url.Parse(myURL) // it will parse the URL and return a URL struct
	if err != nil {
		fmt.Println("Error while parsing URL:", err)
		return
	}
	fmt.Printf("Type of parsedUrl is: %T\n", parsedUrl)
	// we can access the different parts of the URL using the parsedUrl variable
	fmt.Println("Scheme:", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("RawQuery:", parsedUrl.RawQuery)

	// Modify the url
	parsedUrl.Path = "newpath"
	parsedUrl.RawQuery = "iampriyansh-vashistha"
	parsedUrl.Host = "mywebsite.com"

	fmt.Println("Modified UR: ", parsedUrl.String())
}
