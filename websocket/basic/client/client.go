package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:12345/ws", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	done := make(chan struct{})

	// 监听ctrl + c 信号以便优雅的关闭
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		defer close(done)

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Printf("reade message error :%v\n", err)
				return
			}

			log.Printf("Received :%s\n", message)
		}
	}()

	//定期发送消息给服务器
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// 关闭连接
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}

}
