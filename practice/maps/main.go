package main

import "fmt"

func main() {
	fmt.Println("Let's learn Maps in Go Lang")

	studentsGrades := make(map[string]int)

	studentsGrades["Priyansh"] = 100
	studentsGrades["Bhavesh"] = 80
	studentsGrades["Ayush"] = 90

	fmt.Println("Marks of ayush is:", studentsGrades["Ayush"])

	// If key is not present in data, it will it's value as zero/default
	fmt.Println("Rohan's marks is:", studentsGrades["Rohan"])

	// Checking if a key exists
	grade, exists := studentsGrades["Rohan"]
	fmt.Println("Grade of Rohan:", grade)
	fmt.Println("Rohan exists:", exists)

	// Checking if a key exists
	marks, Exists := studentsGrades["Priyansh"]
	fmt.Println("Grade of Priyansh:", marks)
	fmt.Println("Priyansh exists:", Exists)

	for index, value := range studentsGrades {
		fmt.Printf("Student %s has %d marks.\n", index, value)
	}

	// students := make(map[int]string)

	// students['1'] = "Priyansh"
	// students['5'] = "Bhavesh"
	// students['3'] = "Saurabh"

	// fmt.Println("Roll no. 3 student is:", students['3'])

	// delete(students, '1') // it will delete the value which is stored in key 1 means Priyansh
	// fmt.Printf("current value of key 1 is: %q\n", students['1'])
	// students['1'] = "Ashwin"
	// fmt.Println("updated value of key 1 is: ", students['1'])

	// fmt.Printf("Roll no. 6 student is: %q\n", students['6'])

	// another way to declare map
	person := map[string]int{
		"A": 40,
		"B": 20,
		"C": 69,
	}

	for index, value := range person {
		fmt.Printf("Roll No of %s is %d.\n", index, value)
	}
}
