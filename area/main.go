package main

import "fmt"

type area interface {
	Area() float64
}

type circle struct {
	pi     float64
	radius float64
}
type sqaure struct {
	length float64
}
type rectangle struct {
	length float64
	width  float64
}

func (c circle) Area() float64 {

	return c.pi * c.radius * c.radius
}
func (s sqaure) Area() float64 {

	return s.length * s.length
}
func (r rectangle) Area() float64 {

	return r.length * r.width
}
func main() {

	a1 := circle{3.14, 9}
	a2 := sqaure{25}
	a3 := rectangle{23, 13.5}
	var a11 area

	a11 = a1
	fmt.Println("Area of circle is : ", a11.Area())
	a11 = a2
	fmt.Println("Area of square is : ", a11.Area())
	a11 = a3
	fmt.Println("Area of rectangle is : ", a11.Area())

}
