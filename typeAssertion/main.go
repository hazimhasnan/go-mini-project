package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- type assertion ---")

	var myEmptyInterface interface{}
	fmt.Printf("The type is %+v\n", myEmptyInterface)

	var myOtherEmptyInterface interface{}
	fmt.Println("The type is:", myOtherEmptyInterface.(string))

	fmt.Println("Hello")
}
