package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
	"github.com/x14n/goExperimental/chat_01/common"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := "ws://localhost:8080/ws"
	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	// 发送登录信息
	loginInfo := common.LoginInfo{Username: "clientA", Password: "passwordA"}
	err = ws.WriteJSON(loginInfo)
	if err != nil {
		fmt.Println(err)
		return
	}

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			var msg common.Message
			err := ws.ReadJSON(&msg)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("[clientA] Received message: %+v\n", msg)
		}
	}()

	for {
		select {
		case <-interrupt:
			fmt.Println("[clientA] Interrupt received, closing connection...")
			err := ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				fmt.Println(err)
			}
			return
		}
	}
}
