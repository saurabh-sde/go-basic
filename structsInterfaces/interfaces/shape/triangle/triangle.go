package triangle

import "fmt"

type Triangle struct {
	baseLength float32
	height     float32
}

func NewTriangle(l, h float32) *Triangle {
	fmt.Println("New Triangle created with Lenght: ", l, " Height: ", h)
	return &Triangle{baseLength: l, height: h}
}

func (s *Triangle) GetArea() float32 {
	return s.height * s.baseLength * 0.5
}
