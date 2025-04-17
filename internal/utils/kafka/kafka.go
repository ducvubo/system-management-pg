package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func SendMessageToKafka(topic string, message string) error {
	writer := &kafka.Writer{
		Addr:         kafka.TCP("160.187.229.179:19092"),
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireOne,
		Async:        false,
	}
	defer writer.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := writer.WriteMessages(ctx,
		kafka.Message{
			Topic: topic,
			Value: []byte(message),
		},
	)

	if err != nil {
		return fmt.Errorf("failed to send message to Kafka: %w", err)
	}

	fmt.Printf("📤 Tin nhắn đã được gửi tới topic [%s]: %s\n", topic, message)
	return nil
}
