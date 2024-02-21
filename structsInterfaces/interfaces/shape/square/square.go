package square

import "fmt"

// Square
type Square struct {
	sideLength float32
}

func NewSquare(l float32) *Square {
	fmt.Println("New Square created with Lenght: ", l)
	return &Square{sideLength: l}
}

func (s *Square) GetArea() float32 {
	return s.sideLength * s.sideLength
}
