package main

import "fmt"

func main() {

	var subject string = "Gophers"

	fmt.Println("First element of Gophers string: ", string("Gophers"[0]))
	fmt.Printf("The first value of the subject string: %v\n", string(subject[0]))
	fmt.Printf("The last value of the subject string: %v\n", string(subject[len(subject)-1]))
	fmt.Println("Hello " + subject + "!")

}
