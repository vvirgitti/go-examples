package main

import "fmt"

func main () {
	type Bob struct {
		Funny *bool  `json:"funny"`
	}

	nilBob := Bob{
		Funny: nil,
	}

	bob := true
	funnyBob := Bob{
		Funny: &bob,
	}
	fmt.Printf("Is NilBob funny? %v\n", nilBob.Funny)
	fmt.Printf("Is FunnyBob funny? %v", *funnyBob.Funny)
}
