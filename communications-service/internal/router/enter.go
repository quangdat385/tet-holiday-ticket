package router

import (
	"github.com/quangdat385/holiday-ticket/communications-service/internal/router/manager"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/router/socket"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/router/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.ManagerRouterGroup
	Socket  socket.SocketRouterGroup
}

var RouterGroupApp = new(RouterGroup)
