package main

import (
	"fmt"
	"math"
)

func main() {

	result := Square(16)
	fmt.Println("Square of 16 is", result)
}

func Square(value float64) float64 {
	return math.Sqrt(value)
}
