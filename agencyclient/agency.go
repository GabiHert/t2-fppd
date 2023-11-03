package agencyclient

import (
	"github.com/GabiHert/t2-fppd/rpcservice"
	"net/rpc"
)

type Agency struct {
	adm *rpcservice.Adm
}

func (a *Agency) CreateAccount(name string, password int) error {
	token, err := a.adm.InitSession()
	if err != nil {
		return err
	}

	err = a.adm.CreateAccount(name, password, token)
	if err != nil {
		return err
	}
	return nil
}

func (a *Agency) DeleteAccount(name string, password int) error {
	token, err := a.adm.InitSession()
	if err != nil {
		return err
	}

	err = a.adm.DeleteAccount(name, password, token)
	if err != nil {
		return err
	}
	return nil
}

func (a *Agency) GetBalance(name string, password int) (float32, error) {
	token, err := a.adm.InitSession()
	if err != nil {
		return 0, err
	}

	balance, err := a.adm.GetBalance(name, password, token)
	if err != nil {
		return 0, err
	}
	return balance, nil
}

func (a *Agency) Auth(name string, password int) error {
	err := a.adm.Auth(name, password)
	if err != nil {
		return err
	}
	return nil
}

func NewAgency(client *rpc.Client) *Agency {

	adm := rpcservice.NewAdm(client)
	return &Agency{adm: adm}
}
