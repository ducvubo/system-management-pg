package service

import (
	"context"
	"system-management-pg/internal/model"
	"system-management-pg/pkg/response"
)

type (
	IOperationalCosts interface {
		CreateOperationalCosts(ctx context.Context, createOperationalCosts *model.CreateOperationalCostsDto, Account *model.Account) (err error, statusCode int)
		FindOperationalCosts(ctx context.Context, OperaCostID string, Account *model.Account) (out *model.OperationalCostsOutput, err error, statusCode int)
		UpdateOperationalCosts(ctx context.Context, updateOperationalCosts *model.UpdateOperationalCostsDto, Account *model.Account) (err error, statusCode int)
		DeleteOperationalCosts(ctx context.Context, OperaCostID string, Account *model.Account) (err error, statusCode int)
		RestoreOperationalCosts(ctx context.Context, OperaCostID string, Account *model.Account) (err error, statusCode int)
		GetAllOperationalCosts(ctx context.Context, Limit int32, Offset int32, isDeleted int32, OperaCostType string, Account *model.Account) (out response.ModelPagination[[]*model.OperationalCostsOutput], err error, statusCode int)
		UpdateOperationalCostsStatus(ctx context.Context, updateOperationalCostsStatus *model.UpdateOperationalCostsStatusDto, Account *model.Account) (err error, statusCode int)
	}
)

var (
	localOperationalCosts IOperationalCosts
)

func OperationalCosts() IOperationalCosts {
	if localOperationalCosts == nil {
		panic("implement.............................................")
	}
	return localOperationalCosts
}

func InitOperationalCosts(i IOperationalCosts) {
	localOperationalCosts = i
}
