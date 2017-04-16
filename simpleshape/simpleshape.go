package simpleshape

type Shape interface {
	Area() float64
}

func Area(s Shape) float64 {

	return s.Area()

}
