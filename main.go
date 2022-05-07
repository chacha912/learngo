package main

import "fmt"

func main() {
	a := 2
	b := a
	c := &a
	a = 10
	*c = 200
	fmt.Println(a, b, c, *c)
	fmt.Println(&a, &b, &c)
}
