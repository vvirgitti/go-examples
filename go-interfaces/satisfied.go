package main

import "fmt"

type Animal interface {
	Name(name string) string
	Age() int
}
func ReturnData(animal Animal) string {
	name := "Barney"
	return fmt.Sprintf("Hi, my name is %s and I am %d old", animal.Name(name), animal.Age())
}
type Dog struct {}
func (d Dog) Name(name string) string {
	return name
}
func (d Dog) Age() int {
	return 5
}

func main() {
	barney := Dog{}
	fmt.Println(ReturnData(barney))
}