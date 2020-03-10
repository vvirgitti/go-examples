package main

import "fmt"

func main () {
	type Bob struct {
		Name string `json:"name"`
	}

	bob := Bob{
		Name: nil,
	}
	fmt.Printf("Bob's name is %q\n", bob.Name)
}
