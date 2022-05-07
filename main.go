package main

import (
	"fmt"

	"github.com/chacha912/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("yourim")
	fmt.Println(account)
	fmt.Printf("%p", account)
}
