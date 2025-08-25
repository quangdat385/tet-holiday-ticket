package initialize

import (
	"github.com/quangdat385/holiday-ticket/ticket-service/global"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/database"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/service"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/service/impl"
)

func InitService() {
	queries := database.New(global.Mdb)
	// Kiểm tra Redis client có bị nil không
	if global.Rdb == nil {
		panic("global.Rdb is nil! Redis chưa được khởi tạo.")
	}

	// Ticker Service Interface
	// If this service use many services then pls use wire(Section wire)
	redisCache := impl.NewRedisCache(global.Rdb) // Khởi tạo IRedisCache implementation.
	localCache, err := impl.NewRistrettoCache()  // initialize ILocalCache implementation
	if err != nil {
		panic("failed to initialize local cache")
	}
	kafkaConsumer := InitKafkaConsumer()
	if kafkaConsumer == nil {
		panic("failed to initialize kafka consumer")
	}
	// Init ticket item service
	service.InitTicketItem(impl.NewTicketItemImpl(queries, redisCache, localCache))
	// Init ticket home service
	service.InitTicketHome(impl.NewTicketImpl(queries))
	// Init ticket order service
	service.InitStationService(impl.NewStationImpl(queries))
	// Init train service
	service.InitTrainService(impl.NewTrainImpl(queries))
	// Init seat service
	service.InitSeatService(impl.NewSeatImpl(queries))
	// Init seat reservation service
	service.InitSeatReservationService(impl.NewSeatReservation(queries))
	// Init route segment service
	service.InitRouteSegmentService(impl.NewRouteSegmentImpl(queries))
	// Init ticket segment price service
	service.InitTicketSegmentPriceService(impl.NewTicketSegmentPriceImpl(queries))
	// Init kafka consumer service
	service.InitKafkaConsumerService(impl.NewKafkaConsumerServiceImpl(kafkaConsumer, queries))
	go service.KafkaConsumerService().Consume()
}
