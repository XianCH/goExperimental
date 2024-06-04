package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"os"
	"os/signal"
	"time"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Received message: %s\n", message)
		}
	}()

	for {
		select {
		case <-interrupt:
			fmt.Println("Interrupt received, closing connection...")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				fmt.Println(err)
			}
			return
		case <-time.After(time.Second):
			message := "Hello, server!"
			err := c.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Sent message: %s\n", message)
		}
	}
}
