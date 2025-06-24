package initialize

import (
	"system-management-pg/global"
	"system-management-pg/internal/database"
	"system-management-pg/internal/service"
	"system-management-pg/internal/service/impl"
)

func InitServiceInterface() {
	queries := database.New(global.Mdbc)
	
	service.InitInternalNote(impl.NewInternalNoteImpl(queries))
	service.InitInternalProposal(impl.NewInternalProposalImpl(queries))
	service.InitEquipmentMaintenance(impl.NewEquipmentMaintenanceImpl(queries))
	service.InitOperationManual(impl.NewOperationManualImpl(queries))
	service.InitOperationalCosts(impl.NewOperationalCostsImpl(queries))
}
