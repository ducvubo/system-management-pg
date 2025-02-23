package usermanagement

import (
	"net/http"
	"strconv"
	"system-management-pg/internal/model"
	"system-management-pg/internal/service"
	"system-management-pg/internal/utils/context"
	"system-management-pg/internal/utils/validator"
	"system-management-pg/pkg/response"

	"github.com/gin-gonic/gin"
)

var UserManagementProfile = new(cUserManagementProfile)

type cUserManagementProfile struct{}

// CreateUserManagementProfile
// @Summary      CreateUserManagementProfile
// @Description  CreateUserManagementProfile
// @Tags         User Management Profile
// @Accept       json
// @Produce      json
// @Param        payload body model.CreateUserManagementProfileDto true "payload"
// @Success      201  {object}  response.ResponseData
// @Router       /user-management-profile [post]
func (c *cUserManagementProfile) CreateUserManagementProfile(ctx *gin.Context) {
	var params model.CreateUserManagementProfileDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}
	User := context.GetUserProfileFromCtx(ctx)
	if User == nil {
		response.ErrorResponse(ctx, 401, "Thất bại", nil)
	}
	err, statusCode := service.UserManagementProfile().CreateUserManagementProfile(ctx, &params, User)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Tạo mới thành công", nil)
}

// FindUserManagementProfile
// @Summary      FindUserManagementProfile
// @Description  FindUserManagementProfile
// @Tags         User Management Profile
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Success      200  {object}  response.ResponseData
// @Router       /user-management-profile/{id} [get]
func (c *cUserManagementProfile) FindUserManagementProfile(ctx *gin.Context) {
	id := ctx.Param("id")

	data, err, statusCode := service.UserManagementProfile().FindUserManagementProfile(ctx, id)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Thành công", data)
}

// UpdateUserManagementProfile
// @Summary      UpdateUserManagementProfile
// @Description  UpdateUserManagementProfile
// @Tags         User Management Profile
// @Accept       json
// @Produce      json
// @Param        payload body model.UpdateUserManagementProfileDto true "payload"
// @Success      200  {object}  response.ResponseData
// @Router       /user-management-profile [patch]
func (c *cUserManagementProfile) UpdateUserManagementProfile(ctx *gin.Context) {
	var params model.UpdateUserManagementProfileDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}
	User := context.GetUserProfileFromCtx(ctx)
	if User == nil {
		response.ErrorResponse(ctx, 401, "Thất bại", nil)
	}
	err, statusCode := service.UserManagementProfile().UpdateUserManagementProfile(ctx, &params, User)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Cập nhật thành công", nil)
}

// DeleteUserManagementProfile
// @Summary      DeleteUserManagementProfile
// @Description  DeleteUserManagementProfile
// @Tags         User Management Profile
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Success      200  {object}  response.ResponseData
// @Router       /user-management-profile/{id} [delete]
func (c *cUserManagementProfile) DeleteUserManagementProfile(ctx *gin.Context) {
	id := ctx.Param("id")
	User := context.GetUserProfileFromCtx(ctx)
	if User == nil {
		response.ErrorResponse(ctx, 401, "Thất bại", nil)
	}
	err, statusCode := service.UserManagementProfile().DeleteUserManagementProfile(ctx, id, User)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Xóa thành công", nil)
}

// RestoreUserManagementProfile
// @Summary      RestoreUserManagementProfile
// @Description  RestoreUserManagementProfile
// @Tags         User Management Profile
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Success      200  {object}  response.ResponseData
// @Router       /user-management-profile/restore/{id} [patch]
func (c *cUserManagementProfile) RestoreUserManagementProfile(ctx *gin.Context) {
	id := ctx.Param("id")
	User := context.GetUserProfileFromCtx(ctx)
	if User == nil {
		response.ErrorResponse(ctx, 401, "Thất bại", nil)
	}
	err, statusCode := service.UserManagementProfile().RestoreUserManagementProfile(ctx, id, User)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Khôi phục thành công", nil)
}

// GetAllUserManagementProfile
// @Summary      GetAllUserManagementProfile
// @Description  GetAllUserManagementProfile
// @Tags         User Management Profile
// @Accept       json
// @Produce      json
// @Param        pageIndex query int false "pageIndex"
// @Param        pageSize query int false "pageSize"
// @Param        usName query string false "usName"
// @Success      200  {object}  response.ResponseData
// @Router       /user-management-profile [get]
func (c *cUserManagementProfile) GetAllUserManagementProfile(ctx *gin.Context) {
	pageIndexStr := ctx.DefaultQuery("pageIndex", "1")
	pageSizeStr := ctx.DefaultQuery("pageSize", "10")
	usName := ctx.DefaultQuery("usName", "")

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		response.ErrorResponse(ctx, http.StatusBadRequest, "Giá trị pageSize không hợp lệ", nil)
		return
	}

	pageIndex, err := strconv.Atoi(pageIndexStr)
	if err != nil || pageIndex < 1 {
		response.ErrorResponse(ctx, http.StatusBadRequest, "Giá trị pageIndex không hợp lệ", nil)
		return
	}

	limit := int32(pageSize)
	offset := int32((pageIndex - 1) * pageSize)

	data, err, statusCode := service.UserManagementProfile().GetAllUserManagementProfile(ctx, limit, offset, 0, usName)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponseWithPaging(ctx, statusCode, "Thành công", data.Result, int32(pageIndex), int32(pageSize), data.MetaPagination.TotalPage, data.MetaPagination.TotalItems)
}

// GetAllUserManagementProfileRecycle
// @Summary      GetAllUserManagementProfileRecycle
// @Description  GetAllUserManagementProfileRecycle
// @Tags         User Management Profile
// @Accept       json
// @Produce      json
// @Param        pageIndex query int false "pageIndex"
// @Param        pageSize query int false "pageSize"
// @Param        usName query string false "usName"
// @Success      200  {object}  response.ResponseData
// @Router       /user-management-profile/recycle [get]
func (c *cUserManagementProfile) GetAllUserManagementProfileRecycle(ctx *gin.Context) {
	pageIndexStr := ctx.DefaultQuery("pageIndex", "1")
	pageSizeStr := ctx.DefaultQuery("pageSize", "10")
	usName := ctx.DefaultQuery("usName", "")

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		response.ErrorResponse(ctx, http.StatusBadRequest, "Giá trị pageSize không hợp lệ", nil)
		return
	}

	pageIndex, err := strconv.Atoi(pageIndexStr)
	if err != nil || pageIndex < 1 {
		response.ErrorResponse(ctx, http.StatusBadRequest, "Giá trị pageIndex không hợp lệ", nil)
		return
	}

	limit := int32(pageSize)
	offset := int32((pageIndex - 1) * pageSize)

	data, err, statusCode := service.UserManagementProfile().GetAllUserManagementProfile(ctx, limit, offset, 0, usName)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponseWithPaging(ctx, statusCode, "Thành công", data.Result, int32(pageIndex), int32(pageSize), data.MetaPagination.TotalPage, data.MetaPagination.TotalItems)
}
