package main

import "fmt"

// import "fmt"

func fizbuzz() {

	for i := 1; i <= 100; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fiz-buzz", i)
		} else if i%3 == 0 {
			fmt.Println("fiz", i)
		} else if i%5 == 0 {
			fmt.Println("buzz", i)
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	// Long method
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	// i = i + 1
	// 	i++
	// }

	// Short method
	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("Number %d\n", i)
	// }
	fizbuzz()
}
