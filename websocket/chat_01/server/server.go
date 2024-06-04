package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/x14n/goExperimental/chat_01/common"
)

var (
	clients     = make(map[*websocket.Conn]common.User)
	clientMutex sync.Mutex
	upgrader    = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func handlerConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer ws.Close()

	var loginInfo common.LoginInfo
	err = ws.ReadJSON(&loginInfo)
	if err != nil {
		fmt.Println(err)
		return
	}

	if isValidLogin(loginInfo) {
		clientMutex.Lock()
		clients[ws] = common.User{Username: loginInfo.Username, Conn: ws}
		clientMutex.Unlock()

		//处理用户消息
		for {
			var msg common.Message
			err := ws.ReadJSON(&msg)
			if err != nil {
				fmt.Println(err)
				deletCilent(ws)
				panic(err)
			}
			broadcast(msg)
		}
	} else {
		fmt.Println("login is attempt")
	}
}

func deletCilent(ws *websocket.Conn) {
	clientMutex.Lock()
	delete(clients, ws)
	clientMutex.Unlock()
}

func isValidLogin(loginInfo common.LoginInfo) bool {
	// 在这里可以进行更复杂的身份验证逻辑，例如检查用户名和密码
	return true
}

func broadcast(msg common.Message) {
	clientMutex.Lock()
	defer clientMutex.Unlock()

	for _, client := range clients {
		if client.Username == msg.To {
			err := client.Conn.WriteJSON(&msg.Content)
			if err != nil {
				fmt.Println(err)
				deletCilent(client.Conn)
			}
		}
	}
}

func main() {
	http.HandleFunc("/ws", handlerConnections)
	fmt.Println("server is start running on 8080:")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
