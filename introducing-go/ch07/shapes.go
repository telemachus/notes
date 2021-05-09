package geometry

import "math"

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	l, w float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r *Rectangle) area() float64 {
	return r.l * r.w
}

func (r *Rectangle) perimeter() float64 {
	return 2*r.l + 2*r.w
}
