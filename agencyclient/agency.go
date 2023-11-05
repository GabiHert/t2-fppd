package agencyclient

import (
	"fmt"
	"github.com/GabiHert/t2-fppd/agencyclient/interfaces"
	"github.com/GabiHert/t2-fppd/commom"
	"github.com/GabiHert/t2-fppd/rpcservice"
	"net/rpc"
)

type Agency struct {
	adm interfaces.Adm
}

func (a *Agency) CreateAccount(name string, password int) error {
	token, err := a.adm.InitSession()
	if err != nil {
		fmt.Printf("Agency: Error starting session: %s\n", err.Error())
		return err
	}

	_, err = commom.Retry(func() (interface{}, error) {
		err = a.adm.CreateAccount(name, password, token)
		if err != nil {
			fmt.Printf("Agency: Error creating account: %s\n", err.Error())
			return nil, err
		}
		return nil, nil
	})

	if err == nil {
		fmt.Println("Agency: Account created successfully")
	}
	return err
}

func (a *Agency) DeleteAccount(name string, password int) error {
	token, err := a.adm.InitSession()
	if err != nil {
		fmt.Printf("Agency: Error starting session: %s\n", err.Error())
		return err
	}

	_, err = commom.Retry(func() (interface{}, error) {
		err = a.adm.DeleteAccount(name, password, token)
		if err != nil {
			fmt.Printf("Agency: Error deletting account: %s\n", err.Error())
			return nil, err
		}
		return nil, nil
	})

	if err == nil {
		fmt.Println("Agency: Account deleted successfully")
	}
	return err
}

func (a *Agency) Withdraw(name string, password int, amount float32) error {
	token, err := a.adm.InitSession()
	if err != nil {
		fmt.Printf("Agency: Error starting session: %s\n", err.Error())
		return err
	}

	_, err = commom.Retry(func() (interface{}, error) {
		err = a.adm.Withdraw(name, password, amount, token)
		if err != nil {
			fmt.Printf("Agency: Error during withdraw: %s\n", err.Error())
			return nil, err
		}
		return nil, nil
	})

	if err == nil {
		fmt.Println("Agency: Withdraw successful")
	}
	return err
}

func (a *Agency) Deposit(name string, password int, amount float32) error {
	token, err := a.adm.InitSession()
	if err != nil {
		fmt.Printf("Agency: Error starting session: %s\n", err.Error())
		return err
	}

	_, err = commom.Retry(func() (interface{}, error) {
		err = a.adm.Deposit(name, password, amount, token)
		if err != nil {
			fmt.Printf("Agency: Error during deposit: %s\n", err.Error())
			return nil, err
		}
		return nil, nil
	})

	if err == nil {
		fmt.Println("Agency: Deposit successful")
	}
	return err
}

func (a *Agency) GetBalance(name string, password int) (float32, error) {
	token, err := a.adm.InitSession()
	if err != nil {
		fmt.Printf("Agency: Error starting session: %s\n", err.Error())
		return 0, err
	}

	balance, err := commom.Retry(func() (float32, error) {
		balance, err := a.adm.GetBalance(name, password, token)
		if err != nil {
			fmt.Printf("Agency: Error getting balance: %s\n", err.Error())
			return 0, err
		}
		return balance, nil
	})
	if err == nil {
		fmt.Println("Agency: Balance got with success")
	}
	return balance, err
}

func (a *Agency) Auth(name string, password int) error {
	err := a.adm.Auth(name, password)
	if err != nil {
		fmt.Printf("Agency: Error during auth: %s\n", err.Error())
		return err
	}
	fmt.Println("Agency: Auth successful")

	return nil
}

func NewAgency(client *rpc.Client) *Agency {

	adm := rpcservice.NewAdm(client)
	return &Agency{adm: adm}
}
