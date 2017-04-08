// In this example we use the greetingspackage
package main

import (
	"fmt"
	"gocodingisfun/greetingspackage"
)

func main() {

	greetingspackage.PrintGreetings()
	greetingspackage.GopherGreetings()

	fmt.Println("The value of the Magic Number is:", greetingspackage.MagicNumber)

}
