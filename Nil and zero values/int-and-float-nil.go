package main

import "fmt"

func main () {
	type Rob struct {
		Age int `json:"age"`
		Height float32 `json:"height"`
	}

	rob := Rob{
		Age: nil,
		Height: nil,
	}
	fmt.Printf("rob's age is %d and height is %d", rob.Age, rob.Height)
}
