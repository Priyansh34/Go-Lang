/*
package packageone

var privateVar = "I am private" // unexported variable, accessible only within this package

var PublicVar = "I am public or exported" // exported variable, accessible from other packages

func notExported() {
	// unexported function, can only be accessed within this package
}

func ExportedFunc() {
	// exported function, can be accessed from other packages
}
*/

package packageone

import (
	"fmt"
)

var PackageVar = "I am a package level variable in packageone"

func PrintMe(str1, str2 string) {
	fmt.Println(str1, str2, PackageVar)
}
