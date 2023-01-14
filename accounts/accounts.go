package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

func (a *Account) Deposit(balance int) {
	a.balance += balance
}

func (a *Account) Withdraw(balance int) error {
	if balance > a.balance{
		return errors.New("Cant withdraw")
	}
	a.balance -= balance
	return nil
}

func (a *Account) Rename (name string) {
	a.owner = name
}

func (a Account) GetBalance() int {
	return a.balance
}

func (a Account) GetOwner() string{
	return a.owner
}

func (a Account) String() string{
	return fmt.Sprint(a.owner, "'s accouunt \nHas: ", a.balance)
}