package main

import "fmt"

func main() {
	// Function with multiple return values
	x, y := getCoords(10, 20)
	fmt.Print("Length: ", x, " Breadth: ", y)
}

// Named return values
func getCoords(x, y int) (length, breadth int) {
	x = 10 * 30
	y = 20 * 30
	return x, y
}
