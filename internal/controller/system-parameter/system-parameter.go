package systemparameter

import (
	"system-management-pg/internal/model"
	"system-management-pg/internal/service"
	"system-management-pg/internal/utils/context"
	"system-management-pg/internal/utils/validator"
	"system-management-pg/pkg/response"

	"github.com/gin-gonic/gin"
)

var SystemParameter = new(cSystemParameter)

type cSystemParameter struct{}

// SaveSystemParameterDto
// @Summary      SaveSystemParameterDto
// @Description  SaveSystemParameterDto
// @Tags         System Parameter
// @Accept       json
// @Produce      json
// @Param        payload body model.SaveSystemParameterDto true "payload"
// @Success      201  {object}  response.ResponseData
// @Router       /system-parameter/save [post]
func (c *cSystemParameter) SaveSystemParameterDto(ctx *gin.Context) {
	var params model.SaveSystemParameterDto

	if !validator.BindAndValidate(ctx, &params) {
		return
	}

	User := context.GetUserProfileFromCtx(ctx)
	if User == nil {
		response.ErrorResponse(ctx, 401, "Thất bại", nil)
	}
	err, statusCode := service.SystemParameter().SaveSystemParameter(ctx, &params, User)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Tạo mới thành công", nil)
}

// GetSystemParameter
// @Summary      GetSystemParameter
// @Description  GetSystemParameter
// @Tags         System Parameter
// @Accept       json
// @Produce      json
// @Param        sysParaID path string true "sysParaID"
// @Success      200  {object}  response.ResponseData
// @Router       /system-parameter/{sysParaID} [get]
func (c *cSystemParameter) GetSystemParameter(ctx *gin.Context) {
	sysParaID := ctx.Param("sysParaID")

	out, err, statusCode := service.SystemParameter().GetSystemParameter(ctx, sysParaID)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Thành công", out)
}

// GetAllSystemParameter
// @Summary      GetAllSystemParameter
// @Description  GetAllSystemParameter
// @Tags         System Parameter
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.ResponseData
// @Router       /system-parameter [get]
func (c *cSystemParameter) GetAllSystemParameter(ctx *gin.Context) {
	out, err, statusCode := service.SystemParameter().GetAllSystemParameter(ctx)
	if err != nil {
		response.ErrorResponse(ctx, statusCode, "Đã có lỗi xảy ra", err.Error())
		return
	}
	response.SuccessResponse(ctx, statusCode, "Thành công", out)
}
