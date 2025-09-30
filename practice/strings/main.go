package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Let's learn Strings package in Go Lang")
	data := "apple,mango,orange"
	parts := strings.Split(data, ",") // strings.Split() func convert the data into slice.
	//Split slices s into all substrings separated by sep and returns a slice of the substrings between those separators
	fmt.Println(parts)

	// checking how many time a substring repeats
	// func strings.Count(s string, substr string) int
	// this function take string and the substring which we want to check the count and returns the int value
	str := "one two four eight two two nine"
	count := strings.Count(str, "two")
	fmt.Println("count of Two is:", count)

	// Trim function is used to remvove black space from start and the end of our string
	// it takes data from first charactor till last
	// this function don't remove
	str = "      Priyansh    vashsitha      "
	trimmed := strings.TrimSpace(str)
	fmt.Println("Trimmed:", trimmed)

	// join strings
	// to join strings, we use Join function in Strings
	// Join function takes the strings input & saperator value and returns string value
	// func strings.Join(elems []string, sep string) string
	str1 := "Priyansh"
	str2 := "Vashistha"
	fullname := strings.Join([]string{str1, str2}, " ")
	addingName := strings.Join([]string{str1, "Kumar", str2}, " ")
	fmt.Println("Fullname: ", fullname)
	fmt.Println("Adding middle name:", addingName)

	// Compare function
	// func strings.Compare(a string, b string) int
	// Compare returns an integer comparing two strings lexicographically. The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
	name := strings.Compare(str1, str2)
	fmt.Println("comparision: ", name)

}
