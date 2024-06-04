package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/x14n/goExperimental/go_advanced/thread/experimental/chatServer"
)

type Server struct {
	conn      map[string]*websocket.Conn
	Rooms     map[string]*Room
	Read      chan []byte
	Boradcast chan []string
}

type User struct {
	UserId   string
	Conn     *websocket.Conn
	IsActive bool
}

type Room struct {
	Name      string
	users     map[string]*User
	mu        sync.Mutex
	readchan  chan string
	writechan chan string
	msgList   *chatServer.List
}

type Message struct {
	Time       time.Duration
	FromUserId string
	toRoom     string
	data       string
}

func (r *Room) addUser(userId string, conn *websocket.Conn) *User {
	r.mu.Lock()
	defer r.mu.Unlock()

	user := &User{
		UserId: userId,
		Conn:   conn,
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

func NewServer() *Server {
	return &Server{
		conn:      make(map[string]*websocket.Conn),
		Rooms:     make(map[string]*Room),
		Read:      make(chan []byte),
		Boradcast: make(chan []string),
	}
}

func (s *Server) handleWebsocket(conn *websocket.Conn, userId string, selectRoom string) {
	defer conn.Close()

	//如果房间不存在则返回
	room, exit := s.Rooms[selectRoom]
	if !exit {
		conn.WriteMessage(websocket.TextMessage, []byte("房间不存在"))
		return
	}
	//将用户添加到房间内
	room.addUser(userId, conn)
	defer room.removeUser(userId)
	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				// 处理连接断开的情况
				log.Printf("User %s disconnected from room %s", userId, selectRoom)
				return
			}

			// 处理用户发送的消息，例如将消息发送到 s.Read
			// 这里的示例是将消息格式化成字符串并发送到 s.Read
			s.Read <- []byte(msg)
		}
	}()
}

func (s *Server) Run() {
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		upgrade := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}

		conn, err := upgrade.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Error upgrade websocket", err)
			return
		}
		userId := r.URL.Query().Get("userId")
		room := r.URL.Query().Get("roomId")
		go s.handleWebsocket(conn, userId, room)
	})

	go func() {
		for {
			select {
			case data := <-s.Read:
				msg := string(data)
				// 因为消息的格式是:time/user/msg/group
				strSlice := strings.Split(msg, "/")
				if _, exit := s.Rooms[strSlice[3]]; !exit {
					log.Printf("room :%s is not exit", strSlice[3])

					return
				}
				s.Boradcast <- strSlice
			case data := <-s.Boradcast:
				room := s.Rooms[data[3]]
				for _, user := range room.users {
					if user.IsActive {
						err := user.Conn.WriteJSON(data)
						if err != nil {
							log.Println("error boradcast message:", err)
							return
						}
					}
				}
			default:
				// 处理 WebSocket 连接关闭的情况
				log.Println("WebSocket connection closed")
				// 在这里进行一些清理工作，例如从相关的数据结构中移除用户等
				return
			}
		}
	}()

	fmt.Println("server is running on :12345")
	http.ListenAndServe(":12345", nil)
}

func (s *Server) RoomInit() {
	roomSlice := []string{"game", "life", "study"}
	for _, roomName := range roomSlice {
		s.Rooms[roomName] = &Room{
			Name:      roomName,
			users:     make(map[string]*User),
			readchan:  make(chan string),
			writechan: make(chan string),
			msgList:   chatServer.NewList(),
		}
	}
}

func main() {
	s := NewServer()
	s.RoomInit()
	s.Run()
}
