package service

import (
	"context"
	"system-management-pg/internal/model"
	"system-management-pg/pkg/response"
)

type (
	IBlog interface {
		CreateBlog(ctx context.Context, createBlog *model.CreateBlogDto, User *model.UserManagementProfileOutput) (err error, statusCode int)
		UpdateBlog(ctx context.Context, updateBlog *model.UpdateBlogDto, User *model.UserManagementProfileOutput) (err error, statusCode int)
		UpdateStatusBlog(ctx context.Context, updateStatusBlog *model.UpdateStatusBlogDto, User *model.UserManagementProfileOutput) (err error, statusCode int)
		FindBlogById(ctx context.Context, blId string) (out *model.BlogOutput, err error, statusCode int)
		GetAllBlog(ctx context.Context, Limit int32, Offset int32, isDeleted int32, BlTitle string) (out response.ModelPagination[[]*model.BlogOutput], err error, statusCode int)
		DeleteBlog(ctx context.Context, blId string, User *model.UserManagementProfileOutput) (err error, statusCode int)
		RestoreBlog(ctx context.Context, blId string, User *model.UserManagementProfileOutput) (err error, statusCode int)
	}
)

var (
	localBlog IBlog
)

func Blog() IBlog {
	if localBlog == nil {
		panic("implement.............................................")
	}
	return localBlog
}

func InitBlog(i IBlog) {
	localBlog = i
}