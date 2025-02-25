package upload

import (
	"context"
	"fmt"
	"path/filepath"
	contextUtil "system-management-pg/internal/utils/context"
	"system-management-pg/pkg/response"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var Upload = new(cUpload)

type cUpload struct{}

// ImageResponse định nghĩa struct chứa hai URL
type ImageResponse struct {
	ImageCustom string `json:"image_custom"`
	ImageCloud  string `json:"image_cloud"`
}

// UploadFile
// @Summary      Upload a file to Cloudinary
// @Description  Upload an image file to Cloudinary with a specified folder from header and return original and transformed URLs
// @Tags         File Upload
// @Accept       multipart/form-data
// @Produce      json
// @Param        file    formData  file   true  "File to upload (jpg, jpeg, png, gif)"
// @Param        folder  header    string false "Folder name in Cloudinary (optional)"
// @Success      200  {object}  response.ResponseData{data=ImageResponse}  "Upload successful with original and transformed file URLs"
// @Failure      400  {object}  response.ResponseData  "Invalid file or request"
// @Failure      500  {object}  response.ResponseData  "Server or Cloudinary error"
// @Router       /upload [post]
func (c *cUpload) UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		response.ErrorResponse(ctx, 400, "Thất bại", "Không tìm thấy file")
		return
	}

	if file.Size > 1024*1024*1024 {
		response.ErrorResponse(ctx, 400, "Thất bại", "File quá lớn")
		return
	}

	fileExtension := filepath.Ext(file.Filename)
	if fileExtension != ".jpg" && fileExtension != ".jpeg" && fileExtension != ".png" && fileExtension != ".gif" {
		response.ErrorResponse(ctx, 400, "Thất bại", "File không đúng định dạng")
		return
	}

	fileName := uuid.New().String()

	User := contextUtil.GetUserProfileFromCtx(ctx)
	if User == nil {
		response.ErrorResponse(ctx, 401, "Thất bại", nil)
		return
	}

	folder := ctx.GetHeader("folder")
	if folder == "" {
		folder = User.UsID
	}

	// Khởi tạo Cloudinary
	cld, err := cloudinary.NewFromParams("dqdc3gvh6", "512436567492936", "WwEJfE3U_CCwzu0B8v0ydpbHxVM")
	if err != nil {
		response.ErrorResponse(ctx, 500, "Thất bại", "Không thể kết nối tới Cloudinary")
		return
	}

	// Mở file để upload
	fileReader, err := file.Open()
	if err != nil {
		response.ErrorResponse(ctx, 400, "Thất bại", "Không thể đọc file")
		return
	}
	defer fileReader.Close()

	// Upload file lên Cloudinary (file gốc)
	uploadResult, err := cld.Upload.Upload(context.Background(), fileReader, uploader.UploadParams{
		PublicID: fileName,
		Folder:   folder,
	})
	if err != nil {
		response.ErrorResponse(ctx, 500, "Thất bại", "Không thể upload file lên Cloudinary")
		return
	}

	// URL gốc
	originalURL := uploadResult.SecureURL

	// Tạo URL custom thủ công (height: 100, width: 100, fetch_format: jpg)
	customURL := fmt.Sprintf(
		"https://res.cloudinary.com/%s/image/upload/h_100,w_100,f_jpg/%s/%s",
		cld.Config.Cloud.CloudName, // Lấy cloud_name từ config
		folder,
		fileName,
	)

	// Trả về phản hồi với 2 URL trong struct ImageResponse
	response.SuccessResponse(ctx, 201, "Thành công", ImageResponse{
		ImageCustom: customURL,
		ImageCloud:  originalURL,
	})
}
