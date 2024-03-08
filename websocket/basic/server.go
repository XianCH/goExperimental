package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("http upgrade error :%v\n", err)
		return
	}

	defer conn.Close()

	//read message from the client
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("conn read message error :%v\n", err)
			break
		}

		log.Printf("Received message :%s\v", message)

	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	log.Fatal(http.ListenAndServe(":12345", nil))
}
