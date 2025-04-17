package initialize

import (
	"fmt"
	"time"

	"system-management-pg/global"

	"github.com/segmentio/kafka-go"
)

// Init kafka Producer
var KafkaProducer *kafka.Writer

func CheckKafkaConnection(brokerAddress string) error {
	conn, err := kafka.Dial("tcp", brokerAddress)
	if err != nil {
		return fmt.Errorf("không thể kết nối Kafka tại %s: %w", brokerAddress, err)
	}
	defer conn.Close()

	// Ping để kiểm tra kết nối
	conn.SetDeadline(time.Now().Add(5 * time.Second))
	_, err = conn.Brokers()
	if err != nil {
		return fmt.Errorf("kết nối được nhưng không lấy được thông tin broker: %w", err)
	}
	return nil
}

func InitKafka() {
	// global.KafkaProducer = &kafka.Writer{
	// 	Addr:     kafka.TCP("160.187.229.179:19092"),
	// 	Topic:    "otp-auth-topic", // topic
	// 	Balancer: &kafka.LeastBytes{},
	// }
	broker := "160.187.229.179:19092"
	err := CheckKafkaConnection(broker)
	if err != nil {
		global.Logger.Info("Initialized Kafka successfully")
	}

	global.KafkaProducer = &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    "otp-auth-topic",
		Balancer: &kafka.LeastBytes{},
	}
	// log.Println("✅ Kafka connection established!")
	global.Logger.Info("✅ Kafka connection established!")
}

func CloseKafka() {
	if err := global.KafkaProducer.Close(); err != nil {
		// log.Fatalf("Failed to close kafka producer: %v", err)
		global.Logger.Error("Failed to close kafka producer")
	}
}
