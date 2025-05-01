package initialize

import (
	"system-management-pg/global"
	"system-management-pg/internal/database"
	"system-management-pg/internal/service"
	"system-management-pg/internal/service/impl"
)

func InitServiceInterface() {
	// param := consts.SystemEmail

	// fmt.Println("param: %v", param)
	queries := database.New(global.Mdbc)
	// User Service Interface
	service.InitUserLogin(impl.NewUserLoginImpl(queries))
	service.InitUserManagementProfile(impl.NewUserManagementProfileImpl(queries))
	service.InitUserManagementAccount(impl.NewUserManagementAccountImpl(queries))
	service.InitSystemParameter(impl.NewSystemParameterImpl(queries))
	service.InitUserPatoAccount(impl.NewUserPatoAccountImpl(queries))
	service.InitInternalNote(impl.NewInternalNoteImpl(queries))
	service.InitInternalProposal(impl.NewInternalProposalImpl(queries))
	service.InitEquipmentMaintenance(impl.NewEquipmentMaintenanceImpl(queries))
	service.InitOperationManual(impl.NewOperationManualImpl(queries))
	service.InitOperationalCosts(impl.NewOperationalCostsImpl(queries))
}
