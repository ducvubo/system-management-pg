package initialize

import (
	"fmt"
	"system-management-pg/global"
	consts "system-management-pg/internal/const"
	"system-management-pg/internal/database"
	"system-management-pg/internal/service"
	"system-management-pg/internal/service/impl"
)

func InitServiceInterface() {
	param := consts.SystemEmail

	fmt.Println("param: %v", param)
	queries := database.New(global.Mdbc)
	// User Service Interface
	service.InitUserLogin(impl.NewUserLoginImpl(queries))
	service.InitUserManagementProfile(impl.NewUserManagementProfileImpl(queries))
	service.InitUserManagementAccount(impl.NewUserManagementAccountImpl(queries))
	service.InitSystemParameter(impl.NewSystemParameterImpl(queries))

}
