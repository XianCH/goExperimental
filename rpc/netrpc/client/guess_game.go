package client

import (
	"fmt"
	"log"
	"net/rpc"
)

type GuessArgs struct {
	Number int
}

func StartGuessClient() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Printf("Dial error: %v", err)
		return
	}

	for {
		var guess int
		var result int
		fmt.Scan(&guess)

		args := &GuessArgs{Number: guess}
		err = client.Call("GuessGameServer.GuessNumber", args, &result)
		if err != nil {
			log.Printf("Call error: %v", err)
		}

		if result == 0 {
			log.Println("You win!")
		} else if result == -1 {
			log.Println("Too low")
		} else {
			log.Println("Too hight")
		}
	}
}
