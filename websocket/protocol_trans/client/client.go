package main

import (
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/x14n/goExperimental/websocket/protocol_trans/protocol"
)

func main() {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
	log.Println("connect to server", u.String())

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Printf("dial err:%s\n", err)
		return
	}

	defer conn.Close()

	done := make(chan struct{})

	clientMessage := &protocol.Message{
		FromUsername: "client",
		Content:      "hello server",
	}

	clientByteMessage, err := proto.Marshal(clientMessage)

	if err != nil {
		log.Printf("marshal err : %s\n", err)
		return
	}

	err = conn.WriteMessage(websocket.BinaryMessage, clientByteMessage)
	if err != nil {
		log.Printf("write message err : %s\n", err)
		return
	}

	go func() {
		defer close(done)
		for {
			_, p, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}

			serverResponse := &protocol.Message{}
			err = proto.Unmarshal(p, serverResponse)
			if err != nil {
				log.Println(err)
				return
			}

			log.Printf("Received response from server: %+v", serverResponse)
		}
	}()

	select {
	case <-done:
	case <-interrupt:
		log.Println("interrupt")
		err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			log.Println("write close:", err)
			return
		}
		select {
		case <-done:
		case <-time.After(time.Second):
		}
	}

}
