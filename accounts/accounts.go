package accounts

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}

	return &account
}

// Deposit x amnount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of account
func (a Account) Balance() int {
	return a.balance
}

var errNoMoney = errors.New("Can't Withdraw")

// Withdraw from account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount

	return nil
}

// ChangeOwner of account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
