package main

import "fmt"

func main() {
	fmt.Println("--- slices ---")
	defer fmt.Println("--- End slices ---")

	var mySliceOfString []string

	var secondSliceString []string = []string{"hello", "my", "name"}

	var thirdSliceString = []string{"My", "hobby", "is", "eating"}

	fmt.Printf("%v\n", mySliceOfString)

	fmt.Printf("%q\n", secondSliceString)

	fmt.Printf("%v\n", thirdSliceString)

	mySliceOfString = append(mySliceOfString, secondSliceString...)

	fmt.Printf("%q", mySliceOfString)
}
