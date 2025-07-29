package initialize

import (
	"github.com/quangdat385/holiday-ticket/communications-service/global"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/service/socket"
)

func RunHUb() {
	hub := socket.NewHub()
	global.Hub = hub
	go hub.Run()
}
