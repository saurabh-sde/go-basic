package main

import (
	"fmt"
)

type Rectangle struct {
	width  int
	height int
}

// constuctor
func NewRectangle(width, height int) *Rectangle {
	return &Rectangle{width: width, height: height}
}

func (r *Rectangle) getArea() int {
	return r.width * r.height
}

func (r *Rectangle) getPerimeter() int {
	return 2*r.width + 2*r.height
}

func (r *Rectangle) mutiplyDimesions(factor int) {
	r.height *= factor
	r.width *= factor
}

func main() {
	r := NewRectangle(10, 5)
	fmt.Printf("Rectangle: %+v\n", *r)
	// Here we call the 2 methods defined for our struct.
	// NOTE:
	// 1. Only Rectangle type can call to defined functions
	// 2. Using pointers to structs instead of structs allows passing to func by reference instead of value so any modifications in func body will reflect onto original structs data.
	// Ex - check func mutiplyDimesions for details
	fmt.Println("getArea: ", r.getArea())
	fmt.Println("getPerimeter:", r.getPerimeter())
	r.mutiplyDimesions(2)
	fmt.Printf("Modified Rectangle: %+v\n", *r)
	fmt.Println("Modified Rectangle: getArea: ", r.getArea())
	fmt.Println("Modified Rectangle: getPerimeter:", r.getPerimeter())
}
