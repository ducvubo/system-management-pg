package routers

import (
	equipmentmaintenance "system-management-pg/internal/routers/equipment-maintenance"
	internalnote "system-management-pg/internal/routers/internal-note"
	internalproposal "system-management-pg/internal/routers/internal-proposal"
	operationmanual "system-management-pg/internal/routers/operation-manual"
	operationalcosts "system-management-pg/internal/routers/operational-costs"
)

type RouterGroup struct {
	InternalNote 	      internalnote.InternalNoteRouter
	InternalProposal      internalproposal.InternalProposalRouter
	EquipmentMaintenance  equipmentmaintenance.EquipmentMaintenanceRouter
	OperationManual 	  operationmanual.OperationManualRouter
	OperationalCosts 	  operationalcosts.OperationalCostsRouter
}

var RouterGroupApp = new(RouterGroup)
