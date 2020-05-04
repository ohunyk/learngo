package main

import (
	"fmt"
	"log"

	"github.com/ohunyk/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("ohunyk")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
}
