package initialize

import (
	"github.com/quangdat385/holiday-ticket/communications-service/global"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/database"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/service"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/service/impl"
)

func InitSocketService() {
	service.InitSocketService(impl.NewSocketImpl(global.Rdb))
}
func InitMessageService() {
	queries := database.New(global.Mdb)
	redisCache := impl.NewRedisCache(global.Rdb)

	service.InitInformationService(impl.NewInformationImpl(queries, redisCache))
	service.InitConversationService(impl.NewConversationImpl(queries))
	service.InitMessageService(impl.NewMessageServiceImpl(queries))
	service.InitNotificationService(impl.NewNotificationServiceImpl(queries))
}
