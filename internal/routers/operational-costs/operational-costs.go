package operationalcosts

import (
	operationalcosts "system-management-pg/internal/controller/operational-costs"
	"system-management-pg/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type OperationalCostsRouter struct{}

func (pr *OperationalCostsRouter) InitOperationalCostsRouter(Router *gin.RouterGroup) {

	operationalCostsRouterPrivate := Router.Group("/operational-costs")
	operationalCostsRouterPrivate.Use(middlewares.AuthenMiddlewareAccount())
	operationalCostsRouterPrivate.Use(middlewares.LogApiMiddleware())
	{
		operationalCostsRouterPrivate.POST("", operationalcosts.OperationalCosts.CreateOperationalCosts)
		operationalCostsRouterPrivate.GET("/:id", operationalcosts.OperationalCosts.FindOperationalCosts)
		operationalCostsRouterPrivate.PATCH("", operationalcosts.OperationalCosts.UpdateOperationalCosts)
		operationalCostsRouterPrivate.DELETE("/:id", operationalcosts.OperationalCosts.DeleteOperationalCosts)
		operationalCostsRouterPrivate.PATCH("/restore/:id", operationalcosts.OperationalCosts.RestoreOperationalCosts)
		operationalCostsRouterPrivate.GET("", operationalcosts.OperationalCosts.GetAllOperationalCosts)
		operationalCostsRouterPrivate.GET("/recycle", operationalcosts.OperationalCosts.GetAllOperationalCostsRecycle)
		operationalCostsRouterPrivate.PATCH("/update-status", operationalcosts.OperationalCosts.UpdateOperationalCostsStatus)
	}

}
