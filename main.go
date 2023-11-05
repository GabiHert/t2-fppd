package main

import (
	"fmt"
	"github.com/GabiHert/t2-fppd/admserver"
	"github.com/GabiHert/t2-fppd/agencyclient"
	"github.com/GabiHert/t2-fppd/cli"
	"github.com/GabiHert/t2-fppd/commom"
	"github.com/GabiHert/t2-fppd/machineclient"
)

func main() {

	var (
		stop = make(chan bool)
	)

	go func() {
		err := admserver.Serve()
		if err != nil {
			fmt.Println(err.Error())
			stop <- true
		}
	}()

	go func() {
		agencyAdmClient, err := commom.NewRpcClient("8080")
		if err != nil {
			fmt.Println(err.Error())
			stop <- true
		}
		cashMachineAdmClient, err := commom.NewRpcClient("8080")
		if err != nil {
			fmt.Println(err.Error())
			stop <- true
		}
		agency := agencyclient.NewAgency(agencyAdmClient)
		cashMachine := machineclient.NewCashMachine(cashMachineAdmClient)

		for {
			cli.NewCli(agency, cashMachine, stop).Run()
		}

	}()

	<-stop

}
