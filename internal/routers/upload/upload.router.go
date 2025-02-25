package upload

import (
	"system-management-pg/internal/controller/upload"
	"system-management-pg/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type UploadRouter struct{}

func (pr *UploadRouter) InitUploadRouter(Router *gin.RouterGroup) {
	uploadPrivate := Router.Group("/upload")
	uploadPrivate.Use(middlewares.AuthenMiddlewareUserManagement())
	{
		uploadPrivate.POST("", upload.Upload.UploadFile)
	}
}
