package service

import (
	"context"
	"system-management-pg/internal/model"
	"system-management-pg/pkg/response"
)

type (
	IOperationManual interface {
		CreateOperationManual(ctx context.Context, createOperationManual *model.CreateOperationManualDto, Account *model.Account) (err error, statusCode int)
		FindOperationManual(ctx context.Context, OperaManualID string, Account *model.Account) (out *model.OperationManualOutput, err error, statusCode int)
		UpdateOperationManual(ctx context.Context, updateOperationManual *model.UpdateOperationManualDto, Account *model.Account) (err error, statusCode int)
		DeleteOperationManual(ctx context.Context, OperaManualID string, Account *model.Account) (err error, statusCode int)
		RestoreOperationManual(ctx context.Context, OperaManualID string, Account *model.Account) (err error, statusCode int)
		GetAllOperationManual(ctx context.Context, Limit int32, Offset int32, isDeleted int32, OperaManualTitle string, Account *model.Account) (out response.ModelPagination[[]*model.OperationManualOutput], err error, statusCode int)
		UpdateOperationManualStatus(ctx context.Context, updateOperationManualStatus *model.UpdateOperationManualStatusDto, Account *model.Account) (err error, statusCode int)
	}
)

var (
	localOperationManual IOperationManual
)

func OperationManual() IOperationManual {
	if localOperationManual == nil {
		panic("implement.............................................")
	}
	return localOperationManual
}

func InitOperationManual(i IOperationManual) {
	localOperationManual = i
}
