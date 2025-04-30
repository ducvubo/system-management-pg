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

var EquipmentMaintenance = new(cEquipmentMaintenance)

type cEquipmentMaintenance struct{}

// CreateEquipmentMaintenance
// @Summary      CreateEquipmentMaintenance
// @Description  CreateEquipmentMaintenance
// @Tags         Equipment Maintenance
// @Accept       json
// @Produce      json
// @Param        payload body model.CreateEquipmentMaintenanceDto true "payload"
// @Success      201  {object}  response.ResponseData
// @Router       /equipment-maintenance [post]
func (c *cEquipmentMaintenance) CreateEquipmentMaintenance(ctx *gin.Context) {
	var params model.CreateEquipmentMaintenanceDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}
	account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}
	err, statusCode := service.EquipmentMaintenance().CreateEquipmentMaintenance(ctx, &params, account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Tạo mới thành công", nil)
}

// FindEquipmentMaintenance
// @Summary      FindEquipmentMaintenance
// @Description  FindEquipmentMaintenance
// @Tags         Equipment Maintenance
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Success      200  {object}  response.ResponseData
// @Router       /equipment-maintenance/{id} [get]
func (c *cEquipmentMaintenance) FindEquipmentMaintenance(ctx *gin.Context) {
	id := ctx.Param("id")
	account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}

	data, err, statusCode := service.EquipmentMaintenance().FindEquipmentMaintenance(ctx, id,account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Thành công", data)
}

// UpdateEquipmentMaintenance
// @Summary      UpdateEquipmentMaintenance
// @Description  UpdateEquipmentMaintenance
// @Tags         Equipment Maintenance
// @Accept       json
// @Produce      json
// @Param        payload body model.UpdateEquipmentMaintenanceDto true "payload"
// @Success      200  {object}  response.ResponseData
// @Router       /equipment-maintenance [patch]
func (c *cEquipmentMaintenance) UpdateEquipmentMaintenance(ctx *gin.Context) {
	var params model.UpdateEquipmentMaintenanceDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}
	account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}
	err, statusCode := service.EquipmentMaintenance().UpdateEquipmentMaintenance(ctx, &params, account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Cập nhật thành công", nil)
}

// DeleteEquipmentMaintenance
// @Summary      DeleteEquipmentMaintenance
// @Description  DeleteEquipmentMaintenance
// @Tags         Equipment Maintenance
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Success      200  {object}  response.ResponseData
// @Router       /equipment-maintenance/{id} [delete]
func (c *cEquipmentMaintenance) DeleteEquipmentMaintenance(ctx *gin.Context) {
	id := ctx.Param("id")
		account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}
	err, statusCode := service.EquipmentMaintenance().DeleteEquipmentMaintenance(ctx, id, account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Xóa thành công", nil)
}

// RestoreEquipmentMaintenance
// @Summary      RestoreEquipmentMaintenance
// @Description  RestoreEquipmentMaintenance
// @Tags         Equipment Maintenance
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Success      200  {object}  response.ResponseData
// @Router       /equipment-maintenance/restore/{id} [patch]
func (c *cEquipmentMaintenance) RestoreEquipmentMaintenance(ctx *gin.Context) {
	id := ctx.Param("id")
	account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}
	err, statusCode := service.EquipmentMaintenance().RestoreEquipmentMaintenance(ctx, id, account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Khôi phục thành công", nil)
}

// GetAllEquipmentMaintenance
// @Summary      GetAllEquipmentMaintenance
// @Description  GetAllEquipmentMaintenance
// @Tags         Equipment Maintenance
// @Accept       json
// @Produce      json
// @Param        pageIndex query int false "pageIndex"
// @Param        pageSize query int false "pageSize"
// @Param        EqpMtnName query string false "EqpMtnName"
// @Success      200  {object}  response.ResponseData
// @Router       /equipment-maintenance [get]
func (c *cEquipmentMaintenance) GetAllEquipmentMaintenance(ctx *gin.Context) {
	pageIndexStr := ctx.DefaultQuery("pageIndex", "1")
	pageSizeStr := ctx.DefaultQuery("pageSize", "10")
	EqpMtnName := ctx.DefaultQuery("EqpMtnName", "")

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

	data, err, statusCode := service.EquipmentMaintenance().GetAllEquipmentMaintenance(ctx, limit, offset, 0, EqpMtnName,account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponseWithPaging(ctx, statusCode, "Thành công", data.Result, int32(pageIndex), int32(pageSize), data.MetaPagination.TotalPage, data.MetaPagination.TotalItems)
}

// GetAllEquipmentMaintenanceRecycle
// @Summary      GetAllEquipmentMaintenanceRecycle
// @Description  GetAllEquipmentMaintenanceRecycle
// @Tags         Equipment Maintenance
// @Accept       json
// @Produce      json
// @Param        pageIndex query int false "pageIndex"
// @Param        pageSize query int false "pageSize"
// @Param        EqpMtnName query string false "EqpMtnName"
// @Success      200  {object}  response.ResponseData
// @Router       /equipment-maintenance/recycle [get]
func (c *cEquipmentMaintenance) GetAllEquipmentMaintenanceRecycle(ctx *gin.Context) {
	pageIndexStr := ctx.DefaultQuery("pageIndex", "1")
	pageSizeStr := ctx.DefaultQuery("pageSize", "10")
	EqpMtnName := ctx.DefaultQuery("EqpMtnName", "")

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
	data, err, statusCode := service.EquipmentMaintenance().GetAllEquipmentMaintenance(ctx, limit, offset, 1, EqpMtnName,account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponseWithPaging(ctx, statusCode, "Thành công", data.Result, int32(pageIndex), int32(pageSize), data.MetaPagination.TotalPage, data.MetaPagination.TotalItems)
}

// UpdateEquipmentMaintenanceStatus
// @Summary      UpdateEquipmentMaintenanceStatus
// @Description  UpdateEquipmentMaintenanceStatus
// @Tags         Equipment Maintenance
// @Accept       json
// @Produce      json
// @Param        payload body model.UpdateEquipmentMaintenanceStatusDto true "payload"
// @Success      200  {object}  response.ResponseData
// @Router       /equipment-maintenance/update-status [patch]
func (c *cEquipmentMaintenance) UpdateEquipmentMaintenanceStatus(ctx *gin.Context) {
	var params model.UpdateEquipmentMaintenanceStatusDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}
	account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}
	err, statusCode := service.EquipmentMaintenance().UpdateEquipmentMaintenanceStatus(ctx, &params, account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Cập nhật thành công", nil)
}