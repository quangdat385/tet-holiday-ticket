package impl

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumerServiceImpl struct {
	consumer *kafka.Reader
}

func NewKafkaConsumerServiceImpl(consumer *kafka.Reader) *KafkaConsumerServiceImpl {
	return &KafkaConsumerServiceImpl{
		consumer: consumer,
	}
}
func (k *KafkaConsumerServiceImpl) Consume() {

	for {
		m, err := k.consumer.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}
		log.Printf("Message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
		// Process the message
		// For example, you can unmarshal the message value into a struct
		// var msg YourMessageType
		// err = json.Unmarshal(m.Value, &msg)
		// if err != nil {
		// 	log.Println("Error unmarshalling message:", err)
		// 	continue
		// }
		// Perform your business logic here
		// For example, you can save the message to a database or perform some action based on its content
		// After processing the message, commit it
		// This is important to mark the message as processed

		k.consumer.CommitMessages(context.Background(), m)
		log.Printf("Message: %s\n", string(m.Value))
	}
}
