package service

type (
	IPaymentService interface {
	}
)

var (
	localIPaymentService IPaymentService
)

func InitPaymentService(i IPaymentService) {
	localIPaymentService = i
}

func PaymentService() IPaymentService {
	if localIPaymentService == nil {
		panic("PaymentService not initialized")
	}
	return localIPaymentService
}
