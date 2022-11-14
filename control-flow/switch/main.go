package main

import "fmt"

func main() {
	fmt.Println("--- Switch ---")
	defer fmt.Println("--- Switch ---")

	var (
		x int
		l = "hazim"
		n []string
		// myMap = make(map[string]float)
	)

	fmt.Println(x)
	fmt.Println(l)
	n = append(n, "hazim", "haziq", "hariz", "thomas")

	for _, v := range l {
		fmt.Printf("value is: %v and its type is %T\n", string(v), string(v))
	}

}
