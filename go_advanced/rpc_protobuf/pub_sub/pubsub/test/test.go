package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/x14n/goExperimental/go_advanced/rpc_protobuf/pub_sub/pubsub"
)

func main() {
	p := pubsub.NewPublisher(10, time.Millisecond*100)
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	all := p.SubscribeAll()

	go func() {
		for sub := range golang {
			fmt.Println("goalng:", sub)
		}
	}()

	go func() {
		for sub := range all {
			fmt.Println(sub)
		}
	}()

	p.Publishe("hello golang!")
	p.Publishe("hello world!")
}
