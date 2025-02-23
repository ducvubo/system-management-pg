package usermanagement

import (
	usermanagement "system-management-pg/internal/controller/user-management"

	"github.com/gin-gonic/gin"
)

type UserManagementProfileRouter struct{}

func (pr *UserManagementProfileRouter) InitUserManagementProfileRouter(Router *gin.RouterGroup) {

	userManagentRouterPrivate := Router.Group("/user-management-profile")
	{
		userManagentRouterPrivate.POST("", usermanagement.UserManagementProfile.CreateUserManagementProfile)
		userManagentRouterPrivate.GET("/:id", usermanagement.UserManagementProfile.FindUserManagementProfile)
		userManagentRouterPrivate.PATCH("", usermanagement.UserManagementProfile.UpdateUserManagementProfile)
		userManagentRouterPrivate.DELETE("/:id", usermanagement.UserManagementProfile.DeleteUserManagementProfile)
		userManagentRouterPrivate.PATCH("/restore/:id", usermanagement.UserManagementProfile.RestoreUserManagementProfile)
		userManagentRouterPrivate.GET("", usermanagement.UserManagementProfile.GetAllUserManagementProfile)
		userManagentRouterPrivate.GET("/recycle", usermanagement.UserManagementProfile.GetAllUserManagementProfileRecycle)
	}

}
