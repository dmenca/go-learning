package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	result, err := Square(-1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Square of 16 is", result)
}

func Square(value float64) (float64, error) {
	if value < 0 {
		return 0, fmt.Errorf("You can not use negative numbers!")
	}
	return math.Sqrt(value), nil
}
