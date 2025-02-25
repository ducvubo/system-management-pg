package usermanagement

import (
	"system-management-pg/internal/model"
	"system-management-pg/internal/service"
	"system-management-pg/internal/utils/context"
	"system-management-pg/internal/utils/validator"
	"system-management-pg/pkg/response"

	"github.com/gin-gonic/gin"
)

var UserManagementAccount = new(cUserManagementAccount)

type cUserManagementAccount struct{}

// CreateUserManagementAccount
// @Summary      CreateUserManagementAccount
// @Description  CreateUserManagementAccount
// @Tags         User Management Account
// @Accept       json
// @Produce      json
// @Param        payload body model.CreateUserManagementAccountDto true "payload"
// @Success      201  {object}  response.ResponseData
// @Router       /user-management-account [post]
func (c *cUserManagementAccount) CreateUserManagementAccount(ctx *gin.Context) {
	var params model.CreateUserManagementAccountDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}

	User := context.GetUserProfileFromCtx(ctx)
	if User == nil {
		response.ErrorResponse(ctx, 401, "Thất bại", nil)
		return
	}
	err, statusCode := service.UserManagementAccount().CreateUserManagementAccount(ctx, &params, User)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Tạo mới thành công", nil)
}

// LoginUserManagementAccount
// @Summary      LoginUserManagementAccount
// @Description  LoginUserManagementAccount
// @Tags         User Management Account
// @Accept       json
// @Produce      json
// @Param        payload body model.LoginUserManagementAccountDto true "payload"
// @Success      200  {object}  response.ResponseData
// @Router       /user-management-account/login [post]
func (c *cUserManagementAccount) LoginUserManagementAccount(ctx *gin.Context) {
	var params model.LoginUserManagementAccountDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}

	clientId := ctx.GetHeader("x-cl-id")
	if clientId == "" {
		ctx.AbortWithStatusJSON(401, gin.H{"code": 40003, "error": "invalid token3", "description": ""})
		return
	}

	data, err, statusCode := service.UserManagementAccount().LoginUserManagementAccount(ctx, &params, clientId)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Đăng nhập thành công", data)
}

// GetProfileUser
// @Summary      GetProfileUser
// @Description  GetProfileUser
// @Tags         User Management Account
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.ResponseData
// @Router       /user-management-account/profile [get]
func (c *cUserManagementAccount) GetProfileUser(ctx *gin.Context) {
	User := context.GetUserProfileFromCtx(ctx)
	if User == nil {
		response.ErrorResponse(ctx, 401, "Thất bại", nil)
	}

	response.SuccessResponse(ctx, 200, "Thành công", User)
}

// RefreshToken
// @Summary      RefreshToken
// @Description  RefreshToken
// @Tags         User Management Account
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.ResponseData
// @Router       /user-management-account/refresh-token [post]
func (c *cUserManagementAccount) RefreshToken(ctx *gin.Context) {
	clientId := ctx.GetHeader("id_user_guest")
	var params model.RefreshTokenInput

	if !validator.BindAndValidate(ctx, &params) {
		return
	}
	data, err, statusCode := service.UserManagementAccount().RefreshToken(ctx, &params, clientId)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Thành công", data)
}
