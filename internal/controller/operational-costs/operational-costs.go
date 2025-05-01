package equipmentmaintenance

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

var OperationalCosts = new(cOperationalCosts)

type cOperationalCosts struct{}

// CreateOperationalCosts
// @Summary      CreateOperationalCosts
// @Description  CreateOperationalCosts
// @Tags         Operational Costs
// @Accept       json
// @Produce      json
// @Param        payload body model.CreateOperationalCostsDto true "payload"
// @Success      201  {object}  response.ResponseData
// @Router       /operational-costs [post]
func (c *cOperationalCosts) CreateOperationalCosts(ctx *gin.Context) {
	var params model.CreateOperationalCostsDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}
	account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}
	err, statusCode := service.OperationalCosts().CreateOperationalCosts(ctx, &params, account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Tạo mới thành công", nil)
}

// FindOperationalCosts
// @Summary      FindOperationalCosts
// @Description  FindOperationalCosts
// @Tags         Operational Costs
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Success      200  {object}  response.ResponseData
// @Router       /operational-costs/{id} [get]
func (c *cOperationalCosts) FindOperationalCosts(ctx *gin.Context) {
	id := ctx.Param("id")
	account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}

	data, err, statusCode := service.OperationalCosts().FindOperationalCosts(ctx, id,account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Thành công", data)
}

// UpdateOperationalCosts
// @Summary      UpdateOperationalCosts
// @Description  UpdateOperationalCosts
// @Tags         Operational Costs
// @Accept       json
// @Produce      json
// @Param        payload body model.UpdateOperationalCostsDto true "payload"
// @Success      200  {object}  response.ResponseData
// @Router       /operational-costs [patch]
func (c *cOperationalCosts) UpdateOperationalCosts(ctx *gin.Context) {
	var params model.UpdateOperationalCostsDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}
	account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}
	err, statusCode := service.OperationalCosts().UpdateOperationalCosts(ctx, &params, account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Cập nhật thành công", nil)
}

// DeleteOperationalCosts
// @Summary      DeleteOperationalCosts
// @Description  DeleteOperationalCosts
// @Tags         Operational Costs
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Success      200  {object}  response.ResponseData
// @Router       /operational-costs/{id} [delete]
func (c *cOperationalCosts) DeleteOperationalCosts(ctx *gin.Context) {
	id := ctx.Param("id")
		account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}
	err, statusCode := service.OperationalCosts().DeleteOperationalCosts(ctx, id, account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Xóa thành công", nil)
}

// RestoreOperationalCosts
// @Summary      RestoreOperationalCosts
// @Description  RestoreOperationalCosts
// @Tags         Operational Costs
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Success      200  {object}  response.ResponseData
// @Router       /operational-costs/restore/{id} [patch]
func (c *cOperationalCosts) RestoreOperationalCosts(ctx *gin.Context) {
	id := ctx.Param("id")
	account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}
	err, statusCode := service.OperationalCosts().RestoreOperationalCosts(ctx, id, account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Khôi phục thành công", nil)
}

// GetAllOperationalCosts
// @Summary      GetAllOperationalCosts
// @Description  GetAllOperationalCosts
// @Tags         Operational Costs
// @Accept       json
// @Produce      json
// @Param        pageIndex query int false "pageIndex"
// @Param        pageSize query int false "pageSize"
// @Param        OperaCostType query string false "OperaCostType"
// @Success      200  {object}  response.ResponseData
// @Router       /operational-costs [get]
func (c *cOperationalCosts) GetAllOperationalCosts(ctx *gin.Context) {
	pageIndexStr := ctx.DefaultQuery("pageIndex", "1")
	pageSizeStr := ctx.DefaultQuery("pageSize", "10")
	OperaCostType := ctx.DefaultQuery("OperaCostType", "")

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

	data, err, statusCode := service.OperationalCosts().GetAllOperationalCosts(ctx, limit, offset, 0, OperaCostType,account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponseWithPaging(ctx, statusCode, "Thành công", data.Result, int32(pageIndex), int32(pageSize), data.MetaPagination.TotalPage, data.MetaPagination.TotalItems)
}

// GetAllOperationalCostsRecycle
// @Summary      GetAllOperationalCostsRecycle
// @Description  GetAllOperationalCostsRecycle
// @Tags         Operational Costs
// @Accept       json
// @Produce      json
// @Param        pageIndex query int false "pageIndex"
// @Param        pageSize query int false "pageSize"
// @Param        OperaCostType query string false "OperaCostType"
// @Success      200  {object}  response.ResponseData
// @Router       /operational-costs/recycle [get]
func (c *cOperationalCosts) GetAllOperationalCostsRecycle(ctx *gin.Context) {
	pageIndexStr := ctx.DefaultQuery("pageIndex", "1")
	pageSizeStr := ctx.DefaultQuery("pageSize", "10")
	OperaCostType := ctx.DefaultQuery("OperaCostType", "")

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
	data, err, statusCode := service.OperationalCosts().GetAllOperationalCosts(ctx, limit, offset, 1, OperaCostType,account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponseWithPaging(ctx, statusCode, "Thành công", data.Result, int32(pageIndex), int32(pageSize), data.MetaPagination.TotalPage, data.MetaPagination.TotalItems)
}

// UpdateOperationalCostsStatus
// @Summary      UpdateOperationalCostsStatus
// @Description  UpdateOperationalCostsStatus
// @Tags         Operational Costs
// @Accept       json
// @Produce      json
// @Param        payload body model.UpdateOperationalCostsStatusDto true "payload"
// @Success      200  {object}  response.ResponseData
// @Router       /operational-costs/update-status [patch]
func (c *cOperationalCosts) UpdateOperationalCostsStatus(ctx *gin.Context) {
	var params model.UpdateOperationalCostsStatusDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}
	account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}
	err, statusCode := service.OperationalCosts().UpdateOperationalCostsStatus(ctx, &params, account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Cập nhật thành công", nil)
}