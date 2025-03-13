package blog

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

var Blog = new(cBlog)

type cBlog struct{}

// CreateBlogDto
// @Summary      CreateBlogDto
// @Description  CreateBlogDto
// @Tags         Blog
// @Accept       json
// @Produce      json
// @Param        payload body model.CreateBlogDto true "payload"
// @Success      201  {object}  response.ResponseData
// @Router       /blog [post]
func (c *cBlog) CreateBlogDto(ctx *gin.Context) {
	var params model.CreateBlogDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}

	User := context.GetUserProfileFromCtx(ctx)
	if User == nil {
		response.ErrorResponse(ctx, 401, "Thất bại", nil)
	}
	err, statusCode := service.Blog().CreateBlog(ctx, &params, User)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Tạo mới thành công", nil)
}

// UpdateBlogDto
// @Summary      UpdateBlogDto
// @Description  UpdateBlogDto
// @Tags         Blog
// @Accept       json
// @Produce      json
// @Param        payload body model.UpdateBlogDto true "payload"
// @Success      200  {object}  response.ResponseData
// @Router       /blog [patch]
func (c *cBlog) UpdateBlogDto(ctx *gin.Context) {
	var params model.UpdateBlogDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}

	User := context.GetUserProfileFromCtx(ctx)
	if User == nil {
		response.ErrorResponse(ctx, 401, "Thất bại", nil)
	}
	err, statusCode := service.Blog().UpdateBlog(ctx, &params, User)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Cập nhật thành công", nil)
}

// UpdateStatusBlogDto
// @Summary      UpdateStatusBlogDto
// @Description  UpdateStatusBlogDto
// @Tags         Blog
// @Accept       json
// @Produce      json
// @Param        payload body model.UpdateStatusBlogDto true "payload"
// @Success      200  {object}  response.ResponseData
// @Router       /blog/status [patch]
func (c *cBlog) UpdateStatusBlogDto(ctx *gin.Context) {
	var params model.UpdateStatusBlogDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}

	User := context.GetUserProfileFromCtx(ctx)
	if User == nil {
		response.ErrorResponse(ctx, 401, "Thất bại", nil)
	}
	err, statusCode := service.Blog().UpdateStatusBlog(ctx, &params, User)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Cập nhật thành công", nil)
}

// FindBlogById
// @Summary      FindBlogById
// @Description  FindBlogById
// @Tags         Blog
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Success      200  {object}  response.ResponseData
// @Router       /blog/{id} [get]
func (c *cBlog) FindBlogById(ctx *gin.Context) {
	id := ctx.Param("id")

	data, err, statusCode := service.Blog().FindBlogById(ctx, id)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Thành công", data)
}

// GetAllBlog
// @Summary      GetAllBlog
// @Description  GetAllBlog
// @Tags         Blog
// @Accept       json
// @Produce      json
// @Param        pageIndex query int false "pageIndex"
// @Param        pageSize query int false "pageSize"
// @Param        blTitle query string false "blTitle"
// @Success      200  {object}  response.ResponseData
// @Router       /blog [get]
func (c *cBlog) GetAllBlog(ctx *gin.Context) {
pageIndexStr := ctx.DefaultQuery("pageIndex", "1")
	pageSizeStr := ctx.DefaultQuery("pageSize", "10")
	blTitle := ctx.DefaultQuery("blTitle", "")

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

	data, err, statusCode := service.Blog().GetAllBlog(ctx, limit, offset, 0, blTitle)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponseWithPaging(ctx, statusCode, "Thành công", data.Result, int32(pageIndex), int32(pageSize), data.MetaPagination.TotalPage, data.MetaPagination.TotalItems)
}

// DeleteBlog
// @Summary      DeleteBlog
// @Description  DeleteBlog
// @Tags         Blog
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Success      200  {object}  response.ResponseData
// @Router       /blog/{id} [delete]
func (c *cBlog) DeleteBlog(ctx *gin.Context) {
	id := ctx.Param("id")

	User := context.GetUserProfileFromCtx(ctx)
	if User == nil {
		response.ErrorResponse(ctx, 401, "Thất bại", nil)
	}
	err, statusCode := service.Blog().DeleteBlog(ctx, id, User)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Xóa thành công", nil)
}

// RestoreBlog
// @Summary      RestoreBlog
// @Description  RestoreBlog
// @Tags         Blog
// @Accept       json
// @Produce      json
// @Param        id path string true "id"
// @Success      200  {object}  response.ResponseData
// @Router       /blog/restore/{id} [patch]
func (c *cBlog) RestoreBlog(ctx *gin.Context) {
	id := ctx.Param("id")

	User := context.GetUserProfileFromCtx(ctx)
	if User == nil {
		response.ErrorResponse(ctx, 401, "Thất bại", nil)
	}
	err, statusCode := service.Blog().RestoreBlog(ctx, id, User)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Khôi phục thành công", nil)
}

// GetBlogByCatBlId
// @Summary      GetBlogByCatBlId
// @Description  GetBlogByCatBlId
// @Tags         Blog
// @Accept       json
// @Produce      json
// @Param        catBlId path string true "catBlId"
// @Success      200  {object}  response.ResponseData
// @Router       /blog/category/{catBlId} [get]
func (c *cBlog) GetBlogByCatBlId(ctx *gin.Context) {
	catBlId := ctx.Param("catBlId")

	data, err, statusCode := service.Blog().FindBlogById(ctx, catBlId)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Thành công", data)
}

