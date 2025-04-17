package userpato

import (
	userpato "system-management-pg/internal/controller/user-pato"

	// "system-management-pg/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type UserPatoAccountRouter struct{}

func (pr *UserPatoAccountRouter) InitUserPatoAccountRouter(Router *gin.RouterGroup) {

	userManagentRouterPublic := Router.Group("/user-pato-account")
	{
		userManagentRouterPublic.POST("/register", userpato.UserPatoAccount.RegisterUserPatoAccount)
		userManagentRouterPublic.POST("/resend-otp", userpato.UserPatoAccount.ResendOtp)
		userManagentRouterPublic.POST("/activate-account", userpato.UserPatoAccount.ActivateAccount)
		userManagentRouterPublic.POST("/login", userpato.UserPatoAccount.Login)
	}

}
