package secure

import "net/rpc"

const HelloServiceName = "../server/server.HelloService"

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

// 我们将 RPC 服务的接口规范分为三个部分：首先是服务的名字，
//然后是服务要实现的详细方法列表，
//最后是注册该类型服务的函数。
//为了避免名字冲突，我们在 RPC 服务的名字中增加了包路径前缀
//（这个是 RPC 服务抽象的包路径，并非完全等价 Go 语言的包路径）。
//RegisterHelloService 注册服务时，编译器会要求传入的对象满足 HelloServiceInterface 接口。
