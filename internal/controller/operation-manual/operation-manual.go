package operationmanual

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

var OperationManual = new(cOperationManual)

type cOperationManual struct{}

// CreateOperationManual
// @Summary      CreateOperationManual
// @Description  CreateOperationManual
// @Tags         Operation Manual
// @Accept       json
// @Produce      json
// @Param        payload body model.CreateOperationManualDto true "payload"
// @Success      201  {object}  response.ResponseData
// @Router       /operation-manual [post]
func (c *cOperationManual) CreateOperationManual(ctx *gin.Context) {
	var params model.CreateOperationManualDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}
	account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}
	err, statusCode := service.OperationManual().CreateOperationManual(ctx, &params, account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Tạo mới thành công", nil)
}

// FindOperationManual
// @Summary      FindOperationManual
// @Description  FindOperationManual
// @Tags         Operation Manual
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Success      200  {object}  response.ResponseData
// @Router       /operation-manual/{id} [get]
func (c *cOperationManual) FindOperationManual(ctx *gin.Context) {
	id := ctx.Param("id")
	account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}

	data, err, statusCode := service.OperationManual().FindOperationManual(ctx, id,account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Thành công", data)
}

// UpdateOperationManual
// @Summary      UpdateOperationManual
// @Description  UpdateOperationManual
// @Tags         Operation Manual
// @Accept       json
// @Produce      json
// @Param        payload body model.UpdateOperationManualDto true "payload"
// @Success      200  {object}  response.ResponseData
// @Router       /operation-manual [patch]
func (c *cOperationManual) UpdateOperationManual(ctx *gin.Context) {
	var params model.UpdateOperationManualDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}
	account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}
	err, statusCode := service.OperationManual().UpdateOperationManual(ctx, &params, account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Cập nhật thành công", nil)
}

// DeleteOperationManual
// @Summary      DeleteOperationManual
// @Description  DeleteOperationManual
// @Tags         Operation Manual
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Success      200  {object}  response.ResponseData
// @Router       /operation-manual/{id} [delete]
func (c *cOperationManual) DeleteOperationManual(ctx *gin.Context) {
	id := ctx.Param("id")
		account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}
	err, statusCode := service.OperationManual().DeleteOperationManual(ctx, id, account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Xóa thành công", nil)
}

// RestoreOperationManual
// @Summary      RestoreOperationManual
// @Description  RestoreOperationManual
// @Tags         Operation Manual
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Success      200  {object}  response.ResponseData
// @Router       /operation-manual/restore/{id} [patch]
func (c *cOperationManual) RestoreOperationManual(ctx *gin.Context) {
	id := ctx.Param("id")
	account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}
	err, statusCode := service.OperationManual().RestoreOperationManual(ctx, id, account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Khôi phục thành công", nil)
}

// GetAllOperationManual
// @Summary      GetAllOperationManual
// @Description  GetAllOperationManual
// @Tags         Operation Manual
// @Accept       json
// @Produce      json
// @Param        pageIndex query int false "pageIndex"
// @Param        pageSize query int false "pageSize"
// @Param        OperaManualTitle query string false "OperaManualTitle"
// @Success      200  {object}  response.ResponseData
// @Router       /operation-manual [get]
func (c *cOperationManual) GetAllOperationManual(ctx *gin.Context) {
	pageIndexStr := ctx.DefaultQuery("pageIndex", "1")
	pageSizeStr := ctx.DefaultQuery("pageSize", "10")
	OperaManualTitle := ctx.DefaultQuery("OperaManualTitle", "")

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

	account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}

	data, err, statusCode := service.OperationManual().GetAllOperationManual(ctx, limit, offset, 0, OperaManualTitle,account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponseWithPaging(ctx, statusCode, "Thành công", data.Result, int32(pageIndex), int32(pageSize), data.MetaPagination.TotalPage, data.MetaPagination.TotalItems)
}

// GetAllOperationManualRecycle
// @Summary      GetAllOperationManualRecycle
// @Description  GetAllOperationManualRecycle
// @Tags         Operation Manual
// @Accept       json
// @Produce      json
// @Param        pageIndex query int false "pageIndex"
// @Param        pageSize query int false "pageSize"
// @Param        OperaManualTitle query string false "OperaManualTitle"
// @Success      200  {object}  response.ResponseData
// @Router       /operation-manual/recycle [get]
func (c *cOperationManual) GetAllOperationManualRecycle(ctx *gin.Context) {
	pageIndexStr := ctx.DefaultQuery("pageIndex", "1")
	pageSizeStr := ctx.DefaultQuery("pageSize", "10")
	OperaManualTitle := ctx.DefaultQuery("OperaManualTitle", "")

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
	account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}
	data, err, statusCode := service.OperationManual().GetAllOperationManual(ctx, limit, offset, 1, OperaManualTitle,account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponseWithPaging(ctx, statusCode, "Thành công", data.Result, int32(pageIndex), int32(pageSize), data.MetaPagination.TotalPage, data.MetaPagination.TotalItems)
}

// UpdateOperationManualStatus
// @Summary      UpdateOperationManualStatus
// @Description  UpdateOperationManualStatus
// @Tags         Operation Manual
// @Accept       json
// @Produce      json
// @Param        payload body model.UpdateOperationManualStatusDto true "payload"
// @Success      200  {object}  response.ResponseData
// @Router       /operation-manual/update-status [patch]
func (c *cOperationManual) UpdateOperationManualStatus(ctx *gin.Context) {
	var params model.UpdateOperationManualStatusDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}
	account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}
	err, statusCode := service.OperationManual().UpdateOperationManualStatus(ctx, &params, account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Cập nhật thành công", nil)
}