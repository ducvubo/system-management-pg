package service

import (
	"context"
	"system-management-pg/internal/model"
)

type (
	ISystemParameter interface {
		SaveSystemParameter(ctx context.Context, saveSystemParameterDto *model.SaveSystemParameterDto, User *model.UserManagementProfileOutput) (err error, statusCode int)
		GetSystemParameter(ctx context.Context, sysParaID string) (out *model.SaveSystemParameterDto, err error, statusCode int)
		GetAllSystemParameter(ctx context.Context) (out []*model.SaveSystemParameterDto, err error, statusCode int)
	}
)

var (
	localSystemParameter ISystemParameter
)

func SystemParameter() ISystemParameter {
	if localSystemParameter == nil {
		panic("implement.............................................")
	}
	return localSystemParameter
}

func InitSystemParameter(i ISystemParameter) {
	localSystemParameter = i
}
