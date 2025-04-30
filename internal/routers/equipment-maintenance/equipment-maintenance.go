package equipmentmaintenance

import (
	equipmentmaintenance "system-management-pg/internal/controller/equipment-maintenance"
	"system-management-pg/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type EquipmentMaintenanceRouter struct{}

func (pr *EquipmentMaintenanceRouter) InitEquipmentMaintenanceRouter(Router *gin.RouterGroup) {

	equipmentMaintenanceRouterPrivate := Router.Group("/equipment-maintenance")
	equipmentMaintenanceRouterPrivate.Use(middlewares.AuthenMiddlewareAccount())
	{
		equipmentMaintenanceRouterPrivate.POST("", equipmentmaintenance.EquipmentMaintenance.CreateEquipmentMaintenance)
		equipmentMaintenanceRouterPrivate.GET("/:id", equipmentmaintenance.EquipmentMaintenance.FindEquipmentMaintenance)
		equipmentMaintenanceRouterPrivate.PATCH("", equipmentmaintenance.EquipmentMaintenance.UpdateEquipmentMaintenance)
		equipmentMaintenanceRouterPrivate.DELETE("/:id", equipmentmaintenance.EquipmentMaintenance.DeleteEquipmentMaintenance)
		equipmentMaintenanceRouterPrivate.PATCH("/restore/:id", equipmentmaintenance.EquipmentMaintenance.RestoreEquipmentMaintenance)
		equipmentMaintenanceRouterPrivate.GET("", equipmentmaintenance.EquipmentMaintenance.GetAllEquipmentMaintenance)
		equipmentMaintenanceRouterPrivate.GET("/recycle", equipmentmaintenance.EquipmentMaintenance.GetAllEquipmentMaintenanceRecycle)
		equipmentMaintenanceRouterPrivate.PATCH("/update-status", equipmentmaintenance.EquipmentMaintenance.UpdateEquipmentMaintenanceStatus)
	}

}
