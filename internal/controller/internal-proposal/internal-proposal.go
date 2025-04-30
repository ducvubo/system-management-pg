package internalproposal

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

var InternalProposal = new(cInternalProposal)

type cInternalProposal struct{}

// CreateInternalProposal
// @Summary      CreateInternalProposal
// @Description  CreateInternalProposal
// @Tags         Internal Proposal
// @Accept       json
// @Produce      json
// @Param        payload body model.CreateInternalProposalDto true "payload"
// @Success      201  {object}  response.ResponseData
// @Router       /internal-proposal [post]
func (c *cInternalProposal) CreateInternalProposal(ctx *gin.Context) {
	var params model.CreateInternalProposalDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}
	account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}
	err, statusCode := service.InternalProposal().CreateInternalProposal(ctx, &params, account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Tạo mới thành công", nil)
}

// FindInternalProposal
// @Summary      FindInternalProposal
// @Description  FindInternalProposal
// @Tags         Internal Proposal
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Success      200  {object}  response.ResponseData
// @Router       /internal-proposal/{id} [get]
func (c *cInternalProposal) FindInternalProposal(ctx *gin.Context) {
	id := ctx.Param("id")
	account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}

	data, err, statusCode := service.InternalProposal().FindInternalProposal(ctx, id,account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Thành công", data)
}

// UpdateInternalProposal
// @Summary      UpdateInternalProposal
// @Description  UpdateInternalProposal
// @Tags         Internal Proposal
// @Accept       json
// @Produce      json
// @Param        payload body model.UpdateInternalProposalDto true "payload"
// @Success      200  {object}  response.ResponseData
// @Router       /internal-proposal [patch]
func (c *cInternalProposal) UpdateInternalProposal(ctx *gin.Context) {
	var params model.UpdateInternalProposalDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}
	account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}
	err, statusCode := service.InternalProposal().UpdateInternalProposal(ctx, &params, account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Cập nhật thành công", nil)
}

// DeleteInternalProposal
// @Summary      DeleteInternalProposal
// @Description  DeleteInternalProposal
// @Tags         Internal Proposal
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Success      200  {object}  response.ResponseData
// @Router       /internal-proposal/{id} [delete]
func (c *cInternalProposal) DeleteInternalProposal(ctx *gin.Context) {
	id := ctx.Param("id")
		account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}
	err, statusCode := service.InternalProposal().DeleteInternalProposal(ctx, id, account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Xóa thành công", nil)
}

// RestoreInternalProposal
// @Summary      RestoreInternalProposal
// @Description  RestoreInternalProposal
// @Tags         Internal Proposal
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Success      200  {object}  response.ResponseData
// @Router       /internal-proposal/restore/{id} [patch]
func (c *cInternalProposal) RestoreInternalProposal(ctx *gin.Context) {
	id := ctx.Param("id")
	account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}
	err, statusCode := service.InternalProposal().RestoreInternalProposal(ctx, id, account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Khôi phục thành công", nil)
}

// GetAllInternalProposal
// @Summary      GetAllInternalProposal
// @Description  GetAllInternalProposal
// @Tags         Internal Proposal
// @Accept       json
// @Produce      json
// @Param        pageIndex query int false "pageIndex"
// @Param        pageSize query int false "pageSize"
// @Param        ItnProposalTitle query string false "ItnProposalTitle"
// @Success      200  {object}  response.ResponseData
// @Router       /internal-proposal [get]
func (c *cInternalProposal) GetAllInternalProposal(ctx *gin.Context) {
	pageIndexStr := ctx.DefaultQuery("pageIndex", "1")
	pageSizeStr := ctx.DefaultQuery("pageSize", "10")
	ItnProposalTitle := ctx.DefaultQuery("ItnProposalTitle", "")

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

	data, err, statusCode := service.InternalProposal().GetAllInternalProposal(ctx, limit, offset, 0, ItnProposalTitle,account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponseWithPaging(ctx, statusCode, "Thành công", data.Result, int32(pageIndex), int32(pageSize), data.MetaPagination.TotalPage, data.MetaPagination.TotalItems)
}

// GetAllInternalProposalRecycle
// @Summary      GetAllInternalProposalRecycle
// @Description  GetAllInternalProposalRecycle
// @Tags         Internal Proposal
// @Accept       json
// @Produce      json
// @Param        pageIndex query int false "pageIndex"
// @Param        pageSize query int false "pageSize"
// @Param        ItnProposalTitle query string false "ItnProposalTitle"
// @Success      200  {object}  response.ResponseData
// @Router       /internal-proposal/recycle [get]
func (c *cInternalProposal) GetAllInternalProposalRecycle(ctx *gin.Context) {
	pageIndexStr := ctx.DefaultQuery("pageIndex", "1")
	pageSizeStr := ctx.DefaultQuery("pageSize", "10")
	ItnProposalTitle := ctx.DefaultQuery("ItnProposalTitle", "")

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
	data, err, statusCode := service.InternalProposal().GetAllInternalProposal(ctx, limit, offset, 1, ItnProposalTitle,account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponseWithPaging(ctx, statusCode, "Thành công", data.Result, int32(pageIndex), int32(pageSize), data.MetaPagination.TotalPage, data.MetaPagination.TotalItems)
}

// UpdateInternalProposalStatus
// @Summary      UpdateInternalProposalStatus
// @Description  UpdateInternalProposalStatus
// @Tags         Internal Proposal
// @Accept       json
// @Produce      json
// @Param        payload body model.UpdateInternalProposalStatusDto true "payload"
// @Success      200  {object}  response.ResponseData
// @Router       /internal-proposal/update-status [patch]
func (c *cInternalProposal) UpdateInternalProposalStatus(ctx *gin.Context) {
	var params model.UpdateInternalProposalStatusDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}
	account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}
	err, statusCode := service.InternalProposal().UpdateInternalProposalStatus(ctx, &params, account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Cập nhật thành công", nil)
}