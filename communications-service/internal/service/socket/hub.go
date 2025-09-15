package socket

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"slices"

	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/service"
)

type Hub struct {
	// Registered clients.
	Clients map[*Client]bool

	// Inbound messages from the clients.
	Notification chan []byte

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	UnRegister chan *Client

	Message chan []byte
}

func NewHub() *Hub {
	return &Hub{
		Notification: make(chan []byte),
		Message:      make(chan []byte),
		Register:     make(chan *Client),
		UnRegister:   make(chan *Client),
		Clients:      make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.UnRegister:
			if _, ok := h.Clients[client]; ok {
				go service.InformationService().UpdateInformationByUserID(context.Background(), model.InfomationInput{
					UserID: int64(client.Client_Id),
					Status: false,
					Value:  "",
				})
				go service.InformationService().DeleteUserConnected(context.Background(), int64(client.Client_Id))
				delete(h.Clients, client)
				close(client.Send)
			}
		case message := <-h.Notification:
			UserIDS, err := handleNotification(message)
			if err != nil {
				break
			}
			switch ids := UserIDS.(type) {
			case int64:
				for client := range h.Clients {
					if int64(client.Client_Id) == ids {
						select {
						case client.Send <- message:
						default:
							close(client.Send)
							delete(h.Clients, client)
						}
					}
				}
			default:
				for client := range h.Clients {
					select {
					case client.Send <- message:
					default:
						close(client.Send)
						delete(h.Clients, client)
					}
				}
			}

		case message := <-h.Message:
			UserIDS, err := handleMessage(message)
			if err != nil {
				log.Println("Error handling message:", err)
				break
			}
			// Extract int64 user IDs from []model.UserIDS
			var userIDInts []int64
			for _, u := range UserIDS {
				userIDInts = append(userIDInts, u.UserID)
			}
			for client := range h.Clients {
				if slices.Contains(userIDInts, int64(client.Client_Id)) {
					select {
					case client.Send <- message:
					default:
						close(client.Send)
						delete(h.Clients, client)
					}
				}
			}
		}
	}
}
func handleMessage(message []byte) (out []model.UserIDS, err error) {
	fmt.Println("Handle message received")
	var payload Message
	err = json.Unmarshal(message, &payload)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}
	_, err = service.MessageService().CreateMessage(context.Background(), model.MessageInput{
		ConversationID: payload.MessageData.ConversationID,
		UserID:         payload.MessageData.UserID,
		Status:         payload.MessageData.Status,
		Message:        payload.MessageData.Message,
		Type:           payload.MessageData.Type,
	})
	if err != nil {
		log.Println("error creating message: ", err)
		return nil, err
	}
	conversation, err := service.ConversationService().GetConversationById(context.Background(), int64(payload.MessageData.ConversationID))
	if err != nil {
		return nil, err
	}
	UserIDS := conversation.UserIDS
	var filteredUserIDS []model.UserIDS
	for _, u := range UserIDS {
		if u.UserID != int64(payload.MessageData.UserID) {
			filteredUserIDS = append(filteredUserIDS, u)
		}
	}
	UserIDS = filteredUserIDS
	return UserIDS, nil
}
func handleNotification(message []byte) (out any, err error) {
	fmt.Println("Handle notification received")
	var payload Message
	fmt.Println("Message: ", string(message))
	err = json.Unmarshal(message, &payload)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}
	_, err = service.NotificationService().CreateNotification(context.Background(), model.NotificationInput{
		From:    payload.NotificationData.From,
		To:      payload.NotificationData.To,
		Content: payload.NotificationData.Content,
	})
	if err != nil {
		log.Println("error creating notification: ", err)
		return nil, err
	}
	out = payload.NotificationData.To
	return out, nil
}
