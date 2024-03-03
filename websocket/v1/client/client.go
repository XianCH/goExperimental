package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"github.com/x14n/goExperimental/websocket/v1/pt"
)

func main() {
	fmt.Println("connect to server ...")
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:12345/ws", nil)
	if err != nil {
		fmt.Printf("websocket dial error :%v", err)
		return
	}
	defer conn.Close()

	go func() {
		for {
			var msg pt.Message
			err := conn.ReadJSON(&msg)
			if err != nil {
				log.Printf("read json error :%v", err)
				return
			}
			fmt.Println("Received:", msg)
		}

	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		time := time.Now()
		msg := &pt.Message{
			From: "client",
			To:   "server",
			Name: "x14n",
			Msg:  text,
			Time: time.Format("2006-01-02 15:04:05"),
		}
		//send Message
		err := conn.WriteJSON(msg)
		if err != nil {
			log.Println("wirte", err)
			return
		}
	}
}
