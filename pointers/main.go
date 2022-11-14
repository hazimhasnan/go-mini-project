package main

import "fmt"

// primitives cant be preceded with address operator

// use address operator if you want to point to a struct

// new is rarely used

// * is a indirection operator. its is used for dereferencing

// pointer type is a must

// use helper function to convert primitive to pointer in struct or map or other DS

// it takes 1 nanosecond to pass a pointer to function, a millisecond to pass 10mb of data ( like struct etc)

func main() {
	fmt.Println("---- pointers ----")
	defer fmt.Println("--- end of program ----")
}
