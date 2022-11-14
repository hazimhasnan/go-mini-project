package main

import (
	"fmt"
	"time"
)

type MyCustomError struct {
	When time.Time
	What string
}

func (e *MyCustomError) Error() string {
	return fmt.Sprintf("Your error was logged at %s and it was: %s", e.When, e.What)
}

func recreateError() error {
	return &MyCustomError{
		When: time.Now(),
		What: "Something went wrong somewhere",
	}
}

func main() {
	fmt.Println("--- Errors ---")
	defer fmt.Println("--- End Errors ---")

	if err := recreateError(); err != nil {
		fmt.Println("Hey, there is an error")
		fmt.Println(err)
	}

}
