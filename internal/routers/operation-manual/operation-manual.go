package operationmanual

import (
	operationmanual "system-management-pg/internal/controller/operation-manual"
	"system-management-pg/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type OperationManualRouter struct{}

func (pr *OperationManualRouter) InitOperationManualRouter(Router *gin.RouterGroup) {

	operationManualRouterPrivate := Router.Group("/operation-manual")
	operationManualRouterPrivate.Use(middlewares.AuthenMiddlewareAccount())
	{
		operationManualRouterPrivate.POST("", operationmanual.OperationManual.CreateOperationManual)
		operationManualRouterPrivate.GET("/:id", operationmanual.OperationManual.FindOperationManual)
		operationManualRouterPrivate.PATCH("", operationmanual.OperationManual.UpdateOperationManual)
		operationManualRouterPrivate.DELETE("/:id", operationmanual.OperationManual.DeleteOperationManual)
		operationManualRouterPrivate.PATCH("/restore/:id", operationmanual.OperationManual.RestoreOperationManual)
		operationManualRouterPrivate.GET("", operationmanual.OperationManual.GetAllOperationManual)
		operationManualRouterPrivate.GET("/recycle", operationmanual.OperationManual.GetAllOperationManualRecycle)
		operationManualRouterPrivate.PATCH("/update-status", operationmanual.OperationManual.UpdateOperationManualStatus)
	}

}
