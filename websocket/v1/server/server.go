package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/x14n/goExperimental/websocket/v1/pt"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan *pt.Message)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnect(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("ws upgrader error:%s", err)
		return
	}
	defer ws.Close()

	clients[ws] = true

	for {
		var msg pt.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Printf("error : %v", err)
			delete(clients, ws)
			break
		}
		broadcast <- &msg
	}
}

func HandleMessage() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Println("write json error", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	http.HandleFunc("/ws", HandleConnect)
	go HandleMessage()
	http.ListenAndServe(":12345", nil)
}
