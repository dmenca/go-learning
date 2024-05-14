package main

import "os"

func main() {

	panic("a problem")


	if _, err := os.Create("/tmp/file"); err != nil {
		panic(err)
	}
}
