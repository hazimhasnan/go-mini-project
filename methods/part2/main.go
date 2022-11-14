package main

import (
	"fmt"
)

const (
	FileName = "c://hazim/system"
)

type shaper interface {
	calculateArea() float32
	calculateParameter() float32
}

type shapeShifter interface {
	scaleUp(f float32) error
	scaleDown(f float64) error
}

type shape struct {
	length float32
	width  float32
}

func (s *shape) calculateArea() float32 {
	area := s.length * s.length
	return area
}

func (s *shape) calculateParameter() float32 {
	perimeter := (2 * s.length) + (2 * s.width)
	return perimeter
}

func (s *shape) scaleDown(f float32) error {
	if f > 1 {
		return fmt.Errorf("factor cannot be more than 1")
	}
	s.length *= f
	s.width *= f
	return nil
}

func (s *shape) scaleUp(f float32) error {
	if f < 1 {
		return fmt.Errorf("factor cannot be less than 1")
	}
	s.length *= f
	s.width *= f
	return nil
}

func printMyShape(shape shaper) string {
	var description = fmt.Sprintf("My area is %f and my perimeter is %.2f", shape.calculateArea(), shape.calculateParameter())
	return description
}

func main() {
	fmt.Println("--- Method ---")
	defer fmt.Println("--- End Method ---")

	var (
		name = "hazim"
		age  = 18
	)

	myShape := &shape{
		length: 45,
		width:  2,
	}

	var mySecondShape = new(shape)
	mySecondShape.length = 10
	mySecondShape.width = 2

	fmt.Printf("My name is %s and my age is %d and mt file is at this path: %s\n", name, age, FileName)

	fmt.Printf("This is my output:\n%s\n", printMyShape(myShape))

	if err := myShape.scaleDown(0.2); err != nil {
		return
	}

	fmt.Printf("This is my output after scaling:\n%s\n", printMyShape(myShape))

	fmt.Printf("This is my output for secondShape:\n%s\n", printMyShape(mySecondShape))

}
