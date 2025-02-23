package service

import (
	"context"
	"system-management-pg/internal/model"
	"system-management-pg/pkg/response"
)

type (
	IUserManagementProfile interface {
		CreateUserManagementProfile(ctx context.Context, createUserManagementProfile *model.CreateUserManagementProfileDto, User *model.UserManagementProfileOutput) (err error, statusCode int)
		FindUserManagementProfile(ctx context.Context, UsID string) (out *model.UserManagementProfileOutput, err error, statusCode int)
		UpdateUserManagementProfile(ctx context.Context, updateUserManagementProfile *model.UpdateUserManagementProfileDto, User *model.UserManagementProfileOutput) (err error, statusCode int)
		DeleteUserManagementProfile(ctx context.Context, UsID string, User *model.UserManagementProfileOutput) (err error, statusCode int)
		RestoreUserManagementProfile(ctx context.Context, UsID string, User *model.UserManagementProfileOutput) (err error, statusCode int)
		GetAllUserManagementProfile(ctx context.Context, Limit int32, Offset int32, isDeleted int32, UsName string) (out response.ModelPagination[[]*model.UserManagementProfileOutput], err error, statusCode int)
	}
)

var (
	localUserManagementProfile IUserManagementProfile
)

func UserManagementProfile() IUserManagementProfile {
	if localUserManagementProfile == nil {
		panic("implement.............................................")
	}
	return localUserManagementProfile
}

func InitUserManagementProfile(i IUserManagementProfile) {
	localUserManagementProfile = i
}
