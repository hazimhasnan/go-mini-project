package main

import "fmt"

const (
	pi = 3.142
)

var (
	length = 3
	size   = 34
	name   = "triangle"
)

type User struct {
	FirstName, LastName string
	Location            *UserLocation
}

type UserLocation struct {
	City     string
	Country  string
	PostCode int
}

func NewUser(firstName, lastName string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
	}
}

func (u *User) Greeting() string {
	return fmt.Sprintf("Hey there, My name is %s and I am from %v\n", u.FirstName, u.Location)
}

func main() {
	fmt.Println("--- method ---")
	defer fmt.Println("--- End method ---")

	var hazim = NewUser("hazim", "hasnan")

	fmt.Println(hazim.Greeting())

}
