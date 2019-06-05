package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte -- alias for uint8
	// rune -- alias for int32
	// float32 float64
	// complex64 complex128

	// Using vars

	name := "Polo dev"
	//size := 1.3
	var size float32 = 2.3
	//var age int32 = 37

	name, email := "Polo", "polo@gmail.com"
	fmt.Println(name, email)

	const isCool = true
	fmt.Printf("%T\n", size)

}
