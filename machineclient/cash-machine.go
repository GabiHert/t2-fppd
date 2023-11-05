package machineclient

import (
	"fmt"
	"github.com/GabiHert/t2-fppd/commom"
	"github.com/GabiHert/t2-fppd/machineclient/interfaces"
	"github.com/GabiHert/t2-fppd/rpcservice"
	"net/rpc"
)

type CashMachine struct {
	adm interfaces.Adm
}

func (a *CashMachine) Withdraw(name string, password int, amount float32) error {
	token, err := a.adm.InitSession()
	if err != nil {
		fmt.Printf("Cash Machine: Error starting session: %s\n", err.Error())
		return err
	}

	_, err = commom.Retry(func() (interface{}, error) {
		err = a.adm.Withdraw(name, password, amount, token)
		if err != nil {
			fmt.Printf("Cash Machine: Error during withdraw: %s\n", err.Error())
			return nil, err
		}
		return nil, nil
	})

	if err == nil {
		fmt.Println("Cash Machine: Withdraw successful")
	}
	return err
}

func (a *CashMachine) Deposit(name string, password int, amount float32) error {
	token, err := a.adm.InitSession()
	if err != nil {
		fmt.Printf("Cash Machine: Error starting session: %s\n", err.Error())
		return err
	}

	_, err = commom.Retry(func() (interface{}, error) {
		err = a.adm.Deposit(name, password, amount, token)
		if err != nil {
			fmt.Printf("Cash Machine: Error during deposit: %s\n", err.Error())
			return nil, err
		}
		return nil, nil
	})

	if err == nil {
		fmt.Println("Cash Machine: Deposit successful")
	}
	return err
}

func (a *CashMachine) GetBalance(name string, password int) (float32, error) {
	token, err := a.adm.InitSession()
	if err != nil {
		fmt.Printf("Cash Machine: Error starting session: %s\n", err.Error())
		return 0, err
	}

	balance, err := commom.Retry(func() (float32, error) {
		balance, err := a.adm.GetBalance(name, password, token)
		if err != nil {
			fmt.Printf("Cash Machine: Error getting balance: %s\n", err.Error())
			return 0, err
		}
		return balance, nil
	})

	if err == nil {
		fmt.Println("Cash Machine: balance got with success")
	}
	return balance, err
}

func (a *CashMachine) Auth(name string, password int) error {
	err := a.adm.Auth(name, password)
	if err != nil {
		fmt.Printf("Cash Machine: Error during auth: %s\n", err.Error())
		return err
	}

	fmt.Println("Cash Machine: Auth successful")
	return nil
}

func NewCashMachine(client *rpc.Client) *CashMachine {

	adm := rpcservice.NewAdm(client)
	return &CashMachine{adm: adm}
}
