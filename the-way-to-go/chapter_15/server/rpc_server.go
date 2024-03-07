package main

import (
	"hutao/the-way-to-go/chapter_15/rpc_objects"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

func main() {
	calc := new(rpc_objects.Args)
	rpc.Register(calc)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Starting RPC-server -listen error:", err)
	}
	go http.Serve(listener, nil)
	time.Sleep(1000e9)

}
