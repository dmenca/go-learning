package main

import "fmt"

func main() {
	first, second, third := primeNumbers()

	fmt.Println("Price numbers:", first, second, third)

}

func primeNumbers() (int, int, int) {
	return 1, 3, 5
}
