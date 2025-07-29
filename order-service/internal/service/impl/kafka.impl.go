package impl

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/quangdat385/holiday-ticket/order-service/global"
	"github.com/quangdat385/holiday-ticket/order-service/internal/model"
	"github.com/quangdat385/holiday-ticket/order-service/internal/service"
	"github.com/quangdat385/holiday-ticket/order-service/internal/vo"
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
		// Commit the message after processing
		var orderEvent model.OrderEvent
		if err := json.Unmarshal(m.Value, &orderEvent); err != nil {
			log.Printf("Error unmarshalling message: %v", err)
		}
		switch orderEvent.Type {
		case model.OrderEventTypeCreateOrder:
			var checkOrder bool
			for range 3 {
				_, err := service.OrderService().CreateOrder(context.Background(), vo.CreateOrderRequest{
					OrderNUmber: orderEvent.Order.OrderNumber,
					OrderAmount: orderEvent.Order.OrderAmount,
					TerminalID:  orderEvent.Order.TerminalID,
					OrderDate:   orderEvent.Order.OrderDate,
					OrderNotes:  "Order-->Success",
				})
				if err == nil {
					log.Printf("Order created successfully: %+v\n", orderEvent.Order)
					checkOrder = true
					break
				}
				log.Printf("Failed to create order, retrying... Error: %v", err)
				continue
			}
			if !checkOrder {
				log.Printf("Failed to create order after retries: %+v\n", orderEvent.Order)
				var message model.Message
				orderEvent.Order.OrderNotes = "Order-->Failed"
				msgBytes, _ := json.Marshal(orderEvent.Order)
				message.Type = "Notification"
				message.NotificationData = model.NotificationData{
					From:    1,
					To:      1,
					Content: string(msgBytes),
				}
				global.Rdb.Publish(context.Background(), "notification", message)
			}
			log.Printf("Order created successfully after retries: %+v\n", orderEvent.Order)
			arrayOrderNumber := strings.Split(orderEvent.Order.OrderNumber, "-")
			userIdStr := arrayOrderNumber[1]
			userId, _ := strconv.ParseInt(userIdStr, 10, 64)
			var message model.Message
			msgBytes, _ := json.Marshal(orderEvent.Order)
			message.Type = "Notification"
			message.NotificationData = model.NotificationData{
				From:    1,
				To:      int(userId),
				Content: string(msgBytes),
			}
			global.Rdb.Publish(context.Background(), "notification", message)

			// orderEvent.Order.OrderNotes = "Order-->Pending"
			// orderEvent.Type = model.OrderEventTypeReOrder
			// message, _ := json.Marshal(orderEvent)
			// m = kafka.Message{
			// 	Key:   []byte(orderEvent.Order.OrderNumber),
			// 	Value: message,
			// }
			// go global.KafkaProducer.WriteMessages(context.Background(), m)
		case model.OrderEventTypeReOrder:
			log.Printf("Received re-order event: %+v\n", orderEvent.Order)
		case model.OrderEventTypeCancelOrder:
			log.Printf("Received cancel order event: %+v\n", orderEvent.Order)
		case model.OrderEventTypeRefundOrder:
			log.Printf("Received refund order event: %+v\n", orderEvent.Order)
		default:
			log.Printf("Received unknown event type: %s", orderEvent.Type)
		}
		k.consumer.CommitMessages(context.Background(), m)
		log.Printf("Message: %s\n", string(m.Value))
	}
}
