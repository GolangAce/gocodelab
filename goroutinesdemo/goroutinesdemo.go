package main

import (
	"fmt"
)

func printGreetings(source string) {

	for i := 0; i < 9; i++ {
		fmt.Println("Hello Gopher!", i, source)
	}
}

func main() {

	go printGreetings("goroutine")
	//time.Sleep(time.Millisecond * 135)
	printGreetings("main function")

}
