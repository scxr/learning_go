package main

import "fmt"

type car struct {
	speed         uint16
	breakpedal    uint16
	maxspeed      uint16
	steeringwheel int16
}

func main() {
	/*
		mycar := car{speed:230,
					breakpedal:0,
					maxspeed:300,
					steeringwheel:-30}
		this is in long form */

	mycar := car{230, 0, 300, -30} // we can just do this

	fmt.Println("Our cars speed is : ", mycar.speed) // we can reference the cars values

}
