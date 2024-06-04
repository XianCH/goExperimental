package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/x14n/goExperimental/websocket/protocol_trans/protocol"
	"google.golang.org/protobuf/proto"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	addr := "localhost:8080"
	id := "client2" // Example client ID, change as needed
	u := fmt.Sprintf("ws://%s/ws?id=%s", addr, id)

	c, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			var msg protocol.Message
			if err := proto.Unmarshal(message, &msg); err != nil {
				log.Printf("Error unmarshaling message: %v", err)
				continue
			}
			log.Printf("Received message: %+v", msg)
		}
	}()

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("Enter message: ")
			if !scanner.Scan() {
				return
			}
			text := scanner.Text()
			msg := &protocol.Message{
				Avatar:       "https://example.com/avatar.png",
				FromUsername: "User2",
				From:         id,
				To:           "client1", // Change as needed
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
			err = c.WriteMessage(websocket.BinaryMessage, data)
			if err != nil {
				log.Println("write:", err)
				return
			}
		}
	}()

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then waiting for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
