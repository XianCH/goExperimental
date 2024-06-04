package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
	"github.com/x14n/goExperimental/websocket/protomessage/protocol"
	"google.golang.org/protobuf/proto"
)

func main() {
	addr := "localhost:9999"
	u := fmt.Sprintf("ws://%s/ws", addr)

	conn, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter message: ")
		if !scanner.Scan() {
			return
		}
		text := scanner.Text()

		msg := &protocol.Message{
			Avatar:       "https://example.com/avatar.png",
			FromUsername: "User1",
			From:         "client1",
			To:           "client2",
			Content:      text,
			ContentType:  1,
			Type:         "chat",
			MessageType:  1,
		}
		data, err := proto.Marshal(msg)
		if err != nil {
			log.Printf("Error marshaling message: %v", err)
			continue
		}
		err = conn.WriteMessage(websocket.BinaryMessage, data)
		if err != nil {
			log.Println("write:", err)
			return
		}
	}
}
