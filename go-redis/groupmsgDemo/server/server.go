package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"github.com/x14n/goExperimental/go-redis/groupmsgDemo/pb"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

type Message struct {
	Time     string
	Username string
	Message  string
}

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "3953",
		DB:       0,
	})
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("websocket upgrade err:%v\n", err)
		return
	}

	defer conn.Close()

	message, err := rdb.LRange(ctx, "group_chat_message", 0, -1).Result()
	if err != nil && err != redis.Nil {
		log.Printf("redis LRange from redis error :%v\n", err)
		return
	}

	if err != redis.Nil {
		for _, msg := range message {
			conn.WriteMessage(websocket.TextMessage, []byte(msg))
		}
	}

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("conn read message error :%v\n", err)
			return
		}

		message := &pb.Message{}
		err = proto.Unmarshal(p, message)
		if err != nil {
			log.Printf("PROTO UNMARSHAL ERROR :%v\n", err)
			continue
		}

		//save message to redis
		data, err := proto.Marshal(message)
		if err != nil {
			log.Printf("redis marshal message error:%v\n", err)
			continue
		}

		rdb.LPush(ctx, "group_chat_message", data)

		//broadcastMessage to every client
		broadcastMessage(conn, p)
	}
}

func broadcastMessage(conn *websocket.Conn, message []byte) {
	for {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("Error broadcast message error %v\n", err)
			return
		}
	}
}

func saveMessageToRedis() {
	for {
		time.Sleep(10 * time.Second)
		msg := &pb.Message{
			Time:     time.Now().Format("2006-01-02 15:04:05"),
			Username: "Server",
			Message:  "This is a test message from server",
		}
		data, err := proto.Marshal(msg)
		if err != nil {
			log.Printf("proto marshal error :%v\n", err)
			continue
		}
		rdb.LPush(ctx, "group_chat_message", data)
	}
}

func main() {
	http.HandleFunc("/ws", handleWebsocket)
	go saveMessageToRedis()
	log.Println("server start at 12345")
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Printf("server start error %v\n", err)
		return
	}
}
