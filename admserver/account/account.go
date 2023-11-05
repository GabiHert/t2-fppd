package account

import (
	"github.com/GabiHert/t2-fppd/commom"
)

type Account struct {
	Name    string
	Balance float32
	Psw     int
}

func (a *Account) Withdraw(amount float32) *commom.Error {
	if a.Balance < amount {
		return &commom.Error{Message: "insufficient funds"}
	}

	a.Balance -= amount

	return nil
}

func (a *Account) Deposit(amount float32) {
	a.Balance += amount
}
