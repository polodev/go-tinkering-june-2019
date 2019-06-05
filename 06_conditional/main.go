package main

import "fmt"

func main() {
	x := 15
	y := 10

	// if
	if x < y {
		fmt.Printf("%d is less than %d", x, y)
	} else {
		fmt.Printf("%d is less than %d", y, x)
	}

	color := "red"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "green" {
		fmt.Println("color is green")
	} else {
		fmt.Println("different color")
	}

	// switch - same as other languages
}
