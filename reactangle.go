package golang_united_school_homework

type Rectangle struct {
	Height, Weight float64
}

func (rectangle Rectangle) CalcPerimeter() float64 {
	if rectangle.Height < 0.0 || rectangle.Weight < 0.0 {
		return 0.0
	}

	return 2.0*rectangle.Height + 2.0*rectangle.Weight
}

func (rectangle Rectangle) CalcArea() float64 {
	if rectangle.Height < 0.0 || rectangle.Weight < 0.0 {
		return 0.0
	}

	return rectangle.Height * rectangle.Weight
}
