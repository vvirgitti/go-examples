package main

import "fmt"

func main () {
	type Bob struct {
		Funny bool  `json:"funny"`
	}

	bob := Bob{
		Funny: nil,
	}

	fmt.Printf("Is Bob funny? %v\n", bob.Funny)
}
