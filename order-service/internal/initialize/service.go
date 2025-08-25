package initialize

import (
	"fmt"

	"github.com/quangdat385/holiday-ticket/order-service/global"
	"github.com/quangdat385/holiday-ticket/order-service/internal/database"
	"github.com/quangdat385/holiday-ticket/order-service/internal/service"
	"github.com/quangdat385/holiday-ticket/order-service/internal/service/impl"
)

func InitService() {
	queries := database.New(global.Mdb)
	// Kiểm tra Redis client có bị nil không
	if global.Rdb == nil {
		panic("global.Rdb is nil! Redis chưa được khởi tạo.")
	}

	// Ticker Service Interface
	// If this service use many services then pls use wire(Section wire)
	kafkaConsumer := InitKafkaConsumer()
	fmt.Println("kafka consumer initialized:", kafkaConsumer.Config().Brokers)
	if kafkaConsumer == nil {
		panic("failed to initialize kafka consumer")
	}
	// Init order service
	service.InitOrderService(impl.NewOrderServiceImpl(queries))
	// Init order detail service
	service.InitOrderDetailService(impl.NewOrderDetailServiceImpl(queries))

	// Init kafka consumer service
	service.InitKafkaConsumerService(impl.NewKafkaConsumerServiceImpl(kafkaConsumer, queries))
	go service.KafkaConsumerService().Consume()
}
