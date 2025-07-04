package initialize

import (
	"system-management-pg/global"
	"system-management-pg/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Hoặc domain cụ thể
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, id_user_guest")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	MainGroup := r.Group("/api/v1")
	{
		MainGroup.GET("/checkStatus") // tracking monitor
	}
	{
		routers.RouterGroupApp.InternalNote.InitInternalNoteRouter(MainGroup)
	}
	{
		routers.RouterGroupApp.InternalProposal.InitInternalProposalRouter(MainGroup)
	}
	{
		routers.RouterGroupApp.EquipmentMaintenance.InitEquipmentMaintenanceRouter(MainGroup)
	}
	{
		routers.RouterGroupApp.OperationManual.InitOperationManualRouter(MainGroup)
	}
	{
		routers.RouterGroupApp.OperationalCosts.InitOperationalCostsRouter(MainGroup)
	}

	return r
}
