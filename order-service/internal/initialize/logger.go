package initialize

import (
	"github.com/quangdat385/holiday-ticket/order-service/global"
	"github.com/quangdat385/holiday-ticket/order-service/pkg/logger"
)

func InitLogger() {

	global.Logger = logger.NewLogger(global.Config)

}
