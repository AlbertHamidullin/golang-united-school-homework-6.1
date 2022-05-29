package golang_united_school_homework

import "math"

type Triangle struct {
	Side float64
}

func (triangle Triangle) CalcPerimeter() float64 {
	if triangle.Side < 0.0 {
		return 0.0
	}

	return 3.0 * triangle.Side
}

func (triangle Triangle) CalcArea() float64 {
	if triangle.Side < 0.0 {
		return 0.0
	}

	return triangle.Side * triangle.Side * math.Sqrt(3.0) / 4.0
}
