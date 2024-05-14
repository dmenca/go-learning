package main

import "fmt"

type Minute int
type Hour int

func main() {
	minute := Minute(70)
	hour := Hour(10)
	if minute > 60 {
		fmt.Println("Minutes is greater than 60")
	}

	if hour < 15 {
		fmt.Println("Hhours is greater than 15")
	}

}
