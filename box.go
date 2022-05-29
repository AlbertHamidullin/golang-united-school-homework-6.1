package golang_united_school_homework

import (
	"errors"
	"fmt"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

var (
	errorEqualsNil          = errors.New("equals nil")
	errorCapacityExceeded   = errors.New("capacity exceeded")
	errorIndexOutOfRange    = errors.New("index out of range")
	errorCirclesAreNotExist = errors.New("circles are not exist")
)

const (
	errorTemplate = "box error: %w"
)

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	if shapesCapacity < 0 {
		return nil
	}

	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if b == nil {
		return fmt.Errorf(errorTemplate, errorEqualsNil)
	}
	if len(b.shapes) >= b.shapesCapacity {
		return fmt.Errorf(errorTemplate, errorCapacityExceeded)
	}

	b.shapes = append(b.shapes, shape)

	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if b == nil {
		return nil, fmt.Errorf(errorTemplate, errorEqualsNil)
	}
	if !(i >= 0 && i <= len(b.shapes)-1) {
		return nil, fmt.Errorf(errorTemplate, errorIndexOutOfRange)
	}

	return b.shapes[i], nil
}

func (b *box) removeByIndex(i int) {
	copy(b.shapes[i:], b.shapes[i+1:])
	b.shapes[len(b.shapes)-1] = nil
	b.shapes = b.shapes[:len(b.shapes)-1]
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	shape, err := b.GetByIndex(i)

	if err != nil {
		return nil, err
	}

	b.removeByIndex(i)

	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	oldShape, err := b.GetByIndex(i)

	if err != nil {
		return nil, err
	}

	b.shapes[i] = shape

	return oldShape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var r float64 = 0.0

	for _, shape := range b.shapes {
		r += shape.CalcPerimeter()
	}

	return r
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var r float64 = 0.0

	for _, shape := range b.shapes {
		r += shape.CalcArea()
	}

	return r
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	shapes := make([]Shape, 0, len(b.shapes))

	isExists := false

	for _, shape := range b.shapes {
		switch shape.(type) {
		case Circle, *Circle:
			isExists = true
		default:
			shapes = append(shapes, shape)
		}
	}

	if !isExists {
		return fmt.Errorf(errorTemplate, errorCirclesAreNotExist)
	}

	b.shapes = shapes

	return nil
}
