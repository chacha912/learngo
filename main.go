package main

import (
	"fmt"

	"github.com/chacha912/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	dictionary["hello"] = "hello world"

	fmt.Println(dictionary)
	fmt.Println(dictionary["hello"])
}
