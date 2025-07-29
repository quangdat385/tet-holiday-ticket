package initialize

import (
	"github.com/quangdat385/holiday-ticket/payment-service/internal/service"
	"github.com/quangdat385/holiday-ticket/payment-service/internal/service/impl"
)

func InitServiceInterface() {
	service.InitPaymentService(impl.NewPaymentService())
}
