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
		Person *Identity `json:"person"`
	}
	type AnotherBob struct {
		AnotherPerson Identity `json:"another_person"`
	}
	bob := Bob{Person: nil}
	anotherBob := AnotherBob{}
	fmt.Printf("Bob's identity is %v\n", bob.Person)
	fmt.Printf("AnotherBob's identity is %v\n", anotherBob.AnotherPerson)
}

