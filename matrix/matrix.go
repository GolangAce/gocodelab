package main

import (
	"fmt"
)

func main() {

	// Here we create a multi-dimensional array, a 4x3 matrix
	myMatrix := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}

	// First attempt at printing the matrix
	fmt.Printf("%+v\n", myMatrix)
	fmt.Println("\n")

	// A better looking matrix output to the screen
	printFourByThreeMatrix(myMatrix)
}

// A function to print the matrix in a more pretty manner
func printFourByThreeMatrix(inputMatrix [3][4]int) {

	width := len(inputMatrix)
	height := len(inputMatrix[0])

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			fmt.Printf("%5d ", inputMatrix[i][j])
		}
		fmt.Println()
	}

}
