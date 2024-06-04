package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/x14n/goExperimental/websocket/protomessage/protocol"
	"google.golang.org/protobuf/proto"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade:", err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read:", err)
			break
		}
		var msg protocol.Message
		if err := proto.Unmarshal(message, &msg); err != nil {
			log.Printf("Error unmarshaling message: %v", err)
			continue
		}
		log.Printf("Received message: %+v", msg)
	}
}

func main() {
	http.HandleFunc("/ws", handler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
