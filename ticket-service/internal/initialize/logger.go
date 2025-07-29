package initialize

import (
	"github.com/quangdat385/holiday-ticket/ticket-service/global"
	"github.com/quangdat385/holiday-ticket/ticket-service/pkg/logger"
)

func InitLogger() {

	global.Logger = logger.NewLogger(global.Config)

}
