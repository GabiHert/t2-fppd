package admserver

import (
	"github.com/GabiHert/t2-fppd/admserver/adm"
	"net"
	"net/http"
	"net/rpc"
)

func Serve() error {

	admServer := adm.NewAdm()

	err := rpc.Register(admServer)
	if err != nil {
		return err
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	err = http.Serve(listener, nil)
	if err != nil {
		return err
	}

	return nil
}
