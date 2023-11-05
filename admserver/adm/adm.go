package adm

import (
	"errors"
	"github.com/GabiHert/t2-fppd/admserver/account"
	"github.com/GabiHert/t2-fppd/commom"
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

func (a *Adm) Auth(req *commom.AuthReq, _ *struct{}) error {

	_, ac := a.getAccount(req.Name, req.Psw)
	if ac == nil {
		return errors.New("account not found")
	}

	return nil
}

func (a *Adm) InitSession(_ *struct{}, res *string) error {

	token := uuid.NewString()

	a.sessions[token] = true

	res = &token
	return nil
}

func (a *Adm) CreateAccount(req *commom.Req, _ *struct{}) error {
	if a.validateToken(req.Token) {
		return errors.New("invalid token")
	}

	defer a.finishSession(req.Token)

	_, ac := a.getAccount(req.Name, req.Psw)
	if ac != nil {
		return errors.New("acount already exists")
	}

	a.accounts = append(a.accounts, &account.Account{
		Name: req.Name, Psw: req.Psw, Balance: 0})

	return nil
}

func (a *Adm) DeleteAccount(req *commom.Req, _ *struct{}) error {
	if a.validateToken(req.Token) {
		return errors.New("invalid token")
	}

	defer a.finishSession(req.Token)

	i, ac := a.getAccount(req.Name, req.Psw)
	if ac == nil {
		return errors.New("account not found")
	}

	a.accounts = append(a.accounts[:i], a.accounts[i+1:]...)

	return nil
}

func (a *Adm) GetBalance(req *commom.Req, res *float32) error {
	if a.validateToken(req.Token) {
		return errors.New("invalid token")
	}

	defer a.finishSession(req.Token)

	_, ac := a.getAccount(req.Name, req.Psw)
	if ac == nil {
		return errors.New("account not found")
	}

	*res = ac.Balance

	return nil
}

func (a *Adm) Withdraw(req *commom.OperationReq, _ *struct{}) error {
	if a.validateToken(req.Token) {
		return errors.New("invalid token")
	}

	defer a.finishSession(req.Token)

	_, ac := a.getAccount(req.Name, req.Psw)
	if ac == nil {
		return errors.New("account not found")
	}
	return ac.Withdraw(req.Amount)
}

func (a *Adm) Deposit(req *commom.OperationReq, _ *struct{}) error {
	if a.validateToken(req.Token) {
		return errors.New("invalid token")
	}

	defer a.finishSession(req.Token)

	_, ac := a.getAccount(req.Name, req.Psw)
	if ac == nil {
		return errors.New("account not found")
	}

	ac.Deposit(req.Amount)

	return nil
}

func NewAdm() *Adm {
	return &Adm{
		sessions: map[string]bool{},
	}
}
