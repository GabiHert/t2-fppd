package adm

import (
	"errors"
	"github.com/GabiHert/t2-fppd/adm/account"
	"github.com/google/uuid"
)

type Adm struct {
	accounts []*account.Account
	sessions map[string]bool
}

func (a *Adm) validateToken(token string) bool {
	return a.sessions[token]
}

func (a *Adm) finishSession(token string) {
	a.sessions[token] = false
}

func (a *Adm) getAccount(name string, psw int) (int, *account.Account) {
	for i, ac := range a.accounts {
		if ac.Name == name && ac.Psw == psw {
			return i, a.accounts[i]
		}
	}

	return 0, nil
}

func (a *Adm) Autenticate(name string, psw int, token string) error {
	if a.validateToken(token) {
		return errors.New("invalid token")
	}

	defer a.finishSession(token)

	_, account := a.getAccount(name, psw)
	if account == nil {
		return errors.New("Account not found")
	}

	return nil
}

func (a *Adm) InitSession() string {

	token := uuid.NewString()

	a.sessions[token] = true
	return token
}

func (a *Adm) CreateAccount(name string, psw int, token string) error {
	if a.validateToken(token) {
		return errors.New("invalid token")
	}

	defer a.finishSession(token)

	_, ac := a.getAccount(name, psw)
	if ac != nil {
		return errors.New("acount already exists")
	}

	a.accounts = append(a.accounts, &account.Account{Name: name, Psw: psw, Balance: 0})

	return nil
}

func (a *Adm) DeleteAccount(name string, psw int, token string) error {
	if a.validateToken(token) {
		return errors.New("invalid token")
	}

	defer a.finishSession(token)

	i, account := a.getAccount(name, psw)
	if account == nil {
		return errors.New("Account not found")
	}

	a.accounts = append(a.accounts[:i], a.accounts[i+1:]...)

	a.finishSession(token)

	return nil
}

func (a *Adm) GetBalance(name string, psw int, token string) (int, error) {
	if a.validateToken(token) {
		return 0, errors.New("invalid token")
	}

	defer a.finishSession(token)

	_, account := a.getAccount(name, psw)
	if account == nil {
		return 0, errors.New("Account not found")
	}

	return account.Balance, nil
}

func (a *Adm) Withdraw(name string, psw, amount int, token string) error {
	if a.validateToken(token) {
		return errors.New("invalid token")
	}

	defer a.finishSession(token)

	_, account := a.getAccount(name, psw)
	if account == nil {
		return errors.New("Account not found")
	}

	return account.Withdraw(amount)
}

func (a *Adm) Deposit(name string, psw, amount int, token string) error {
	if a.validateToken(token) {
		return errors.New("invalid token")
	}

	defer a.finishSession(token)

	_, account := a.getAccount(name, psw)
	if account == nil {
		return errors.New("Account not found")
	}

	account.Deposit(amount)

	return nil
}

func NewAdm() *Adm {
	return &Adm{
		sessions: map[string]bool{},
	}
}
