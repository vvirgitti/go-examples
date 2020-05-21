package main

import "fmt"

func main () {
	type Bob struct {
		Name *string `json:"name"`
	}

	nilBob := Bob{
		Name: nil,
	}

	bob := "Bob"
	bobWithName := Bob{
		Name: &bob,
	}
	fmt.Printf("Bob's name is %v\n", nilBob.Name)
	fmt.Printf("BobWithname's name is %v", *bobWithName.Name)
}