package main

import "fmt"

type User struct {
	Name, Location string
	Id             int
}

// nesting a struct is not considered as composition
type Player struct {
	User
	GameId int
}

type Gamer interface {
	Login() string
}

func (u *User) Login() string {
	fmt.Println("Hi:", u.Id)
	return fmt.Sprintf("Logged in as %s\n", u.Name)
}

func IsGamer(gamer Gamer) {
	fmt.Println("Yes, Implements Gamer Interface")
}

func main() {
	fmt.Println("-- Start of Composition --")
	defer fmt.Println("--- End of composition ---")

	p := Player{
		GameId: 23,
		User: User{
			Name:     "Hazim",
			Location: "Puchong",
			Id:       1,
		},
	}

	fmt.Println(p.User.Login())
	fmt.Println(p.Login())

	fmt.Println("Checking does it meet Gamer interface")
	IsGamer(&p)

	for i := 0; i < 3; i++ {
		fmt.Println("Inside loop")
	}

	x := make([]int, 0, 2)
	x = append(x, 7)
	x = append(x, 9)

	for i, v := range x {
		fmt.Println(i, v)
	}
}
