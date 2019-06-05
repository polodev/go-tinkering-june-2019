## 1. Hello  world

~~~go
package main;

import "fmt"

func main() {
  fmt.Println("Hello world")
}
~~~



## 2. variable & const

types of variable and const is followings:    


* string
* bool
* int
* int int8 int16 int32 int64
* uint uint8 uint16 uint32 uint64 uintptr
* byte -- alias for uint8
* rune -- alias for int32
* float32 float64
* complex64 complex128



2 way we can write variable.   
Using vars keyword or using `:=` operator

~~~go
var name string = "Polo Dev"
name := "Polo dev"
const isCool = true
~~~

default float is `float64`. For `float32` we have to explicit      


~~~go
size := 1.3 // default float is float 64
var size float32 = 2.3 // explicitly set float32
~~~

multiple assignments at a time 

~~~go
name, email := "Polo", "polo@gmail.com"
~~~

printing type using `%T`

~~~go
fmt.Printf("%T\n", size)
~~~


## 3. Packages

Whenever we import a new package it should be added new line and all import should be inside bracket

~~~go
import (
  "fmt"
  "math"

  "github.com/polodev/go_crash_course/03_packages/strutil"
)
~~~


## 4. functions 

~~~go
func greeting(name string) string {
  return "Hello " + name
}

func getSum(num1 int, num2 int) int {
  return num1 + num2
}

func main() {
  fmt.Println(greeting("Shibu"))
  fmt.Println(getSum(3, 3))
}

~~~

## 5. Array and Slices 


### Array

In case of array we need announce number of index when initialized. However slices like js array. No need to prior announcement.

~~~go
var fruitArr [2]string
//assign values
fruitArr[0] = "Apple"
fruitArr[1] = "Orange"
~~~

Declare and assign array in same time    
~~~go
fruitArr := [2]string{"Apple", "Orange"}
~~~

### Slice

~~~go
fruitSlice := []string{"Apple", "Orange", "Grape"}
~~~

get the length of array or slices    

~~~go
fmt.Println(len(fruitSlice))
~~~

get the value in range 
~~~go
// start:stop index value
fmt.Println(fruitSlice[1:3])
~~~


## 6. Condition 

condition as like any other programming (c based)
~~~go
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
~~~

## 7. Loops

Long method     

~~~go
i := 1
for i <= 10 {
  fmt.Println(i)
  // i = i + 1
  i++
}
~~~

short method 
~~~go
for i := 1; i <= 10; i++ {
  fmt.Printf("Number %d\n", i)
}
~~~

fizbuzz 
~~~go
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
~~~


## 8. maps

Declare and assign value    

~~~go
// Define map
// key is [string] and  value is string 
emails := make(map[string]string)
// assign key value
emails["bob"] = "bob@gmail.com"
emails["sharon"] = "sharon@gmail.com"
emails["mike"] = "mike@gmail.com"
~~~



Declare and assign value at once
~~~go
emails := map[string]string{"bob": "bob@gmail.com", "sharon": "sharon@gmail.com", "mike": "mike@gmail.com"}
~~~

Output    

~~~go
fmt.Println(emails)
fmt.Println(len(emails))
fmt.Println(emails["bob"])
~~~

delete a key
~~~go
// map, map key
delete(emails, "bob")
~~~ 

# 9. Range

~~~go
ids := []int{33, 44, 55, 66, 77}

Loop through ids
for i, id := range ids {
  fmt.Printf("%d - ID: %d\n", i, id)
}
~~~

if you not want to use index use `_` (underscore) as id 

~~~go
not using index
for _, id := range ids {
  fmt.Printf("ID: %d\n", id)
}
~~~


Loop through map

~~~go
emails := map[string]string{"bob": "bob@gmail.com", "sharon": "sharon@gmail.com", "mike": "mike@gmail.com"}
for k, v := range emails {
  fmt.Printf("%s: %s\n", k, v)
}
~~~


## 10. Pointers 


~~~go
a := 5
b := &a

fmt.Println(a, b)   // 5, 0xc0000140f0
fmt.Printf("%T", b) // *int
fmt.Println(a, *b)
fmt.Println(*&a) // *b // a

*b = 10
fmt.Println(a) // 10
~~~

to get the pointer value we have prepand `*` before pointer eg. `*b`


## 11. Closures

~~~go
// adder function will return func 
// return function take int value as argument 
// return_function  will return another integer
// this is the reason return is `func(int) int`
func adder() func(int) int {
  sum := 0

  return func(x int) int {
    sum += x
    return sum
  }
}

func main() {
  sum := adder()
  for i := 0; i < 10; i++ {
    fmt.Println(sum(i))
  }
}
~~~

## 12. Struct

Struct is like Class in oop.     


define a  Struct property

~~~go
// long way
type Person struct {
  firstName string
  lastName  string
  city      string
  gender    string
  age       int
}
// short way
type Person struct {
  firstName, lastName, city, gender string
  age                               int
}
~~~

// assign value to struct 

~~~go
person1 := Person{firstName: "Samantha", lastName: "Saint", city: "New York", gender: "f", age: 26}
// Alternative
person1 := Person{"Samantha", "Saint", "New York", "f", 26}
~~~

**Defining method of struct:** we can define method in 2 ways 

* value receiver  (for getter purpose)
* pointer receiver  (for setter purpose)

value receiver    

~~~go
// require strconv pkg
func (p Person) greet() string {
  return "Hello my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}
~~~

Pointer Receiver    

~~~go
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
~~~


using method 

~~~go
person1 := Person{"Samantha", "Saint", "New York", "f", 26}

// fmt.Println(person1)

person1.birthDay()
person1.getMarried("Williams")

fmt.Println(person1.greet())
~~~


13. Interfaces

~~~go
package main

import (
  "fmt"
  "math"
)

type Shape interface {
  area() float64
}

type Circle struct {
  x, y, radius float64
}
type Rectangle struct {
  width, height float64
}

// value pointer methor

func (c Circle) area() float64 {
  return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
  return r.width * r.height
}

func getArea(s Shape) float64 {
  return s.area()
}

func main() {

  circle := Circle{x: 0, y: 0, radius: 5}
  rectangle := Rectangle{width: 10, height: 5}

  fmt.Println(getArea(circle))
  fmt.Println((getArea(rectangle)))
}

~~~

## 14. web

~~~go
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
~~~


