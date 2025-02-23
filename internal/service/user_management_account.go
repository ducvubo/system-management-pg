package service

import (
	"context"
	"system-management-pg/internal/database"
	"system-management-pg/internal/model"
)

type (
	IUserManagementAccount interface {
		CreateUserManagementAccount(ctx context.Context, createUserManagementAccount *model.CreateUserManagementAccountDto, User *model.UserManagementProfileOutput) (err error, statusCode int)
		FindUserAccountById(ctx context.Context, UsaID string) (out *model.UserManagementAccountOutput, err error, statusCode int)
		LoginUserManagementAccount(ctx context.Context, loginUserManagementAccount *model.LoginUserManagementAccountDto, clientId string) (loginUserManagementAccountOutput *model.LoginUserManagementAccountOutput, err error, statusCode int)
		FindUserSessionBySessionIdAndRefreshToken(ctx context.Context, clientId, refreshToken string) (userSession database.FindUserSessionBySessionIdAndRefreshTokenRow, err error, statusCode int)
		RefreshToken(ctx context.Context, refreshTokenInput *model.RefreshTokenInput, clientId string) (loginUserManagementAccountOutput *model.LoginUserManagementAccountOutput, err error, statusCode int)
	}
)

var (
	localUserManagementAccount IUserManagementAccount
)

func UserManagementAccount() IUserManagementAccount {
	if localUserManagementAccount == nil {
		panic("implement.............................................")
	}
	return localUserManagementAccount
}

func InitUserManagementAccount(i IUserManagementAccount) {
	localUserManagementAccount = i
}
