package initialize

import (
	"system-management-pg/global"
	"system-management-pg/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
