package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/quangdat385/holiday-ticket/order-service/global"
	"github.com/quangdat385/holiday-ticket/order-service/internal/database"
	"github.com/quangdat385/holiday-ticket/order-service/internal/model"
	"github.com/quangdat385/holiday-ticket/order-service/internal/model/mapper"
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
	log.Println("Starting Kafka consumer...")

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
			var order model.OrderOutPut
			for range 3 {
				orderDB, err := k.CreateOrder(orderEvent)
				if err != nil {
					continue
				}
				if orderDB.ID != 0 {
					continue
				}
				order = orderDB
				checkOrder = true
				break
			}
			if !checkOrder {
				orderEvent.Order.OrderNotes = "Order-->Failed"
				log.Printf("Failed to create order after retries: %+v\n", orderEvent.Order)
				var message model.Message
				msgBytes, _ := json.Marshal(orderEvent.Order)
				message.Type = "Notification"
				message.NotificationData = model.NotificationData{
					From:    1,
					To:      1,
					Content: string(msgBytes),
				}
				global.Rdb.Publish(context.Background(), "notification", message)
				k.consumer.CommitMessages(context.Background(), m)
				log.Printf("Message: %s\n", string(m.Value))
				continue
			}
			log.Printf("Order created successfully after retries: %+v\n", orderEvent.Order)
			orderEvent.Type = model.OrderEventTypeConfirmOrder
			orderEventJSON, _ := json.Marshal(orderEvent)
			key := fmt.Sprintf("%s-%s", order.OrderNumber, model.OrderEventTypeConfirmOrder)
			msg := kafka.Message{
				Key:   []byte(key),
				Value: orderEventJSON,
				Time:  time.Now(),
			}
			go global.KafkaProducer.WriteMessages(context.Background(), msg)
			k.consumer.CommitMessages(context.Background(), m)
			log.Printf("Message: %s\n", string(m.Value))
			continue
		case model.OrderEventTypeOrderSuccess:
			var checkUpdate bool
			for range 3 {
				_, err := k.r.UpdateOrderNote(context.Background(), database.UpdateOrderNoteParams{
					OrderNumber: orderEvent.Order.OrderNumber,
					OrderNotes:  "Order-->Success",
				})
				if err != nil {
					log.Printf("Error updating order note: %v", err)
					continue
				}
				checkUpdate = true
				break
			}
			if !checkUpdate {
				orderEvent.Order.OrderNotes = "Order-->Update-->Failed"
				log.Printf("Failed to update order after retries: %+v\n", orderEvent.Order)
				var message model.Message
				msgBytes, _ := json.Marshal(orderEvent.Order)
				message.Type = "Notification"
				message.NotificationData = model.NotificationData{
					From:    1,
					To:      1,
					Content: string(msgBytes),
				}
				global.Rdb.Publish(context.Background(), "notification", message)
				k.consumer.CommitMessages(context.Background(), m)
				log.Printf("Message: %s\n", string(m.Value))
				continue
			}
			content := model.ContentType{
				OrderNumber: orderEvent.Order.OrderNumber,
				Message:     "Order processed successfully",
				Status:      true,
			}
			var message model.Message
			msgBytes, _ := json.Marshal(content)
			message.Type = "Notification"
			message.NotificationData = model.NotificationData{
				From:    1,
				To:      orderEvent.Order.UserID,
				Content: string(msgBytes),
			}
			global.Rdb.Publish(context.Background(), "notification", message)
			log.Printf("Received order success event: %+v\n", orderEvent.Order)
			k.consumer.CommitMessages(context.Background(), m)
			log.Printf("Message: %s\n", string(m.Value))
			continue
		case model.OrderEventTypeCancelOrder:
			log.Printf("Received cancel order event: %+v\n", orderEvent.Order)
		case model.OrderEventTypeRefundOrder:
			log.Printf("Received refund order event: %+v\n", orderEvent.Order)
		default:
			continue
		}

	}
}
func (k *KafkaConsumerServiceImpl) Close() error {
	return k.consumer.Close()
}
func (k *KafkaConsumerServiceImpl) CreateOrder(in model.OrderEvent) (out model.OrderOutPut, err error) {
	orderItemJSON, err := json.Marshal(in.Order.OrderItem)
	if err != nil {
		return out, err
	}
	result, err := k.r.InsertOrder(context.Background(), database.InsertOrderParams{
		OrderNumber: in.Order.OrderNumber,
		UserID:      in.Order.UserID,
		StationCode: in.Order.StationCode,
		OrderAmount: strconv.FormatFloat(float64(in.Order.OrderAmount), 'f', -1, 32),
		TerminalID:  in.Order.TerminalID,
		OrderDate:   in.Order.OrderDate,
		OrderNotes:  in.Order.OrderNotes,
		OrderItem:   orderItemJSON,
	})
	if err != nil {
		return out, err
	}
	lastOrderId, err := result.LastInsertId()
	if err != nil {
		return out, err
	}
	order, err := k.r.GetOrderById(context.Background(), int32(lastOrderId))
	if err != nil {
		return out, err
	}
	out = mapper.ToOrderDTO(order)
	return out, nil
}
