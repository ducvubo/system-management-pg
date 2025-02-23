package usermanagement

import (
	usermanagement "system-management-pg/internal/controller/user-management"
	"system-management-pg/internal/middlewares"

	// "system-management-pg/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type UserManagementAccountRouter struct{}

func (pr *UserManagementAccountRouter) InitUserManagementAccountRouter(Router *gin.RouterGroup) {

	userManagentRouterPublic := Router.Group("/user-management-account")
	{
		userManagentRouterPublic.POST("/login", usermanagement.UserManagementAccount.LoginUserManagementAccount)
		userManagentRouterPublic.POST("/refresh-token", usermanagement.UserManagementAccount.RefreshToken)
	}
	userManagentRouterPrivate := Router.Group("/user-management-account")
	userManagentRouterPrivate.Use(middlewares.AuthenMiddlewareUserManagement())
	{
		userManagentRouterPrivate.POST("", usermanagement.UserManagementAccount.CreateUserManagementAccount)
		userManagentRouterPrivate.GET("/profile", usermanagement.UserManagementAccount.GetProfileUser)
	}

}
