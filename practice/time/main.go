package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Let's start Time Package in Go lang")

	currentTime := time.Now()
	fmt.Println("Current Time is: ", currentTime)
	fmt.Printf("Type of current time %T\n", currentTime)

	// formatted time and date
	// we want date in dd-mm-yyyy format
	// for this we use a function Format
	// to add the format, we have to pass "02-01-2006" so it can return today's date in dd-mm-yyyy format

	formattedTime := currentTime.Format("02-01-2006")
	desireTime := currentTime.Format("2006-01-02")
	fmt.Println("Formatted date format:", formattedTime)
	fmt.Println("Desired date format", desireTime)

	// to get the Day, we have to pass Monday so it can return today's week day with time
	date := currentTime.Format("02-01-2006, Monday")
	fmt.Println("Date with day:", date)

	// to get the time, we need to pass time which followed by Go Lang
	// time format is 15:04:05

	timeDate := currentTime.Format("2006-01-02, 15:04:05, Monday")
	fmt.Println("Today's date, time, day:", timeDate)

	// to get the time in 12hours format
	// we should pass the format like this
	// 3:04 PM
	// here we can pass only PM not AM
	updatedTime := currentTime.Format("2006/01/02, 3:04:05 PM, Monday")
	fmt.Println("Time in 12 hour format:", updatedTime)

	// convert string into time
	// user input it datestring so we have to give the layout string in same format
	datestr := "2023-06-20" // if it is dd-mm-yyyy then we have to pass "02-01-2006"
	// if it is yyyy/mm/dd then we have to pass "2006/01/02"
	// here it is yyyy-mm-dd so we have passed "2006-01-02"
	layoutstr := "2006-01-02"
	convertedTime, _ := time.Parse(layoutstr, datestr)
	fmt.Println("Converted time is:", convertedTime)

	// add 1 more day in current time
	newDay := currentTime.Add(24 * time.Hour)
	fmt.Println("New day is:", newDay)
	formattedNewDay := newDay.Format("2006/01/02, 3:04:05 PM, Monday")
	fmt.Println("Formatted new day is:", formattedNewDay)

}
