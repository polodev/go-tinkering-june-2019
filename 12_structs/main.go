package main

import (
	"fmt"
	"strconv"
)

// struct works like class in oop
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// There are 2 kinds of method in go
// value receiver and pointer receiver

// Gretting - value receiver
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// birthday // pointer receiver

func (p *Person) birthDay() {
	p.age++
}

// get Married

func (p *Person) getMarried(husbandLastName string) {
	if p.gender == "m" {
		return
	}
	p.lastName += " " + husbandLastName
}

func main() {
	// person1 := Person{firstName: "Samantha", lastName: "Saint", city: "New York", gender: "f", age: 26}

	// Alternative
	person1 := Person{"Samantha", "Saint", "New York", "f", 26}

	// fmt.Println(person1)

	person1.birthDay()
	person1.getMarried("Williams")

	// getting single value
	fmt.Println(person1.greet())

}
