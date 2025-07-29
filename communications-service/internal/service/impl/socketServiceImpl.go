package impl

import (
	"context"
	"log"
	"net/http"
	"strconv"

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
	conn, err := s.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	auth, _ := c.Get("UserID")
	client_id, err := strconv.Atoi(auth.(string))
	if err != nil {
		log.Println("Error converting client_id:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client_id"})
		return
	}
	log.Println("client_id: ", client_id)
	client := socket.InitClient(global.Hub, conn, utils.GenerateRandomString(10), int64(client_id)) // 0 is client_id
	log.Println("client: ", client)
	result, err := service.InformationService().GetInformationByUserID(s.ctx, int64(client_id))
	if err != nil {
		log.Println("Error fetching communication info:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch communication info"})
		return
	}
	if result.ID == 0 {
		_, err = service.InformationService().InsertInformationByUserID(s.ctx, model.InfomationInput{
			UserID: int64(client_id),
			Status: true,
			Value:  client.Id,
			Type:   "socket_id",
		})
		if err != nil {
			log.Println("Error inserting communication info:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert communication info"})
			return
		}
	}
	_, err = service.InformationService().UpdateInformationByUserID(s.ctx, model.InfomationInput{
		UserID: int64(client_id),
		Status: true,
		Value:  client.Id,
		Type:   "socket_id",
	})
	if err != nil {
		log.Println("Error updating communication info:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update communication info"})
		return
	}
	_, err = service.InformationService().SetUserConnected(s.ctx, int64(client_id))
	if err != nil {
		log.Println("Error setting user connected status:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set userconnected status"})
		return
	}
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
