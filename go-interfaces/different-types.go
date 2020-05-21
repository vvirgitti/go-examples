package main

import "fmt"

type animal interface {
	Name(name string) string
	Age() int
}
func returnData(animal animal) string {
	name := "Barney"
	return fmt.Sprintf("Hi, my name is %s and I am %d old", animal.Name(name), animal.Age())
}

type dog struct {}
func (d dog) Name(name string) string {
	return name
}

func main() {
	barney := dog{}
	fmt.Println(returnData(barney))
}
