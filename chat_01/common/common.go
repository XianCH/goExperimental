package common

import "github.com/gorilla/websocket"

type User struct {
	Username string
	Conn     *websocket.Conn
}

// LoginInfo 结构体表示登录信息
type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Message 结构体表示即时通讯消息
type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
}
