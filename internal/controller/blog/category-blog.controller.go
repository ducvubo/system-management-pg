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

var CategoryBlog = new(cCategoryBlog)

type cCategoryBlog struct{}

// CreateCategoryBlogDto
// @Summary      CreateCategoryBlogDto
// @Description  CreateCategoryBlogDto
// @Tags         Category Blog
// @Accept       json
// @Produce      json
// @Param        payload body model.CreateCategoryBlogDto true "payload"
// @Success      201  {object}  response.ResponseData
// @Router       /category-blog [post]
func (c *cCategoryBlog) CreateCategoryBlogDto(ctx *gin.Context) {
	var params model.CreateCategoryBlogDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}

	User := context.GetUserProfileFromCtx(ctx)
	if User == nil {
		response.ErrorResponse(ctx, 401, "Thất bại", nil)
	}
	err, statusCode := service.CategoryBlog().CreateCategoryBlog(ctx, &params, User)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Tạo mới thành công", nil)
}

// UpdateCategoryBlogDto
// @Summary      UpdateCategoryBlogDto
// @Description  UpdateCategoryBlogDto
// @Tags         Category Blog
// @Accept       json
// @Produce      json
// @Param        payload body model.UpdateCategoryBlogDto true "payload"
// @Success      200  {object}  response.ResponseData
// @Router       /category-blog [patch]
func (c *cCategoryBlog) UpdateCategoryBlogDto(ctx *gin.Context) {
	var params model.UpdateCategoryBlogDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}

	User := context.GetUserProfileFromCtx(ctx)
	if User == nil {
		response.ErrorResponse(ctx, 401, "Thất bại", nil)
	}
	err, statusCode := service.CategoryBlog().UpdateCategoryBlog(ctx, &params, User)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Cập nhật thành công", nil)
}

// UpdateStatusCategoryBlogDto
// @Summary      UpdateStatusCategoryBlogDto
// @Description  UpdateStatusCategoryBlogDto
// @Tags         Category Blog
// @Accept       json
// @Produce      json
// @Param        payload body model.UpdateStatusCategoryBlogDto true "payload"
// @Success      200  {object}  response.ResponseData
// @Router       /category-blog/status [patch]
func (c *cCategoryBlog) UpdateStatusCategoryBlogDto(ctx *gin.Context) {
	var params model.UpdateStatusCategoryBlogDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}

	User := context.GetUserProfileFromCtx(ctx)
	if User == nil {
		response.ErrorResponse(ctx, 401, "Thất bại", nil)
	}
	err, statusCode := service.CategoryBlog().UpdateStatusCategoryBlog(ctx, &params, User)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Cập nhật thành công", nil)
}

// FindCategoryBlogById
// @Summary      FindCategoryBlogById
// @Description  FindCategoryBlogById
// @Tags         Category Blog
// @Accept       json
// @Produce      json
// @Param        catBlId path string true "catBlId"
// @Success      200  {object}  response.ResponseData
// @Router       /category-blog/{catBlId} [get]
func (c *cCategoryBlog) FindCategoryBlogById(ctx *gin.Context) {
	catBlId := ctx.Param("catBlId")
	User := context.GetUserProfileFromCtx(ctx)
	if User == nil {
		response.ErrorResponse(ctx, 401, "Thất bại", nil)
	}
	out, err, statusCode := service.CategoryBlog().FindCategoryBlogById(ctx, catBlId)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Thành công", out)
}

// GetAllCategoryBlog
// @Summary      GetAllCategoryBlog
// @Description  GetAllCategoryBlog
// @Tags         Category Blog
// @Accept       json
// @Produce      json
// @Param        pageIndex query int false "pageIndex"
// @Param        pageSize query int false "pageSize"
// @Param        catName query string false "catName"
// @Success      200  {object}  response.ResponseData
// @Router       /category-blog [get]
func (c *cCategoryBlog) GetAllCategoryBlog(ctx *gin.Context) {
	pageIndexStr := ctx.DefaultQuery("pageIndex", "1")
	pageSizeStr := ctx.DefaultQuery("pageSize", "10")
	catName := ctx.DefaultQuery("catName", "")

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

	data, err, statusCode := service.CategoryBlog().GetAllCategoryBlog(ctx, limit, offset, 0, catName)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponseWithPaging(ctx, statusCode, "Thành công", data.Result, int32(pageIndex), int32(pageSize), data.MetaPagination.TotalPage, data.MetaPagination.TotalItems)
}

// DeleteCategoryBlog
// @Summary      DeleteCategoryBlog
// @Description  DeleteCategoryBlog
// @Tags         Category Blog
// @Accept       json
// @Produce      json
// @Param        catBlId path string true "catBlId"
// @Success      200  {object}  response.ResponseData
// @Router       /category-blog/{catBlId} [delete]
func (c *cCategoryBlog) DeleteCategoryBlog(ctx *gin.Context) {
	catBlId := ctx.Param("catBlId")
	User := context.GetUserProfileFromCtx(ctx)
	if User == nil {
		response.ErrorResponse(ctx, 401, "Thất bại", nil)
	}
	err, statusCode := service.CategoryBlog().DeleteCategoryBlog(ctx, catBlId, User)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Xóa thành công", nil)
}

// RestoreCategoryBlog
// @Summary      RestoreCategoryBlog
// @Description  RestoreCategoryBlog
// @Tags         Category Blog
// @Accept       json
// @Produce      json
// @Param        catBlId path string true "catBlId"
// @Success      200  {object}  response.ResponseData
// @Router       /category-blog/restore/{catBlId} [patch]
func (c *cCategoryBlog) RestoreCategoryBlog(ctx *gin.Context) {
	catBlId := ctx.Param("catBlId")
	User := context.GetUserProfileFromCtx(ctx)
	if User == nil {
		response.ErrorResponse(ctx, 401, "Thất bại", nil)
	}
	err, statusCode := service.CategoryBlog().RestoreCategoryBlog(ctx, catBlId, User)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Khôi phục thành công", nil)
}