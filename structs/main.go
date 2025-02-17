package main

import "fmt"

type sendMessage struct {
	message     string
	phoneNumber int
}

/*
type Wheel struct {
	radius   int
	material string
}
*/

type car struct {
	brand string
	model string
	year  int
	/*	backWheel  Wheel
		frontWheel Wheel */
	Wheel struct {
		radius   int
		material string
	}
}

func (car car) setCar(values car) car {
	car.brand = values.brand
	car.model = values.model
	car.year = values.year
	car.Wheel.material = values.Wheel.material
	car.Wheel.radius = values.Wheel.radius
	return car
}

func printCar(c car) {
	car := c.setCar(c)
	// write a print statement to print the car details
	fmt.Printf("Brand: %s\nModel: %s\nYear: %d\nWheel Material: %s\nWheel Radius: %d\n",
		car.brand, car.model, car.year, car.Wheel.material, car.Wheel.radius)
}

func sendMsg(msg sendMessage) {
	fmt.Printf("Sending message: %s to phone number: %d\n", msg.message, msg.phoneNumber)
	fmt.Printf("=============================")
}

func main() {
	sendMsg(sendMessage{"Hello, World!", 1234567890})
	swift := car{
		brand: "suzuki",
		model: "swift",
		year:  2020,
	}
	printCar(swift)
}
