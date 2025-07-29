package service

import "github.com/gin-gonic/gin"

type (
	ISocketService interface {
		Connect(c *gin.Context)
		PublishMessage(channel string, message []byte) error
		Subscribe(channel string) error
	}
)

var (
	localSocketService ISocketService
)

func SocketService() ISocketService {
	if localSocketService == nil {
		panic("implement localSocketService not found for interface ISocketService")
	}
	return localSocketService
}
func InitSocketService(socketService ISocketService) {
	localSocketService = socketService
}
