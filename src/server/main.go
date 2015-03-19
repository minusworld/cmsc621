package main

import (
	"proj1util"
	"./dict3"
	"fmt"
	"os"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func initializeDict(storageFile string) dict3.DICT3 {
	d := dict3.DICT3{Dict: make(map[dict3.KeyRelationship]interface{}),	StorageFile: storageFile}
	d.Load(storageFile)
	return d
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: server <serverParametersFile>")
		os.Exit(0)
	}
	paramFile := os.Args[1]
	serverParams := proj1util.ConfigureParameters(paramFile)
	d := initializeDict(serverParams.StorageContainer)
	rpc.Register(&d)
	fmt.Println("setting up listener...")
	listener, listenerErr := net.Listen(serverParams.Protocol, ":"+serverParams.Port)
	if listenerErr != nil {
		panic(listenerErr)
	}
	for {
		fmt.Println("waiting for connection...")
		conn, acceptErr := listener.Accept()
		if acceptErr != nil {
			panic(acceptErr)
		}
		fmt.Println("serving...")
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
		fmt.Println("done;")
	}
}
