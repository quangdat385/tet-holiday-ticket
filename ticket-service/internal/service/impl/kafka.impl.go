package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/quangdat385/holiday-ticket/ticket-service/global"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/database"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/service"
	"github.com/segmentio/kafka-go"
)

type KafkaConsumerServiceImpl struct {
	consumer *kafka.Reader
	r        *database.Queries
}

func NewKafkaConsumerServiceImpl(consumer *kafka.Reader, r *database.Queries) *KafkaConsumerServiceImpl {
	return &KafkaConsumerServiceImpl{
		consumer: consumer,
		r:        r,
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
		// Commit the message after processing
		var orderEvent model.OrderEvent
		if err := json.Unmarshal(m.Value, &orderEvent); err != nil {
			log.Printf("Error unmarshalling message: %v", err)
		}
		switch orderEvent.Type {
		case model.OrderEventTypeConfirmOrder:
			_, code := service.TicketItem().DecreaseStock(context.Background(), int(orderEvent.Order.OrderItem.ItemID), int(orderEvent.Order.OrderItem.ItemCount))
			if code == 1 {
				orderEvent.Type = model.OrderEventTypeReConfirmOrder
				orderEventJSON, _ := json.Marshal(orderEvent)
				key := fmt.Sprintf("%s-%s", orderEvent.Order.OrderNumber, orderEvent.Type)
				msg := kafka.Message{
					Key:   []byte(key),
					Value: orderEventJSON,
					Time:  time.Now(),
				}
				go global.KafkaProducer.WriteMessages(context.Background(), msg)
				k.consumer.CommitMessages(context.Background(), m)
				log.Printf("Message: %s\n", string(m.Value))
				continue
			}
			if code == 2 {
				msg := model.ContentType{
					OrderNumber: orderEvent.Order.OrderNumber,
					Message:     "Stock not available",
					Status:      false,
				}
				var message model.Message
				msgBytes, _ := json.Marshal(msg)
				message.Type = "Notification"
				message.NotificationData = model.NotificationData{
					From:    1,
					To:      int(orderEvent.Order.UserID),
					Content: string(msgBytes),
				}
				go global.Rdb.Publish(context.Background(), "notification", message)
				k.consumer.CommitMessages(context.Background(), m)
				continue
			}
			orderEvent.Type = model.OrderEventTypeOrderSuccess
			orderEventJSON, _ := json.Marshal(orderEvent)
			key := fmt.Sprintf("%s-%s", orderEvent.Order.OrderNumber, orderEvent.Type)
			msg := kafka.Message{
				Key:   []byte(key),
				Value: orderEventJSON,
				Time:  time.Now(),
			}
			go global.KafkaProducer.WriteMessages(context.Background(), msg)
			k.consumer.CommitMessages(context.Background(), m)
			log.Printf("Message: %s\n", string(m.Value))
			continue
		case model.OrderEventTypeReConfirmOrder:
			orderEvent.Type = model.OrderEventTypeConfirmOrder
			orderEventJSON, _ := json.Marshal(orderEvent)
			key := fmt.Sprintf("%s-%s", orderEvent.Order.OrderNumber, orderEvent.Type)
			msg := kafka.Message{
				Key:   []byte(key),
				Value: orderEventJSON,
				Time:  time.Now(),
			}
			go global.KafkaProducer.WriteMessages(context.Background(), msg)
			k.consumer.CommitMessages(context.Background(), m)
			log.Printf("Message: %s\n", string(m.Value))
			continue
		default:
			continue
		}
	}
}
