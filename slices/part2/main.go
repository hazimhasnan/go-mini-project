package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("--- slices ---")
	defer fmt.Println("--- End slices ---")

	originalSlices := []int{4, 4, 4, 4, 4, 4}

	loopSlices(originalSlices)

	sliceOfSlice := originalSlices[:2]
	sliceOfSlice[0] = 5

	newSlices := append(originalSlices, 34)

	fmt.Printf("%v is of type: %T\n", sliceOfSlice, sliceOfSlice)

	fmt.Println("New original slice:", originalSlices)
	fmt.Println("New slice:", newSlices)
	fmt.Println("New new new original slice:", originalSlices)
}

func loopSlices(s interface{}) (err error) {
	// get the concrete type value of the interface
	cs, ok := s.([]int)
	if !ok {
		return errors.New("not a slice of integers")
	}

	// print the concrete interface
	for i, v := range cs {
		fmt.Printf("The key is %d, value is %d\n", i, v)
	}
	return nil
}
