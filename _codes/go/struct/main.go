package main

import "fmt"

type User struct {
	Name string
}

func NewUser(name *string) *User {
	return &User{
		Name: *name,
	}
}

func main() {
	name := "John"
	nameP := &name
	// nameP := *name

	// * getter
	// & setter

	user := NewUser(nameP)

	fmt.Printf("Name: %s\n", user.Name)
	// fmt.Printf("Name: %s\n", &name)
	fmt.Printf("Name: %s\n", *nameP)
}
