package main

import "math"

//Perimeter returns the perimeter of a rectangle given height and width
func Perimeter(rectangle Rectangle) (perimeter float64) {
	perimeter = 2 * (rectangle.Width + rectangle.Height)
	return
}

//Area returns area given passed height and width of rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

//Area returns area given passed height and width of rectangle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//Area returns area given passed height and width of rectangle
func (t Triangle) Area() float64 {
	return t.Height * t.Base / 2
}

func main() {

}

//Rectangle you fool
type Rectangle struct {
	Width  float64
	Height float64
}

//Circle you fool
type Circle struct {
	Radius float64
}

//Shape you fool
type Shape interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}
