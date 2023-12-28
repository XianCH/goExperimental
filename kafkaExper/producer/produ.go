package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func main() {
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: "test_01",
		Key:   sarama.StringEncoder("find"),
		Value: sarama.StringEncoder("hello"),
	}

	// Send the message and capture the offset
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("partition:%d offset:%d\n", partition, offset)
}
