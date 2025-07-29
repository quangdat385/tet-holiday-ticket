package global

import (
	"database/sql"

	"github.com/quangdat385/holiday-ticket/payment-service/pkg/logger"
	"github.com/quangdat385/holiday-ticket/payment-service/pkg/setting"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
)

var (
	Config        setting.Config
	Logger        *logger.LoggerZap
	Mdb           *sql.DB
	Rdb           *redis.Client
	KafkaProducer *kafka.Writer
)
