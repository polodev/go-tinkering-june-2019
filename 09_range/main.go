package main

import "fmt"

func main() {
	// ids := []int{33, 44, 55, 66, 77}

	// Loop through ids
	// for i, id := range ids {
	// 	fmt.Printf("%d - ID: %d\n", i, id)
	// }
	// not using index
	// for _, id := range ids {
	// 	fmt.Printf("ID: %d\n", id)
	// }

	// sum := 0

	// for _, id := range ids {
	// 	sum += id
	// }

	// fmt.Println("sum", sum)

	emails := map[string]string{"bob": "bob@gmail.com", "sharon": "sharon@gmail.com", "mike": "mike@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

}
