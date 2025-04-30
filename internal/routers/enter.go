package routers

import (
	equipmentmaintenance "system-management-pg/internal/routers/equipment-maintenance"
	internalnote "system-management-pg/internal/routers/internal-note"
	internalproposal "system-management-pg/internal/routers/internal-proposal"
	"system-management-pg/internal/routers/manage"
	systemparameter "system-management-pg/internal/routers/system-parameter"
	"system-management-pg/internal/routers/upload"
	"system-management-pg/internal/routers/user"
	usermanagement "system-management-pg/internal/routers/user-management"
	userpato "system-management-pg/internal/routers/user-pato"
)

type RouterGroup struct {
	User                  user.UserRouterGroup
	Manage                manage.ManageRouterGroup
	UserManagementProfile usermanagement.UserManagementProfileRouter
	UserManagementAccount usermanagement.UserManagementAccountRouter
	SystemParameter       systemparameter.SystemParameterRouter
	Upload                upload.UploadRouter
	UserPatoAccount       userpato.UserPatoAccountRouter
	InternalNote 	      internalnote.InternalNoteRouter
	InternalProposal      internalproposal.InternalProposalRouter
	EquipmentMaintenance  equipmentmaintenance.EquipmentMaintenanceRouter
}

var RouterGroupApp = new(RouterGroup)
