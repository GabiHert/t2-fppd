package rpcservice

import (
	"fmt"
	"github.com/GabiHert/t2-fppd/commom"
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
	req := commom.AuthReq{
		Name: name,
		Psw:  psw,
	}
	var res commom.Res

	errCall := a.client.Call("Adm.Auth", &req, &res)
	if errCall != nil {
		return errCall
	}

	if res.Err != nil {
		fmt.Println(res.Err.Error())
		return nil
	}

	return nil
}

func (a *Adm) CreateAccount(name string, psw int, token string) error {
	req := commom.Req{
		Name:  name,
		Psw:   psw,
		Token: token,
	}

	var res commom.Res

	errCall := a.client.Call("Adm.CreateAccount", &req, &res)
	if errCall != nil {
		return errCall
	}

	if res.Err != nil {
		fmt.Println(res.Err.Error())
		return nil
	}

	return nil
}

func (a *Adm) DeleteAccount(name string, psw int, token string) error {

	req := commom.Req{
		Name:  name,
		Psw:   psw,
		Token: token,
	}

	var res commom.Res

	errCall := a.client.Call("Adm.DeleteAccount", &req, &res)
	if errCall != nil {
		return errCall
	}

	if res.Err != nil {
		fmt.Println(res.Err.Error())
		return nil
	}
	return nil
}

func (a *Adm) GetBalance(name string, psw int, token string) (float32, error) {

	req := commom.Req{
		Name:  name,
		Psw:   psw,
		Token: token,
	}
	var res commom.GetBalanceRes

	err := a.client.Call("Adm.GetBalance", &req, &res)
	if err != nil {
		return 0, err
	}

	if res.Err != nil {
		fmt.Println(res.Err.Error())
		return 0, nil
	}

	return res.Balance, nil
}

func (a *Adm) Deposit(name string, psw int, amount float32, token string) error {

	req := commom.OperationReq{
		Req: commom.Req{
			Name:  name,
			Psw:   psw,
			Token: token,
		},
		Amount: amount,
	}

	var res commom.Res
	errCall := a.client.Call("Adm.Deposit", &req, &res)
	if errCall != nil {
		return errCall
	}

	if res.Err != nil {
		fmt.Println(res.Err.Error())
		return nil
	}
	return nil
}

func (a *Adm) Withdraw(name string, psw int, amount float32, token string) error {

	req := commom.OperationReq{
		Req: commom.Req{
			Name:  name,
			Psw:   psw,
			Token: token,
		},
		Amount: amount,
	}

	var res commom.Res
	errCall := a.client.Call("Adm.Withdraw", &req, &res)
	if errCall != nil {
		return errCall
	}

	if res.Err != nil {
		fmt.Println(res.Err.Error())
		return nil
	}
	return nil
}

func NewAdm(client *rpc.Client) *Adm {
	return &Adm{client: client}
}
