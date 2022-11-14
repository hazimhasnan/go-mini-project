package myPackage

import (
	"fmt"

	"github.com/pborman/uuid"
)

func Run() {
	fmt.Println("Running...")
	fmt.Println(uuid.NewRandom())
}
