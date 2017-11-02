package main

import "fmt"

func main() {

	for i := 60; i < 170; i++ {
		fmt.Printf("%d \t %b \t %#X %q \t \n", i, i, i, i)
	}
}
