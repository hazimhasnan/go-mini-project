package main

import "fmt"

func main() {
	fmt.Println("--- slices ---")
	defer fmt.Println("--- End slices ---")

	var mySliceOfString []string

	var secondSliceString = make([]string, 0)

	var thirdSliceString = []string{"My", "hobby", "is", "eating"}

	fmt.Printf("%+v\n", mySliceOfString)

	fmt.Printf("%q\n", secondSliceString)

	fmt.Printf("%v\n", thirdSliceString)

	if mySliceOfString == nil {
		fmt.Println("MySliceofString is nil")
		fmt.Println(len(mySliceOfString))
	}

	if secondSliceString == nil {
		fmt.Println("secondSliceOfString is nil")
	}

	if thirdSliceString == nil {
		fmt.Println("thirdSliceOrString is nil")
	}
}
