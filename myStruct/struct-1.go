package myStruct

import "fmt"

type Student struct {
	FirstName string
	LastName  string
	Age       int
	Hobby     string
}

func (student Student) PrintStudent() string {
	return fmt.Sprintf("The student is %+v", student)
}
