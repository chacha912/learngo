package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	people := [2]string{"nico", "urim"}
	for _, person := range people {
		go isSexy(person, c)
	}
	result := <-c
	fmt.Println(result)
	fmt.Println(<-c)
}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	c <- true
}
