package main

import "fmt"

type Moto struct {
	Type string
	NumberOfWheels int
}

func (m Moto) WheelsCounter() int  {
	return m.NumberOfWheels
}

func MotoWheelChecker(moto Moto) int {
	return moto.WheelsCounter()
}

func main() {
	moto := Moto{"moto", 2}
	fmt.Printf("This %s has %d wheels\n", moto.Type, moto.WheelsCounter())
	fmt.Printf("Checking that the %s has %d wheels", moto.Type, MotoWheelChecker(moto))
}
