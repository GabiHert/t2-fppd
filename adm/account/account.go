package account

import "errors"

type Account struct {
	Name         string
	Psw, Balance int
}

func (a *Account) Withdraw(amount int) error {
	if a.Balance < amount {
		return errors.New("insufficient funds")
	}

	a.Balance -= amount

	return nil
}

func (a *Account) Deposit(amount int) {
	a.Balance += amount
}
