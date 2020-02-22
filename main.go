package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if err := kafkaConsumer.Close(); err != nil {
			log.Println(err)
		}
	}()

	if err := kafkaConsumer.SubscribeTopics([]string{"image_process", "^aRegex.*[Tt]opic"}, nil); err != nil {
		log.Println(err)
		return
	}

	for {
		msg, err := kafkaConsumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

}
