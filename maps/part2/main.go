package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("--- Map ---")
	defer fmt.Println("--- End Map ---")

	myMaps := make(map[string]string)

	myMaps["name"] = "hazim"
	myMaps["age"] = "seventeen"

	fmt.Printf("%#v and %v\n\n", myMaps, myMaps)

	fmt.Printf("the value of the key is: %v\n", myMaps["level"])

	v, ok := myMaps["level"]
	if ok {
		fmt.Println("the key does not exist")
		os.Exit(0)
	}

	println("the value is:", v)
}
