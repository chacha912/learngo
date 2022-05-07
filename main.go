package main

import "fmt"

func main() {
	names := [5]string{"nico", "urim", "ong"}
	names[3] = "ala"
	names[4] = "ala"
	// names[5] = "ala"
	fmt.Println(names)

	fruits := []string{"🍎", "🍇", "🍌"}
	fruits = append(fruits, "🍓", "🍉", "🍍")
	fmt.Println(fruits)
}
