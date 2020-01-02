package main

import "fmt"

func add(x, y float64) float64 {
	return x + y
}

func main() {
	num1, num2 := 5.6, 4.4
	var result float64 = add(num1, num2)
	fmt.Println("Your result is ", result)
}
