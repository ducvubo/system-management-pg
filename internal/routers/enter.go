package routers

import (
	"system-management-pg/internal/routers/manage"
	systemparameter "system-management-pg/internal/routers/system-parameter"
	"system-management-pg/internal/routers/upload"
	"system-management-pg/internal/routers/user"
	usermanagement "system-management-pg/internal/routers/user-management"
)

type RouterGroup struct {
	User                  user.UserRouterGroup
	Manage                manage.ManageRouterGroup
	UserManagementProfile usermanagement.UserManagementProfileRouter
	UserManagementAccount usermanagement.UserManagementAccountRouter
	SystemParameter       systemparameter.SystemParameterRouter
	Upload                upload.UploadRouter
}

var RouterGroupApp = new(RouterGroup)
