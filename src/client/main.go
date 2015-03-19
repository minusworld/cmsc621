package main

import (
	"fmt"
	"encoding/json"
	"os"
	"bufio"
	"net"
	"proj1util"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: client <serverParametersFile>")
		os.Exit(0)
	}
	paramFile := os.Args[1]
	serverParams := proj1util.ConfigureParameters(paramFile)

	jsonObject := new(interface{})
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter json RPC call:")
	jsonMessage, readerErr := reader.ReadString('\n')
	checkError(readerErr)
	
	conn, dialErr := net.Dial(serverParams.Protocol, serverParams.IpAddress+":"+serverParams.Port)
	checkError(dialErr)
	
	fmt.Fprintf(conn, jsonMessage)
	response, responseErr := bufio.NewReader(conn).ReadString('\n')
	checkError(responseErr)
	
	jsonUnmarshalErr := json.Unmarshal([]byte(response), &jsonObject)
	checkError(jsonUnmarshalErr)
	
	reJson, _ := json.Marshal(jsonObject)
	fmt.Println(string(reJson))
}
