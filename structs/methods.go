package main

import "fmt"

const kmhtomph float32 = 1.60934

type car struct {
	speed         float32
	breakpedal    uint16
	maxspeed      uint16
	steeringwheel int16
}

func (mycar car) convtomph() float32 {
	return (mycar.speed / kmhtomph)
}

func main() {
	mycar := car{230, 0, 300, -30} // we can just do this

	fmt.Println("Our cars speed is (in kmh): ", mycar.speed) // we can reference the cars values
	fmt.Println("Our cars speed is (in mph): ", mycar.convtomph())
}
