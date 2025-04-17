package userpato

import (
	"system-management-pg/internal/model"
	"system-management-pg/internal/service"
	"system-management-pg/internal/utils/validator"
	"system-management-pg/pkg/response"

	"github.com/gin-gonic/gin"
)

var UserPatoAccount = new(cUserPatoAccount)

type cUserPatoAccount struct{}

// CreateUserPatoAccount
// @Summary      CreateUserPatoAccount
// @Description  CreateUserPatoAccount
// @Tags         User Pato Account
// @Accept       json
// @Produce      json
// @Param        payload body model.RegisterUserPatoAccountDto true "payload"
// @Success      201  {object}  response.ResponseData
// @Router       /user-pato-account/register [post]
func (c *cUserPatoAccount) RegisterUserPatoAccount(ctx *gin.Context) {
	var params model.RegisterUserPatoAccountDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}

	user, err, statusCode := service.UserPatoAccount().RegisterUserPatoAccount(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Đăng ký thành công vui lòng xác nhận trong email", user)
}

// ResendOtp
// @Summary      ResendOtp
// @Description  ResendOtp
// @Tags         User Pato Account
// @Accept       json
// @Produce      json
// @Param        payload body model.ResendOtpDto true "payload"
// @Success      201  {object}  response.ResponseData
// @Router       /user-pato-account/resend-otp [post]
func (c *cUserPatoAccount) ResendOtp(ctx *gin.Context) {
	var params model.ResendOtpDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}

	err, statusCode := service.UserPatoAccount().ResendOtp(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Gửi lại mã OTP thành công", nil)
}

// ActivateAccount
// @Summary      ActivateAccount
// @Description  ActivateAccount
// @Tags         User Pato Account
// @Accept       json
// @Produce      json
// @Param        payload body model.ActiveAccountDto true "payload"
// @Success      201  {object}  response.ResponseData
// @Router       /user-pato-account/activate-account [post]
func (c *cUserPatoAccount) ActivateAccount(ctx *gin.Context) {
	var params model.ActiveAccountDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}

	err, statusCode := service.UserPatoAccount().ActivateAccount(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Kích hoạt tài khoản thành công", nil)
}

// Login
// @Summary      Login
// @Description  Login
// @Tags         User Pato Account
// @Accept       json
// @Produce      json
// @Param        payload body model.LoginUserPatoAccountDto true "payload"
// @Success      201  {object}  response.ResponseData
// @Router       /user-pato-account/login [post]
func (c *cUserPatoAccount) Login(ctx *gin.Context) {
	var params model.LoginUserPatoAccountDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}

	user, err, statusCode := service.UserPatoAccount().Login(ctx, &params, ctx.ClientIP())
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Đăng nhập thành công", user)
}
