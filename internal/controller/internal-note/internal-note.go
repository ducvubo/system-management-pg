package internalnote

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

var InternalNote = new(cInternalNote)

type cInternalNote struct{}

// CreateInternalNote
// @Summary      CreateInternalNote
// @Description  CreateInternalNote
// @Tags         Internal Note
// @Accept       json
// @Produce      json
// @Param        payload body model.CreateInternalNoteDto true "payload"
// @Success      201  {object}  response.ResponseData
// @Router       /internal-note [post]
func (c *cInternalNote) CreateInternalNote(ctx *gin.Context) {
	var params model.CreateInternalNoteDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}
	account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}
	err, statusCode := service.InternalNote().CreateInternalNote(ctx, &params, account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Tạo mới thành công", nil)
}

// FindInternalNote
// @Summary      FindInternalNote
// @Description  FindInternalNote
// @Tags         Internal Note
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Success      200  {object}  response.ResponseData
// @Router       /internal-note/{id} [get]
func (c *cInternalNote) FindInternalNote(ctx *gin.Context) {
	id := ctx.Param("id")
	account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}

	data, err, statusCode := service.InternalNote().FindInternalNote(ctx, id,account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Thành công", data)
}

// UpdateInternalNote
// @Summary      UpdateInternalNote
// @Description  UpdateInternalNote
// @Tags         Internal Note
// @Accept       json
// @Produce      json
// @Param        payload body model.UpdateInternalNoteDto true "payload"
// @Success      200  {object}  response.ResponseData
// @Router       /internal-note [patch]
func (c *cInternalNote) UpdateInternalNote(ctx *gin.Context) {
	var params model.UpdateInternalNoteDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}
	account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}
	err, statusCode := service.InternalNote().UpdateInternalNote(ctx, &params, account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Cập nhật thành công", nil)
}

// DeleteInternalNote
// @Summary      DeleteInternalNote
// @Description  DeleteInternalNote
// @Tags         Internal Note
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Success      200  {object}  response.ResponseData
// @Router       /internal-note/{id} [delete]
func (c *cInternalNote) DeleteInternalNote(ctx *gin.Context) {
	id := ctx.Param("id")
		account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}
	err, statusCode := service.InternalNote().DeleteInternalNote(ctx, id, account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Xóa thành công", nil)
}

// RestoreInternalNote
// @Summary      RestoreInternalNote
// @Description  RestoreInternalNote
// @Tags         Internal Note
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Success      200  {object}  response.ResponseData
// @Router       /internal-note/restore/{id} [patch]
func (c *cInternalNote) RestoreInternalNote(ctx *gin.Context) {
	id := ctx.Param("id")
	account := context.GetAccoutFromCtx(ctx)
	if account == nil {
		response.ErrorResponse(ctx, 401, "Không tìm thấy thông tin người dùng", nil)
		return
	}
	err, statusCode := service.InternalNote().RestoreInternalNote(ctx, id, account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Khôi phục thành công", nil)
}

// GetAllInternalNote
// @Summary      GetAllInternalNote
// @Description  GetAllInternalNote
// @Tags         Internal Note
// @Accept       json
// @Produce      json
// @Param        pageIndex query int false "pageIndex"
// @Param        pageSize query int false "pageSize"
// @Param        ItnNoteTitle query string false "ItnNoteTitle"
// @Success      200  {object}  response.ResponseData
// @Router       /internal-note [get]
func (c *cInternalNote) GetAllInternalNote(ctx *gin.Context) {
	pageIndexStr := ctx.DefaultQuery("pageIndex", "1")
	pageSizeStr := ctx.DefaultQuery("pageSize", "10")
	ItnNoteTitle := ctx.DefaultQuery("ItnNoteTitle", "")

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

	data, err, statusCode := service.InternalNote().GetAllInternalNote(ctx, limit, offset, 0, ItnNoteTitle,account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponseWithPaging(ctx, statusCode, "Thành công", data.Result, int32(pageIndex), int32(pageSize), data.MetaPagination.TotalPage, data.MetaPagination.TotalItems)
}

// GetAllInternalNoteRecycle
// @Summary      GetAllInternalNoteRecycle
// @Description  GetAllInternalNoteRecycle
// @Tags         Internal Note
// @Accept       json
// @Produce      json
// @Param        pageIndex query int false "pageIndex"
// @Param        pageSize query int false "pageSize"
// @Param        ItnNoteTitle query string false "ItnNoteTitle"
// @Success      200  {object}  response.ResponseData
// @Router       /internal-note/recycle [get]
func (c *cInternalNote) GetAllInternalNoteRecycle(ctx *gin.Context) {
	pageIndexStr := ctx.DefaultQuery("pageIndex", "1")
	pageSizeStr := ctx.DefaultQuery("pageSize", "10")
	ItnNoteTitle := ctx.DefaultQuery("ItnNoteTitle", "")

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
	data, err, statusCode := service.InternalNote().GetAllInternalNote(ctx, limit, offset, 1, ItnNoteTitle,account)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponseWithPaging(ctx, statusCode, "Thành công", data.Result, int32(pageIndex), int32(pageSize), data.MetaPagination.TotalPage, data.MetaPagination.TotalItems)
}
