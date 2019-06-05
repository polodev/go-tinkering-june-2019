package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)   // 5, 0xc0000140f0
	fmt.Printf("%T", b) // *int
	fmt.Println(a, *b)
	fmt.Println(*&a) // *b // a

	*b = 10
	fmt.Println(a) // 10

}
