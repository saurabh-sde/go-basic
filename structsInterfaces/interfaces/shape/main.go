package main

import (
	"fmt"

	"github.com/saurabh-sde/go-basic/structsInterfaces/interfaces/shape/square"
	"github.com/saurabh-sde/go-basic/structsInterfaces/interfaces/shape/triangle"
)

type Shape interface {
	GetArea() float32
}

func GetArea(s Shape) float32 {
	fmt.Printf("Area: %+v\n", s.GetArea())
	return s.GetArea()
}

func main() {
	sq := square.NewSquare(5)
	GetArea(sq)
	t := triangle.NewTriangle(5, 2)
	GetArea(t)

}
