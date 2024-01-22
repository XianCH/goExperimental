package server

import (
	"errors"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type User struct {
	UserId    string
	Conn      *websocket.Conn
	IsActive  bool
	Subscribe map[string]bool
}

type Room struct {
	Name  string
	users map[string]*User
	mu    sync.Mutex
}

type Message struct {
	Time time.Duration
	From string
	data string
}

func (r *Room) addUser(userId string, conn *websocket.Conn) *User {
	r.mu.Lock()
	defer r.mu.Unlock()

	user := &User{
		UserId:    userId,
		Conn:      conn,
		Subscribe: make(map[string]bool),
	}
	r.users[userId] = user
	return user
}

func (r *Room) removeUser(userId string) error {
	//check user if exit
	if _, ok := r.users[userId]; ok {
		return errors.New("user not exit")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.users, userId)
	return nil
}
