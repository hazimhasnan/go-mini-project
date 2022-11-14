package main

import "fmt"

func main() {
	fmt.Println("--- slices ---")
	defer fmt.Println("--- End slices ---")

	var my3DArray [3][3]int

	fmt.Printf("%v\n", my3DArray)

	for i := 0; i < len(my3DArray); i++ {
		for j := 0; j < len(my3DArray[0]); j++ {
			my3DArray[i][j] = 1
		}
	}

	fmt.Printf("%d\n", my3DArray)
}
