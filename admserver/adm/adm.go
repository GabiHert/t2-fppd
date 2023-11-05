package adm

import (
	"fmt"
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
	fmt.Printf("Adm: Session finished: %s\n", token)
}

func (a *Adm) getAccount(name string, psw int) (int, *account.Account) {
	for i, ac := range a.accounts {
		if ac.Name == name && ac.Psw == psw {
			return i, a.accounts[i]
		}
	}

	return 0, nil
}

func (a *Adm) Auth(req *commom.AuthReq, res *commom.Res) error {
	if err := commom.Fail(); err != nil {
		return err
	}
	_, ac := a.getAccount(req.Name, req.Psw)
	if ac == nil {
		res.Err = &commom.Error{Message: "account not found"}
		return nil
	}

	return nil
}

func (a *Adm) InitSession(_ *struct{}, res *string) error {

	token := uuid.NewString()

	a.sessions[token] = true

	*res = token
	fmt.Printf("Adm: Session started: %s\n", token)

	return nil
}

func (a *Adm) CreateAccount(req *commom.Req, res *commom.Res) error {
	if err := commom.Fail(); err != nil {
		return err
	}

	if !a.validateToken(req.Token) {
		res.Err = &commom.Error{Message: "invalid token"}
		return nil
	}

	defer a.finishSession(req.Token)

	_, ac := a.getAccount(req.Name, req.Psw)
	if ac != nil {
		res.Err = &commom.Error{Message: "account already exists"}
		return nil
	}

	a.accounts = append(a.accounts, &account.Account{
		Name: req.Name, Psw: req.Psw, Balance: 0})

	return nil
}

func (a *Adm) DeleteAccount(req *commom.Req, res *commom.Res) error {
	if !a.validateToken(req.Token) {
		res.Err = &commom.Error{Message: "invalid token"}
		return nil
	}

	defer a.finishSession(req.Token)

	i, ac := a.getAccount(req.Name, req.Psw)
	if ac == nil {
		res.Err = &commom.Error{Message: "account not found"}
		return nil
	}

	a.accounts = append(a.accounts[:i], a.accounts[i+1:]...)

	return nil
}

func (a *Adm) GetBalance(req *commom.Req, res *commom.GetBalanceRes) error {
	if err := commom.Fail(); err != nil {
		return err
	}

	if !a.validateToken(req.Token) {
		res.Err = &commom.Error{Message: "invalid token"}
		return nil
	}

	defer a.finishSession(req.Token)

	_, ac := a.getAccount(req.Name, req.Psw)
	if ac == nil {
		res.Err = &commom.Error{Message: "account not found"}
		return nil
	}

	res.Balance = ac.Balance

	return nil
}

func (a *Adm) Withdraw(req *commom.OperationReq, res *commom.Res) error {
	if err := commom.Fail(); err != nil {
		return err
	}

	if !a.validateToken(req.Token) {
		res.Err = &commom.Error{Message: "invalid token"}
		return nil
	}

	defer a.finishSession(req.Token)

	_, ac := a.getAccount(req.Name, req.Psw)
	if ac == nil {
		res.Err = &commom.Error{Message: "account not found"}
		return nil
	}
	res.Err = ac.Withdraw(req.Amount)
	return nil
}

func (a *Adm) Deposit(req *commom.OperationReq, res *commom.Res) error {
	if err := commom.Fail(); err != nil {
		return err
	}

	if !a.validateToken(req.Token) {
		res.Err = &commom.Error{Message: "invalid token"}
		return nil
	}

	defer a.finishSession(req.Token)

	_, ac := a.getAccount(req.Name, req.Psw)
	if ac == nil {
		res.Err = &commom.Error{Message: "account not found"}
		return nil
	}

	ac.Deposit(req.Amount)

	return nil
}

func NewAdm() *Adm {
	return &Adm{
		sessions: map[string]bool{},
	}
}
