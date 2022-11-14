package main

import "fmt"

func main() {
	fmt.Println("--- Map ---")
	defer fmt.Println("--- End Map ---")

	var myMaps = map[string]string{}

	var mySecondMaps = make(map[string]string)

	var myNilMaps map[string]string

	var myOldMap map[string]string = map[string]string{}

	fmt.Printf("myMaps: %#v", myMaps)
	fmt.Println(mySecondMaps)
	fmt.Println(myOldMap)

	if myOldMap != nil {
		fmt.Println("Its not nil")
	}

	if mySecondMaps == nil {
		fmt.Println("Its nil")
	}

	if myNilMaps == nil {
		fmt.Println("myNilMaps is nil")
	}

}
