package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http server address")

var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
}

func echo(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade", err)
		return
	}
	defer func() {
		log.Printf("Closeing connecting for %s", conn.RemoteAddr())
	}()
	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Print("readMessage", err)
			break
		}
		fmt.Printf("%s :recv:%s\n", conn.RemoteAddr(), string(msg))
		err = conn.WriteMessage(messageType, msg)
		if err != nil {
			log.Print("write message", err)
			break
		}

	}
}

func main() {
	flag.Parse()
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
