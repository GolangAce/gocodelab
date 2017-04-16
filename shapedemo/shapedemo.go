package main

import (
	"fmt"
	"gocodelab/simpleshape"
)

func main() {

	r := simpleshape.NewRectangle(9, 6)
	t := simpleshape.NewTriangle(3, 6)

	fmt.Println("Area of myRectangle: ", simpleshape.Area(r))
	fmt.Println("Area of myCircle: ", simpleshape.Area(t))

}
