package main

import (
	"fmt"
	"github.com/GabiHert/t2-fppd/admserver"
	"github.com/GabiHert/t2-fppd/agencyclient"
	"github.com/GabiHert/t2-fppd/commom/config"
	"sync"
)

func main() {

	var (
		wg      sync.WaitGroup
		errChan = make(chan error)
	)

	wg.Add(3)

	go func() {
		defer wg.Done()
		err := admserver.Serve()
		if err != nil {
			errChan <- err
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case err := <-errChan:
				fmt.Println(err.Error())
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		agencyAdmClient, err := config.NewRpcClient("8080")
		if err != nil {
			panic(err)
		}

		agency := agencyclient.NewAgency(agencyAdmClient)

		err = agency.Auth("teste", 123)
		if err != nil {
			errChan <- err
		}
	}()

	wg.Wait()

}
