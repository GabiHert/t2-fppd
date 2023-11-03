package account

import "errors"

type Account struct {
	Name    string
	Balance float32
	Psw     int
}

func (a *Account) Withdraw(amount float32) error {
	if a.Balance < amount {
		return errors.New("insufficient funds")
	}

	a.Balance -= amount

	return nil
}

func (a *Account) Deposit(amount float32) {
	a.Balance += amount
}
