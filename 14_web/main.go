package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world from home page</h1>")
}
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About page</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server Started")
	http.ListenAndServe(":3000", nil)
}
