package initialize

import (
	"github.com/quangdat385/holiday-ticket/communications-service/global"
	"github.com/quangdat385/holiday-ticket/communications-service/pkg/logger"
)

func InitLogger() {

	global.Logger = logger.NewLogger(global.Config)

}
