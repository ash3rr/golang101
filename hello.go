package main

import (
	"errors"
	"fmt"
	"math"
)

type person struct {
	name string
	age  int
}

func main() {

	fmt.Println("Hello, world!")
	var x int = 1
	y := 5 // short hard to create and assign value, type is inferred
	var a [5]int
	a[2] = 7
	a[4] = 55

	bslice := []int{5, 2} // a slice array, is not fixed in size, so you can easily extend with an append

	bslice = append(bslice, 15)

	b := [5]int{5, 2, 3, 4, 2}

	// maps hold key value pairs
	// this example is saying make a map,
	// make the key a string and the value an int
	vertices := make(map[string]int)
	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12

	delete(vertices, "square")

	fmt.Println(vertices)

	fmt.Println(x)

	p := person{name: "asher", age: 100}

	fmt.Println(p.age)

	fmt.Println(p)

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 0
	// while loop
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// increment over and print an array
	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}

	// increment over and print a map
	m := make(map[string]string)

	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println("key:", key, "value:", value)
	}

	// if else if else example

	if x > y {
		fmt.Println("x is more than y")

	} else if x < y {
		fmt.Println("x is less than y")
	} else {
		fmt.Println("go fuck yourself. Pun intended, get it? Go... ")
	}

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(sum(50, 50))

	result, err := sqrt(-16)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println(result)
	}

	k := 7
	incmemoryaddress(&k) // send pointer

	fmt.Println(k)
}

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}

// this should increment any value passed, but it doesnt.
func incmemoryaddress(x *int) {
	*x++ //without the star, you're incrementing the memory address, the star deferences the memory address to use the value instead
}

// by prefixing the type with ana asterix
// when calling it, prefixing the variable with ampersand to send the pointer instead
