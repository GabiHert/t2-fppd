package main

import (
	"github.com/GabiHert/t2-fppd/adm/adm"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {

	admServer := adm.NewAdm()
	// Register the timeserver object upon which the GiveServerTime
	// function will be called from the RPC server (from the client)
	err := rpc.Register(admServer)
	if err != nil {
		panic(err)
	}
	// Registers an HTTP handler for RPC messages
	rpc.HandleHTTP()
	// Start listening for the requests on port 1234
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listener error: ", err)
	}
	// Serve accepts incoming HTTP connections on the listener l, creating
	// a new service goroutine for each. The service goroutines read requests
	// and then call handler to reply to them
	http.Serve(listener, nil)
}
