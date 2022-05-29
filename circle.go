package golang_united_school_homework

import "math"

type Circle struct {
	Radius float64
}

func (circle Circle) CalcPerimeter() float64 {
	if circle.Radius < 0.0 {
		return 0.0
	}

	return 2.0 * circle.Radius * math.Pi
}

func (circle Circle) CalcArea() float64 {
	if circle.Radius < 0.0 {
		return 0.0
	}

	return circle.Radius * circle.Radius * math.Pi
}
