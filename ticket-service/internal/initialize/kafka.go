package initialize

import (
	"fmt"
	"log"
	"time"

	"github.com/quangdat385/holiday-ticket/ticket-service/global"
	"github.com/segmentio/kafka-go"
)

// Init kafka Producer
var KafkaProducer *kafka.Writer

func InitKafka() {
	kafkaHost := global.Config.Kafka.Host
	kafkaPort := global.Config.Kafka.Port
	url := fmt.Sprintf("%s:%d", kafkaHost, kafkaPort)
	global.KafkaProducer = &kafka.Writer{
		Addr:     kafka.TCP(url),
		Topic:    "order-ticket", // topic
		Balancer: &kafka.LeastBytes{},
	}
}

func CloseKafka() {
	if err := global.KafkaProducer.Close(); err != nil {
		log.Fatalf("Failed to close kafka producer: %v", err)
	}
}

// generate Kafka Consumer
func InitKafkaConsumer() *kafka.Reader {
	kafkaHost := global.Config.Kafka.Host
	kafkaPort := global.Config.Kafka.Port
	url := fmt.Sprintf("%s:%d", kafkaHost, kafkaPort)
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{url},
		Topic:          "order-ticket",
		GroupID:        "ticket-group",
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		StartOffset:    kafka.LastOffset,
	})
}
