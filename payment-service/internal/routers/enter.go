package routers

import (
	"github.com/quangdat385/holiday-ticket/payment-service/internal/routers/manager"
	"github.com/quangdat385/holiday-ticket/payment-service/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
