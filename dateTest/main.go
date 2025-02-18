package main

import (
	"fmt"
	"time"

	"github.com/wagslane/go-tinytime"
)

func main() {
	tt := tinytime.New(1585750374)
	tt = tt.Add(time.Hour * 48)
	fmt.Print(tt, "\n")
	fmt.Print(tt.ToTime())
}
