package main

import "fmt"

func main() {

	// // Define map
	// emails := make(map[string]string)
	// // assign key value
	// emails["bob"] = "bob@gmail.com"
	// emails["sharon"] = "sharon@gmail.com"
	// emails["mike"] = "mike@gmail.com"

	// Declare map and add kv
	emails := map[string]string{"bob": "bob@gmail.com", "sharon": "sharon@gmail.com", "mike": "mike@gmail.com"}

	// fmt.Println(emails)
	// fmt.Println(len(emails))
	// fmt.Println(emails["bob"])
	delete(emails, "bob")
	fmt.Println(emails)
}
