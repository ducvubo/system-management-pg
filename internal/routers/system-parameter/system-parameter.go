package systemparameter

import (
	parameter "system-management-pg/internal/controller/system-parameter"
	"system-management-pg/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type SystemParameterRouter struct{}

func (pr *SystemParameterRouter) InitSystemParameterRouter(Router *gin.RouterGroup) {


	systemParameterPublic := Router.Group("/system-parameter")
	systemParameterPublic.Use(middlewares.LogApiMiddleware())
	{
		systemParameterPublic.GET("", parameter.SystemParameter.GetAllSystemParameter)
		systemParameterPublic.GET("/:sysParaID", parameter.SystemParameter.GetSystemParameter)

	}

	systemParameterPrivate := Router.Group("/system-parameter")
	systemParameterPrivate.Use(middlewares.AuthenMiddlewareAccount())
	{
		systemParameterPrivate.POST("/save", parameter.SystemParameter.SaveSystemParameterDto)
	}

}
