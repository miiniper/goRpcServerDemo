package main

import (
	"fmt"
	"goRpcServerDemo/rpcDemo"
)

func main() {
	fmt.Println("goRpcServerDemo server starting ... ")

	//rpcDemo
	rpcDemo.Register()

}
