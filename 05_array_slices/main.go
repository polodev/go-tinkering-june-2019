package main

import "fmt"

func main() {
	// // Arrays
	// var fruitArr [2]string

	// //assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// declare and assign at once
	//fruitArr := [2]string{"Apple", "Orange"}

	//fmt.Println(fruitArr)

	fruitSlice := []string{"Apple", "Orange", "Grape"}

	//fmt.Println(fruitSlice)
	//fmt.Println(len(fruitSlice))

	// start and stops value
	fmt.Println(fruitSlice[1:3])
}
