package main

import (
	"fmt"

	"github.com/chacha912/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err2 := dictionary.Update(baseWord, "Second")

	if err2 != nil {
		fmt.Println(err2)
	}

	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)
}
