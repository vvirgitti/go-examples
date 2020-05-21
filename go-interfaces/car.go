package main

import "fmt"

type Car struct {
	Type string
	NumberOfWheels int
}

func (v Car) WheelsCounter() int  {
	return v.NumberOfWheels
}

func CarWheelChecker(car Car) int {
	return car.WheelsCounter()
}

func main() {
	car := Car{"car", 4}
	fmt.Printf("This %s has %d wheels\n", car.Type, car.WheelsCounter())
	fmt.Printf("Checking that the %s has %d wheels", car.Type, CarWheelChecker(car))
}
