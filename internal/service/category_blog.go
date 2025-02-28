package service

import (
	"context"
	"system-management-pg/internal/model"
	"system-management-pg/pkg/response"
)

type (
	ICategoryBlog interface {
		CreateCategoryBlog(ctx context.Context, createCategoryBlog *model.CreateCategoryBlogDto, User *model.UserManagementProfileOutput) (err error, statusCode int)
		UpdateCategoryBlog(ctx context.Context, updateCategoryBlog *model.UpdateCategoryBlogDto, User *model.UserManagementProfileOutput) (err error, statusCode int)
		UpdateStatusCategoryBlog(ctx context.Context, updateStatusCategoryBlog *model.UpdateStatusCategoryBlogDto, User *model.UserManagementProfileOutput) (err error, statusCode int)
		FindCategoryBlogById(ctx context.Context, catBlId string, User *model.UserManagementProfileOutput) (out *model.CategoryBlogOutput, err error, statusCode int)
		GetAllCategoryBlog(ctx context.Context, Limit int32, Offset int32, isDeleted int32,CatName string, User *model.UserManagementProfileOutput) (out response.ModelPagination[[]*model.CategoryBlogOutput], err error, statusCode int)
		DeleteCategoryBlog(ctx context.Context, catBlId string, User *model.UserManagementProfileOutput) (err error, statusCode int)
		RestoreCategoryBlog(ctx context.Context, catBlId string, User *model.UserManagementProfileOutput) (err error, statusCode int)
	}
)

var (
	localCategoryBlog ICategoryBlog
)

func CategoryBlog() ICategoryBlog {
	if localCategoryBlog == nil {
		panic("implement.............................................")
	}
	return localCategoryBlog
}

func InitCategoryBlog(i ICategoryBlog) {
	localCategoryBlog = i
}