package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/x14n/goExperimental/websocket/protocol_trans/protocol"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
}

func WebsHaddler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("websocket upgrade err :%s\n", err)
		return
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("read msg err: %s\n", err)
			return
		}
		clientMessage := &protocol.Message{}

		err = proto.Unmarshal(msg, clientMessage)
		if err != nil {
			log.Printf("unmarshal err:%s\n", err)
			return
		}

		log.Printf("revec message from client:%s\n", clientMessage.Content)

		if clientMessage.Content == "hello server" && clientMessage.FromUsername == "client" {
			serverMessage := &protocol.Message{
				FromUsername: "server",
				Content:      "hello client",
			}

			response, err := proto.Marshal(serverMessage)
			if err != nil {
				log.Printf("marshal err : %s\n", err)
				return
			}

			err = conn.WriteMessage(websocket.BinaryMessage, response)
			if err != nil {
				log.Printf("write message err : %s\n", err)
				return
			}

		}
	}
}

func main() {
	http.HandleFunc("/ws", WebsHaddler)
	http.ListenAndServe(":8080", nil)
}
