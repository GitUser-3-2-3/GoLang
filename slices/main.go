package main

import "fmt"

func signUpMessage() [3]string {
	return [3]string{"Welcome to the platform!",
		"Please verify your email address",
		"Thank you for signing up!",
	}
}

func slices() []int {
	a := []int{1, 2, 3, 4, 5}
	b := a[1:4]
	c := make([]int, 5)
	c[1] = 10
	str := "string"
	fmt.Print(" ", len(str))
	fmt.Print("\n", c, len(c), " ", cap(c), " ")
	return []int{a[1], b[1], c[1]}
}

func main() {
	signUp := signUpMessage()
	fmt.Print(signUp)
	fmt.Print(slices())
}
