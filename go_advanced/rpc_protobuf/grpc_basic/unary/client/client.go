package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/x14n/goExperimental/go_advanced/rpc_protobuf/grpc_basic/protos"
	"google.golang.org/grpc"
)

func main() {
	//echoServer()
}

func echoServer() {
	//http2 默认使用安全连接，我们这里不是用安全连接
	conn, err := grpc.Dial("127.0.0.1:12345", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	c := protos.NewEchoServiceClient(conn)

	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Println("reader line error :", err)
			return
		}
		req := protos.EchoRequest{Req: string(line)}
		res, err := c.GetUnaryEcho(context.Background(), &req)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(res.GetRes())
	}
}
