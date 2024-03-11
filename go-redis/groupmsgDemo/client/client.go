package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"github.com/x14n/goExperimental/go-redis/groupmsgDemo/pb"
	"google.golang.org/protobuf/proto"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

type Message struct {
	Time     string
	Username string
	Message  string
}

func init() {
	// 连接 Redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 设置密码
		DB:       0,  // 使用默认数据库
	})
}

func main() {
	// 建立 WebSocket 连接
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:12345/ws", nil)
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer conn.Close()

	go readMessages(conn)

	// 读取用户输入并发送消息
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message: ")
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		// 构造消息
		msg := &pb.Message{
			Time:     time.Now().Format("2006-01-02 15:04:05"),
			Username: "Client",
			Message:  message,
		}

		// 将消息序列化为 Protocol Buffers 格式
		data, err := proto.Marshal(msg)
		if err != nil {
			log.Println("Error marshalling message:", err)
			continue
		}

		// 发送消息
		err = conn.WriteMessage(websocket.BinaryMessage, data)
		if err != nil {
			log.Println("Error sending message:", err)
			continue
		}
	}
}

func readMessages(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message from server:", err)
			return
		}

		// 解析并打印消息
		message := &pb.Message{}
		err = proto.Unmarshal(p, message)
		if err != nil {
			log.Println("Error unmarshalling message:", err)
			continue
		}
		fmt.Printf("[%s] %s: %s\n", message.Time, message.Username, message.Message)
	}
}
