package kafka

import (
	"context"
	"encoding/json"

	"system-management-pg/global"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

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
func SendMessageToKafka(ctx context.Context, topic string, payload NotificationPayload) {
	go func() {
		if global.KafkaProducer == nil {
			global.Logger.Error("Kafka producer not initialized")
			return
		}

		messageBytes, err := json.Marshal(payload)
		if err != nil {
			global.Logger.Error("Failed to marshal Kafka message", zap.Error(err))
			return
		}

		err = global.KafkaProducer.WriteMessages(ctx,
			kafka.Message{
				Topic: topic,
				Value: messageBytes,
			},
		)
		if err != nil {
			global.Logger.Error("Failed to send message to Kafka", zap.String("topic", topic), zap.Error(err))
		} else {
			global.Logger.Info("Message sent to Kafka", zap.String("topic", topic), zap.Any("payload", payload))
		}
	}()
}