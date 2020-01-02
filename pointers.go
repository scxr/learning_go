package main

import "fmt"

func main() {
	x := 200
	pointx := &x // this is the memory address where x is stored & denotes this

	fmt.Println(pointx)
	fmt.Println(*pointx) // this reads the value in the memory address denoted by *

	*pointx = 25 // this replaces the value in the memory address where x is stored thus making x the value at the pointx
	fmt.Println(x)
}
