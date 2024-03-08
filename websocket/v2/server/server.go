package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 将HTTP连接升级为WebSocket协议
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	// 向客户端发送消息
	if err := conn.WriteMessage(websocket.TextMessage, []byte("Hello from server")); err != nil {
		log.Println("Error writing message to client:", err)
		return
	}

	// 读取来自客户端的消息
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Println("Error reading message:", err)
		return
	}
	log.Printf("Received message from client: %s", message)
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	log.Fatal(http.ListenAndServe(":12345", nil))
}
