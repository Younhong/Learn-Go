package bank

import (
	"fmt"

	"github.com/younhong/Learn-Go/accounts"
)

func Bank() {
	account := accounts.NewAccount("younhong")
	account.Deposit(10)
	err := account.Withdraw(5)
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
	}
	fmt.Println(account)
}
