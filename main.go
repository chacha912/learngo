package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func lenAndUpper2(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	fmt.Println(multiply(2, 2))

	totalLength, upperName := lenAndUpper("urim")
	fmt.Println(totalLength, upperName)

	totalLength, upperName = lenAndUpper2("ong")
	fmt.Println(totalLength, upperName)

	repeatMe("nico", "yourim", "dal", "ong")
}
