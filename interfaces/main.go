package main

import "fmt"

// zero value for interface is nil

// interface can be considered as tuple of its value and type
// type is basically the type of variable that meets the interface
// value is basically the exact value

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) (returnedData string) {
	return
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	var data = "data"
	c.L.Process(data)
}

func main() {
	fmt.Println("--- Interface ---")
	defer fmt.Println("--- End Interface ---")

	c := Client{
		L: LogicProvider{},
	}

	c.Program()
}
