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

func (mycar *car) newspeed(newspeed float32) {
	mycar.speed = newspeed
}

func main() {
	/*
		mycar := car{speed:230,
					breakpedal:0,
					maxspeed:300,
					steeringwheel:-30}
		this is in long form */

	mycar := car{230, 0, 300, -30} // we can just do this

	fmt.Println("Our cars speed is (in kmh): ", mycar.speed) // we can reference the cars values
	fmt.Println("Our cars speed is (in mph): ", mycar.convtomph())

	mycar.newspeed(200)

	fmt.Println("Our cars speed is (in kmh): ", mycar.speed) // we can reference the cars values
	fmt.Println("Our cars speed is (in mph): ", mycar.convtomph())

}
