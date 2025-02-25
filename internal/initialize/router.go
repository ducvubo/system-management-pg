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
	// middlewares
	// r.Use() // logging
	// r.Use() // cross
	// r.Use() // limiter global
	manageRouter := routers.RouterGroupApp.Manage
	userRouter := routers.RouterGroupApp.User
	userManagementProfileRouter := routers.RouterGroupApp.UserManagementProfile
	userManagementAccountRouter := routers.RouterGroupApp.UserManagementAccount
	MainGroup := r.Group("/api/v1")
	{
		MainGroup.GET("/checkStatus") // tracking monitor
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		manageRouter.InitUserRouter(MainGroup)
		manageRouter.InitAdminRouter(MainGroup)
	}
	{
		userManagementProfileRouter.InitUserManagementProfileRouter(MainGroup)
	}
	{
		userManagementAccountRouter.InitUserManagementAccountRouter(MainGroup)
	}
	{
		routers.RouterGroupApp.SystemParameter.InitSystemParameterRouter(MainGroup)
	}
	{
		routers.RouterGroupApp.Upload.InitUploadRouter(MainGroup)
	}

	return r
}
