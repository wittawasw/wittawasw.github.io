package main

import (
	"fmt"
	"time"
)

type UserProfile struct {
	Name      string
	Email     string
	Birthdate time.Time
}

func NewUserProfile(name string, email string, birthdateStr string) (*UserProfile, error) {
	layout := "2006-Jan-02"

	birthdate, err := time.Parse(layout, birthdateStr)
	if err != nil {
		return nil, err
	}

	profile := UserProfile{
		Name:      name,
		Email:     email,
		Birthdate: birthdate,
	}

	return &profile, nil
}

func main() {
	fmt.Println("Starting...")

	name := "John"
	namePointer := &name
	email := "john@example.com"
	birthdateStr := "2014-Jan-14"

	fmt.Printf("Input name: %s\n", name)
	fmt.Printf("Input name: %s\n", *namePointer)

	profile, err := NewUserProfile(name, email, birthdateStr)

	if err != nil {
		fmt.Printf("Error: Formatting Birthdate: %s\n", err)
	} else {
		fmt.Printf("Created User successfully: %s\n", profile.Name)
		fmt.Printf("Created User successfully: %s\n", profile.Birthdate)
	}

	fmt.Println("...End")
}
