package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/x14n/goExperimental/go_advanced/rpc_protobuf/grpc_basic/protos"
	"google.golang.org/grpc"
)

const (
	ServerAddress string = "127.0.0.1:12345"
)

func main() {
	Bidirectional()
	// clientSide()
	// echoServer()
	//ServerSide()
}

func Bidirectional() {
	conn, err := grpc.Dial(ServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	c := protos.NewBidireactionalClient(conn)
	stream, err := c.BidirectionalHello(context.Background())
	if err != nil {
		log.Fatalf("call rpc func error :%v", err)
	}
	for n := 0; n <= 5; n++ {
		err := stream.Send(&protos.BiRequest{Message: "double stream rpc" + strconv.Itoa(n)})
		if err != nil {
			log.Fatalf("stream request error :%v", err)
		}
		res, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("read has been over")
			break
		}
		if err != nil {
			log.Fatalf("revc message error:%v", err)
		}

		log.Println(res.GetMessage())
	}
}

func clientSide() {
	conn, err := grpc.Dial(ServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	//build grpc connection
	c := protos.NewClientSideClient(conn)
	res, err := c.ClientSideHello(context.Background())
	if err != nil {
		log.Fatalf("Call ClientSide Func error :%v", err)
	}
	for i := 0; i < 5; i++ {
		err = res.Send(&protos.ClientRequest{Req: "fuck u"})
		if err != nil {
			log.Println(err)
			return
		}
	}
	//接受响应
	fmt.Println(res.CloseAndRecv())
}

func ServerSide() {
	conn, err := grpc.Dial(ServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	c := protos.NewServerSideClient(conn)
	req := protos.ServerSideRequest{
		SideReq: "你说你妈呢",
	}
	stream, err := c.ServerSideHello(context.Background(), &req)
	if err != nil {
		log.Fatalf("call SayHello func err : %s", err)
	}
	for n := 0; n < 5; n++ {
		res, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("over")
			break
		}
		if err != nil {
			log.Fatalf("Conversations get stream err : %v", err)
		}
		log.Println(res)
	}
}

func echoServer() {
	//http2 默认使用安全连接，我们这里不是用安全连接
	conn, err := grpc.Dial(ServerAddress, grpc.WithInsecure())
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
