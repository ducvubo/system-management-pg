// package initialize

// import (
// 	"fmt"
// 	"time"

// 	"system-management-pg/global"

// 	"github.com/segmentio/kafka-go"
// )

// // Init kafka Producer
// var KafkaProducer *kafka.Writer

// func CheckKafkaConnection(brokerAddress string) error {
// 	conn, err := kafka.Dial("tcp", brokerAddress)
// 	if err != nil {
// 		return fmt.Errorf("không thể kết nối Kafka tại %s: %w", brokerAddress, err)
// 	}
// 	defer conn.Close()

// 	// Ping để kiểm tra kết nối
// 	conn.SetDeadline(time.Now().Add(5 * time.Second))
// 	_, err = conn.Brokers()
// 	if err != nil {
// 		return fmt.Errorf("kết nối được nhưng không lấy được thông tin broker: %w", err)
// 	}
// 	return nil
// }

// func InitKafka() {
// 	// global.KafkaProducer = &kafka.Writer{
// 	// 	Addr:     kafka.TCP("160.187.229.179:19092"),
// 	// 	Topic:    "otp-auth-topic", // topic
// 	// 	Balancer: &kafka.LeastBytes{},
// 	// }
// 	broker := "160.191.243.201:19092"
// 	err := CheckKafkaConnection(broker)
// 	if err != nil {
// 		global.Logger.Info("Initialized Kafka successfully")
// 	}

// 	global.KafkaProducer = &kafka.Writer{
// 		Addr:     kafka.TCP(broker),
// 		Topic:    "otp-auth-topic",
// 		Balancer: &kafka.LeastBytes{},
// 	}
// 	// log.Println("✅ Kafka connection established!")
// 	global.Logger.Info("✅ Kafka connection established!")
// }

// func CloseKafka() {
// 	if err := global.KafkaProducer.Close(); err != nil {
// 		// log.Fatalf("Failed to close kafka producer: %v", err)
// 		global.Logger.Error("Failed to close kafka producer")
// 	}
// }
package initialize

import (
	"fmt"
	"time"

	"system-management-pg/global"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

// InitKafka khởi tạo Kafka Producer
func InitKafka() {
	brokers := []string{
		"160.191.243.201:19092", // Thay bằng process.env.BROKER_KAFKA_1
		// Thêm các broker khác nếu cần: process.env.BROKER_KAFKA_2, BROKER_KAFKA_3
	}

	// Kiểm tra kết nối đến broker
	for _, broker := range brokers {
		err := CheckKafkaConnection(broker)
		if err != nil {
			global.Logger.Error("Failed to connect to Kafka broker", zap.String("broker", broker), zap.Error(err))
			continue
		}
		global.Logger.Info("Connected to Kafka broker", zap.String("broker", broker))
	}

	// Khởi tạo Kafka Producer
	global.KafkaProducer = &kafka.Writer{
		Addr:                   kafka.TCP(brokers...),
		Balancer:               &kafka.LeastBytes{},
		MaxAttempts:            10,                     // Tương đương retry.retries
		WriteTimeout:           5 * time.Second,        // Timeout khi ghi
		RequiredAcks:           kafka.RequireAll,       // Đảm bảo tất cả replica xác nhận
		AllowAutoTopicCreation: true,                   // Tự động tạo topic nếu chưa tồn tại
	}

	global.Logger.Info("✅ Kafka Producer initialized successfully!")
}

// CheckKafkaConnection kiểm tra kết nối đến Kafka broker
func CheckKafkaConnection(brokerAddress string) error {
	conn, err := kafka.Dial("tcp", brokerAddress)
	if err != nil {
		return fmt.Errorf("không thể kết nối Kafka tại %s: %w", brokerAddress, err)
	}
	defer conn.Close()

	conn.SetDeadline(time.Now().Add(5 * time.Second))
	_, err = conn.Brokers()
	if err != nil {
		return fmt.Errorf("kết nối được nhưng không lấy được thông tin broker: %w", err)
	}
	return nil
}

// CloseKafka đóng Kafka Producer
func CloseKafka() {
	if global.KafkaProducer != nil {
		if err := global.KafkaProducer.Close(); err != nil {
			global.Logger.Error("Failed to close Kafka producer", zap.Error(err))
		} else {
			global.Logger.Info("Kafka producer closed successfully")
		}
	}
}