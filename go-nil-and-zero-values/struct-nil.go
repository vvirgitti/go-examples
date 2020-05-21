package main

import "fmt"

func main () {
	type Identity struct {
		Name string `json:"name"`
		Age int `json:"age"`
		Height float64 `json:"height"`
		IsFunny bool `json:"is_funny"`
	}

	type Bob struct {
		Person Identity `json:"person"`
	}

	bob := Bob{Person: nil}
	fmt.Printf("Bob's identity is %v\n", bob.Person)
}
