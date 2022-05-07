package main

import (
	"fmt"

	"github.com/chacha912/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("yourim")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}

	account.ChangeOwner("ong")
	fmt.Println(account.Balance(), account.Owner())
}
