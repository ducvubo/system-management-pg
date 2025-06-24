package kafka

import (
	"context"
	"encoding/json"

	"system-management-pg/global"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

// func SendMessageToKafka(topic string, message string) error {
// 	writer := &kafka.Writer{
// 		Addr:         kafka.TCP("160.187.229.179:19092"),
// 		Balancer:     &kafka.LeastBytes{},
// 		RequiredAcks: kafka.RequireOne,
// 		Async:        false,
// 	}
// 	defer writer.Close()

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	err := writer.WriteMessages(ctx,
// 		kafka.Message{
// 			Topic: topic,
// 			Value: []byte(message),
// 		},
// 	)

// 	if err != nil {
// 		return fmt.Errorf("failed to send message to Kafka: %w", err)
// 	}

// 	fmt.Printf("📤 Tin nhắn đã được gửi tới topic [%s]: %s\n", topic, message)
// 	return nil
// }


type KafkaMessage struct {
	Topic   string
	Message string
}

// NotificationPayload định nghĩa payload cho thông báo
type NotificationPayload struct {
	RestaurantID  string `json:"restaurantId"`
	NotiContent   string `json:"noti_content"`
	NotiTitle     string `json:"noti_title"`
	NotiType      string `json:"noti_type"`
	NotiMetadata  string `json:"noti_metadata"`
	SendObject    string `json:"sendObject"`
}

// SendMessageToKafka gửi tin nhắn đến Kafka topic bất đồng bộ
func SendMessageToKafka(ctx context.Context, msg KafkaMessage) {
	go func() {
		if global.KafkaProducer == nil {
			global.Logger.Error("Kafka producer not initialized")
			return
		}

		messageBytes, err := json.Marshal(msg.Message)
		if err != nil {
			global.Logger.Error("Failed to marshal Kafka message", zap.Error(err))
			return
		}

		err = global.KafkaProducer.WriteMessages(ctx,
			kafka.Message{
				Topic: msg.Topic,
				Value: messageBytes,
			},
		)
		if err != nil {
			global.Logger.Error("Failed to send message to Kafka", zap.String("topic", msg.Topic), zap.Error(err))
		} else {
			global.Logger.Info("Message sent to Kafka", zap.String("topic", msg.Topic), zap.String("message", msg.Message))
		}
	}()
}
