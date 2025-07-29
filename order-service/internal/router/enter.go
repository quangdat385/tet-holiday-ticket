package router

import (
	"github.com/quangdat385/holiday-ticket/order-service/internal/router/manager"
	"github.com/quangdat385/holiday-ticket/order-service/internal/router/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
