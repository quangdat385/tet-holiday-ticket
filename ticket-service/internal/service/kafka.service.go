package service

type IKafkaConsumerService interface {
	Consume()
}

var (
	localKafkaConsumerService IKafkaConsumerService
)

func KafkaConsumerService() IKafkaConsumerService {
	if localKafkaConsumerService == nil {
		panic("implement localKafkaConsumerService not found for interface IKafkaConsumerService")
	}

	return localKafkaConsumerService
}
func InitKafkaConsumerService(i IKafkaConsumerService) {
	localKafkaConsumerService = i
}
