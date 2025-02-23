package validator

import (
	"fmt"
	"net/http"
	"strings"
	"system-management-pg/pkg/response"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Map lỗi sang thông báo tiếng Việt
var errorMessages = map[string]string{
	"required": "%s không được để trống",
	"min":      "%s phải có ít nhất %s ký tự",
	"max":      "%s không được vượt quá %s ký tự",
	"e164":     "%s không hợp lệ, phải đúng định dạng số điện thoại quốc tế",
	"oneof":    "%s phải là 'male', 'female' hoặc 'other'",
	"datetime": "%s không đúng định dạng YYYY-MM-DD",
	"url":      "%s phải là một đường dẫn hợp lệ",
}

// Chuyển lỗi validation thành danh sách chuỗi
func GetValidationErrors(err error) []string {
	var errors []string
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			fieldName := e.Field() // Tên trường bị lỗi
			tag := e.Tag()         // Kiểu lỗi (vd: required, min, max,...)

			// Kiểm tra nếu lỗi có trong danh sách custom
			if msg, exists := errorMessages[tag]; exists {
				if tag == "min" || tag == "max" {
					errors = append(errors, fmt.Sprintf(msg, fieldName, e.Param()))
				} else {
					errors = append(errors, fmt.Sprintf(msg, fieldName))
				}
			} else {
				// Nếu không có lỗi custom, trả về lỗi mặc định
				errors = append(errors, fmt.Sprintf("Field %s có lỗi: %s", fieldName, tag))
			}
		}
	}
	return errors
}

// Hàm Bind và Validate dữ liệu
func BindAndValidate(ctx *gin.Context, requestData interface{}) bool {
	// Validate dữ liệu JSON
	if err := ctx.ShouldBindJSON(requestData); err != nil {
		validationErrors := GetValidationErrors(err)

		// Trả về response lỗi
		response.ErrorResponse(ctx, http.StatusBadRequest, validationErrors, "Dữ liệu không hợp lệ")
		return false
	}
	return true
}

func IsStrongPassword(password string) bool {
	var hasUpper, hasLower, hasDigit, hasSpecial bool
	specialChars := "!@#$%^&*()-_=+[]{}|;:'\",.<>?/"

	if len(password) < 8 {
		return false
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case strings.ContainsRune(specialChars, char):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasDigit && hasSpecial
}
