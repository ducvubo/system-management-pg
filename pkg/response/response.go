package response

import (
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	StatusCode int         `json:"statusCode"` // HTTP status code
	Code       int         `json:"code"`       // status code
	Message    interface{} `json:"message"`    // thông báo lỗi
	Data       interface{} `json:"data"`       // dữ liệu trả về
	Error      interface{} `json:"error"`      // dữ liệu lỗi
}

type MetaPagination struct {
	Current    int32 `json:"current"`
	PageSize   int32 `json:"pageSize"`
	TotalPage  int32 `json:"totalPage"`
	TotalItems int64 `json:"totalItems"`
}

type ModelPagination[T any] struct {
	Result         T              `json:"result"`
	MetaPagination MetaPagination `json:"meta"`
}

func SuccessResponseWithPaging[T any](c *gin.Context, statusCode int, message string, data T, current int32, pageSize int32, totalPage int32, totalItem int64, codes ...int) {
	code := 0
	if len(codes) > 0 {
		code = codes[0]
	}

	response := ResponseData{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
		Data: ModelPagination[T]{
			Result: data,
			MetaPagination: MetaPagination{
				Current:    current,
				PageSize:   pageSize,
				TotalPage:  totalPage,
				TotalItems: totalItem,
			},
		},
	}

	c.JSON(statusCode, response)
}

func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}, codes ...int) {
	code := 0
	if len(codes) > 0 {
		code = codes[0]
	}
	c.JSON(statusCode, ResponseData{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
		Data:       data,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, message interface{}, err interface{}, codes ...int) {
	code := 0
	if len(codes) > 0 {
		code = codes[0]
	}
	c.JSON(statusCode, ResponseData{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
		Data:       nil,
		Error:      err,
	})
}
