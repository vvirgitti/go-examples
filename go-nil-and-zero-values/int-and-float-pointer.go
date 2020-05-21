package main

import "fmt"

func main () {
	type Rob struct {
		Age *int `json:"age"`
		Height *float64 `json:"height"`
	}
	rob := Rob{
		Age: nil,
		Height: nil,
	}
	fmt.Printf("rob's age is %v and height is %v\n", rob.Age, rob.Height)
	robAge := 33
	robHeight := 1.80
	robbie := Rob{
		Age:    &robAge,
		Height: &robHeight,
	}
	fmt.Printf("robbie's age is %v and height is %v\n", *robbie.Age, *robbie.Height)
}
