package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Let's start Data Concersion in GOLang")
	var num int = 42
	fmt.Println("The value of num is:", num)
	fmt.Printf("Type of num is: %T\n", num)

	var data float64 = float64(num)
	fmt.Println("Data is:", data)
	fmt.Printf("Typr of data is: %T\n", data)

	// convert integer into a string
	number := 1234
	str := strconv.Itoa(number)
	fmt.Println("converted string is:", str)
	fmt.Printf("Typr of %q is: %T\n", str, str)

	// converting string into a integer
	number_str := "342001"
	number_int, _ := strconv.Atoi(number_str)
	fmt.Println("converted integer is:", number_int)
	fmt.Printf("Type of data is: %T\n", number_int)

	// converting float string into a float
	num_str := "34.2001"
	//number_float, _ := strconv.Atoi(num_str) //Atoi function convert string into integer if we use it for float data, it will pass 0 as nil
	// so for converting string into float, we use ParseFloat function
	number_float, _ := strconv.ParseFloat(num_str, 64)
	// ParseFloat function take 2 vlues, one is string and other is bitsize for float
	fmt.Println("converted integer is:", number_float)
	fmt.Printf("Type of data is: %T\n", number_float)

}
