package service

import (
	"context"
	"system-management-pg/internal/model"
	"system-management-pg/pkg/response"
)

type (
	IEquipmentMaintenance interface {
		CreateEquipmentMaintenance(ctx context.Context, createEquipmentMaintenance *model.CreateEquipmentMaintenanceDto, Account *model.Account) (err error, statusCode int)
		FindEquipmentMaintenance(ctx context.Context, EqpMtnID string, Account *model.Account) (out *model.EquipmentMaintenanceOutput, err error, statusCode int)
		UpdateEquipmentMaintenance(ctx context.Context, updateEquipmentMaintenance *model.UpdateEquipmentMaintenanceDto, Account *model.Account) (err error, statusCode int)
		DeleteEquipmentMaintenance(ctx context.Context, EqpMtnID string, Account *model.Account) (err error, statusCode int)
		RestoreEquipmentMaintenance(ctx context.Context, EqpMtnID string, Account *model.Account) (err error, statusCode int)
		GetAllEquipmentMaintenance(ctx context.Context, Limit int32, Offset int32, isDeleted int32, EqpMtnName string, Account *model.Account) (out response.ModelPagination[[]*model.EquipmentMaintenanceOutput], err error, statusCode int)
		UpdateEquipmentMaintenanceStatus(ctx context.Context, updateEquipmentMaintenanceStatus *model.UpdateEquipmentMaintenanceStatusDto, Account *model.Account) (err error, statusCode int)
	}
)

var (
	localEquipmentMaintenance IEquipmentMaintenance
)

func EquipmentMaintenance() IEquipmentMaintenance {
	if localEquipmentMaintenance == nil {
		panic("implement.............................................")
	}
	return localEquipmentMaintenance
}

func InitEquipmentMaintenance(i IEquipmentMaintenance) {
	localEquipmentMaintenance = i
}
