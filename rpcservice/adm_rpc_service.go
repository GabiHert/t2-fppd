package rpcservice

import (
	"github.com/GabiHert/t2-fppd/commom/types"
	"net/rpc"
)

type Adm struct {
	client *rpc.Client
}

func (a *Adm) InitSession() (string, error) {
	var session string
	err := a.client.Call("Adm.InitSession", &struct{}{}, &session)
	if err != nil {
		return "", err
	}

	return session, nil
}

func (a *Adm) Auth(name string, psw int) error {
	req := types.AuthReq{
		Name: name,
		Psw:  psw,
	}

	err := a.client.Call("Adm.Auth", &req, &struct{}{})
	if err != nil {
		return err
	}

	return nil
}

func (a *Adm) CreateAccount(name string, psw int, token string) error {
	req := types.Req{
		Name:  name,
		Psw:   psw,
		Token: token,
	}

	err := a.client.Call("Adm.CreateAccount", &req, &struct{}{})
	if err != nil {
		return err
	}

	return nil
}

func (a *Adm) DeleteAccount(name string, psw int, token string) error {

	req := types.Req{
		Name:  name,
		Psw:   psw,
		Token: token,
	}

	err := a.client.Call("Adm.DeleteAccount", &req, &struct{}{})
	if err != nil {
		return err
	}

	return nil
}

func (a *Adm) GetBalance(name string, psw int, token string) (float32, error) {

	req := types.Req{
		Name:  name,
		Psw:   psw,
		Token: token,
	}
	var balance float32

	err := a.client.Call("Adm.GetBalance", &req, &balance)
	if err != nil {
		return 0, err
	}

	return balance, nil
}

func (a *Adm) Deposit(name string, psw int, amount float32, token string) error {

	req := types.OperationReq{
		Req: types.Req{
			Name:  name,
			Psw:   psw,
			Token: token,
		},
		Amount: amount,
	}

	err := a.client.Call("Adm.Deposit", &req, &struct{}{})
	if err != nil {
		return err
	}

	return nil
}

func (a *Adm) Withdraw(name string, psw int, amount float32, token string) error {

	req := types.OperationReq{
		Req: types.Req{
			Name:  name,
			Psw:   psw,
			Token: token,
		},
		Amount: amount,
	}

	err := a.client.Call("Adm.Withdraw", &req, &struct{}{})
	if err != nil {
		return err
	}

	return nil
}

func NewAdm(client *rpc.Client) *Adm {
	return &Adm{client: client}
}
