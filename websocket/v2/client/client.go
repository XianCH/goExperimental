package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	// 连接WebSocket服务器
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:12345/ws", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	// 读取来自服务器的消息
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}
	log.Printf("Received message from server: %s", message)

	// 向服务器发送消息
	if err := conn.WriteMessage(websocket.TextMessage, []byte("Hello from client")); err != nil {
		log.Println("Error writing message to server:", err)
		return
	}

	// 等待一段时间以确保消息传递完成
	time.Sleep(1 * time.Second)
}
