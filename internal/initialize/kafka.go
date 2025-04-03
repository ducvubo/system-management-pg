package initialize

import (
	"log"

	"system-management-pg/global"
	"github.com/segmentio/kafka-go"
)

// Init kafka Producer
var KafkaProducer *kafka.Writer

func InitKafka() {
	global.KafkaProducer = &kafka.Writer{
		Addr:     kafka.TCP("160.187.229.179:19092"),
		Topic:    "otp-auth-topic", // topic
		Balancer: &kafka.LeastBytes{},
	}
}

func CloseKafka() {
	if err := global.KafkaProducer.Close(); err != nil {
		log.Fatalf("Failed to close kafka producer: %v", err)
	}
}
