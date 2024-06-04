package main

import (
	"fmt"
	"log"
	"net/rpc"
)

const HelloServiceName = "../server/server.HelloService"

type HelloServiceClient struct {
	*rpc.Client
}
type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:12345")
	if err != nil {
		log.Fatal("dial error:", err)
	}

	var reply string
	err = client.Call(secure.HelloServiceName+".Hello", "hello", &reply)
	if err != nil {
		log.Fatal("client call error:", err)
	}
	fmt.Println(reply)
}
