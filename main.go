package main

import (
	"fmt"

	"github.com/hazim/go-mini-project/myPackage"
	"github.com/hazim/go-mini-project/myPackage/insidepackage"
	"github.com/hazim/go-mini-project/myStruct"
)

const NAME string = "Hazim"

func main() {

	// can only do multiple declaration and initialisation of the same type
	var number, alphabet string = "b", "a"

	// shortcut for multiple types is to use declaration list
	var (
		x    int
		y    bool   = true
		n    string = "annn"
		f, g int
	)

	//lets use the short variable declaration operator :=
	shortA, shortB := false, 67

	fmt.Printf("shorthand variable %v %v\n", shortB, shortA)
	fmt.Printf("%v %v", shortB, shortA)

	fmt.Println("\nJust to print everything \n", number, alphabet, x, y, n, f, g)

	fmt.Println("My name is:", NAME)
	myPackage.Run()
	insidepackage.Run()
	student1 := myStruct.Student{
		FirstName: "Hazim",
		LastName:  "hasnan",
		Age:       26,
		Hobby:     "coding",
	}
	fmt.Printf("--> The student is %+v", student1)
	fmt.Println(student1.PrintStudent())
}
