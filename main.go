package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi"}
	urim := person{"chayourim", 29, favFood}
	fmt.Println(urim.name)

	nico := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nico.name)
}
