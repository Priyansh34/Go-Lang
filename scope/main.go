/*
package main

import (
	"fmt"
	"myapp/packageone"
)

var one = "one" //package level variable, can be used in any function within this package

func main() {
	one := "this is a block level variable" //block level variable, can only be used within this(main) function
	fmt.Println(one)
	myFunc()

	newString := packageone.PublicVar
	fmt.Println(newString)

	packageone.ExportedFunc()

}

func myFunc() {
	fmt.Println(one)
}
*/

//Challange:
// variables
//declare a package level variable for the main
//package maned myVar

//declare a block level variable for the main function named blockVar

//declare a pakg level variable in the packageone named PackageVar

//create an exported function in packageone called PrintMe

//in the main func, print out the values for myVar, blockVar and PackageVar on one line
//using the PrintMe function in packageone

package main

import (
	"fmt"
	"myapp/packageone"
)

var myVar = "I am a package level Variable"

func main() {
	var blockVar = "I am a block level variable"

	packageone.PrintMe(myVar, blockVar)

	var Myname = "Priyansh"
	fmt.Println("My name is", Myname)

}
