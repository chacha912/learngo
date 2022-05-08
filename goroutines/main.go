package main

import (
	"fmt"
	"time"
)

func main() {
	go sexyCount("nico")
	go sexyCount("yourim")
	time.Sleep(time.Second * 2)
}

func sexyCount(person string) {
	for i := 0; i < 3; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
