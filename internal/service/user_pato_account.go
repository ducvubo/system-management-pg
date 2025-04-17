package service

import (
	"context"
	"system-management-pg/internal/model"
)

type (
	IUserPatoAccount interface {
		RegisterUserPatoAccount(ctx context.Context, registerUserPatoAccount *model.RegisterUserPatoAccountDto) (registerUserPatoAccountOutput *model.RegisterUserPatoAccountOutput, err error, statusCode int)
		ResendOtp(ctx context.Context, resendOtp *model.ResendOtpDto) (err error, statusCode int)
		ActivateAccount(ctx context.Context, activateAccount *model.ActiveAccountDto) (err error, statusCode int)
		Login(ctx context.Context, login *model.LoginUserPatoAccountDto, clientId string) (loginUserPatoAccountOutput *model.LoginUserPatoAccountOutput, err error, statusCode int)
	}
)

var (
	localUserPatoAccount IUserPatoAccount
)

func UserPatoAccount() IUserPatoAccount {
	if localUserPatoAccount == nil {
		panic("implement.............................................")
	}
	return localUserPatoAccount
}

func InitUserPatoAccount(i IUserPatoAccount) {
	localUserPatoAccount = i
}
