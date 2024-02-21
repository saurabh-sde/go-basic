package main

import "fmt"

type Shape interface {
	getArea() float32
}

func getArea(s Shape) float32 {
	fmt.Printf("Area: %+v\n", s.getArea())
	return s.getArea()
}

// Square
type Square struct {
	sideLength float32
}

func NewSquare(l float32) *Square {
	fmt.Println("New Square created with Lenght: ", l)
	return &Square{sideLength: l}
}

func (s *Square) getArea() float32 {
	return s.sideLength * s.sideLength
}

type Triangle struct {
	baseLength float32
	height     float32
}

func NewTriangle(l, h float32) *Triangle {
	fmt.Println("New Triangle created with Lenght: ", l, " Height: ", h)
	return &Triangle{baseLength: l, height: h}
}

func (s *Triangle) getArea() float32 {
	return s.height * s.baseLength * 0.5
}

func main() {
	sq := NewSquare(5)
	getArea(sq)
	t := NewTriangle(5, 2)
	getArea(t)

}
