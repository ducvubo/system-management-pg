package global

import (
	"database/sql"

	"system-management-pg/pkg/logger"
	"system-management-pg/pkg/setting"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

var (
	Config        setting.Config
	Logger        *logger.LoggerZap
	Rdb           *redis.Client
	Mdb           *gorm.DB
	Mdbc          *sql.DB
	KafkaProducer *kafka.Writer
)
