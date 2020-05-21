package main

import "fmt"

type Vehicule interface {
	WheelsCounter() int
}
func WheelDetector(vehicule Vehicule) int {
	return vehicule.WheelsCounter()
}

type Unicycle struct {
	Type string
	NumberOfWheels int
}
func (u Unicycle) WheelsCounter() int  {
	return u.NumberOfWheels
}

func main() {
	unicycle := Unicycle{"unicyle", 1}
	fmt.Printf("I am a %s and I have %d wheel(s)", unicycle.Type, WheelDetector(unicycle))
}
