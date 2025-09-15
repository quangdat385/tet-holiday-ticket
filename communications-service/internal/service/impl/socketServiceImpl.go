package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/quangdat385/holiday-ticket/communications-service/global"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/service"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/service/socket"
	"github.com/quangdat385/holiday-ticket/communications-service/utils"
	"github.com/redis/go-redis/v9"
)

type sSocket struct {
	rdb      *redis.Client
	upgrader websocket.Upgrader
	ctx      context.Context
}

func NewSocketImpl(rdb *redis.Client) *sSocket {
	pubsub := rdb.Subscribe(context.Background(), "notification")
	fmt.Println("Subscribed to Redis channel: notification")
	go func() {
		fmt.Println("Listening for messages from Redis...")
		for msg := range pubsub.Channel() {
			log.Printf("Received message from Redis: %s", msg.Payload)

			var payload socket.Message
			err := json.Unmarshal([]byte(msg.Payload), &payload)
			if err != nil {
				log.Printf("Error unmarshalling message: %v", err)
				continue
			}
			if payload.Type == "Message" {
				global.Hub.Message <- []byte(msg.Payload)
			}
			global.Hub.Notification <- []byte(msg.Payload)
		}
	}()
	return &sSocket{
		rdb: rdb,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		ctx: context.Background(),
	}
}
func (s *sSocket) Connect(c *gin.Context) {
	client_id := c.GetInt("UserID")
	if client_id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client_id"})
		return
	}
	log.Println("client_id: ", client_id)
	clientIDInt := int64(client_id)

	clientID := utils.GenerateRandomString(10)
	_, err := service.InformationService().UpdateInformationByUserID(s.ctx, model.InfomationInput{
		UserID: clientIDInt,
		Status: true,
		Value:  clientID,
		Type:   "socket_id",
	})
	if err != nil {
		return
	}
	conn, err := s.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	client := socket.InitClient(global.Hub, conn, clientID, clientIDInt) // 0 is client_id
	log.Println("client: ", client)
	go service.InformationService().SetUserConnected(s.ctx, clientIDInt)
	client.Hub.Register <- client
	go client.WritePump(s.rdb)
	go client.ReadPump(s.rdb)

}

func (s *sSocket) PrivateConnect(c *gin.Context) {
	// implement your code here
}

func (s *sSocket) PublishMessage(channel string, message []byte) error {
	return s.rdb.Publish(s.ctx, string(channel), message).Err()
}
func (s *sSocket) Subscribe(channel string) error {
	_, err := s.rdb.Subscribe(s.ctx, channel).Receive(s.ctx)
	return err
}
