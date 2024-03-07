package main

import (
	"fmt"
	"hutao/the-way-to-go/chapter_15/rpc_objects"
	"log"
	"net/rpc"
)

const serverAddress = "127.0.0.1"

func main() {
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("Error dialing:", err)
	}
	// Synchronous call
	args := &rpc_objects.Args{7, 8}
	var reply int
	err = client.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Args error:", err)
	}
	fmt.Printf("Args: %d * %d = %d", args.N, args.M, reply)
}
