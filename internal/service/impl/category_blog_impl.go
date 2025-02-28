package impl

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"system-management-pg/internal/database"
	"system-management-pg/internal/model"
	"system-management-pg/internal/utils"
	"system-management-pg/pkg/response"

	"github.com/google/uuid"
)

// CreateCategoryBlog(ctx context.Context, createCategoryBlog *model.CreateCategoryBlogDto) (err error, statusCode int)
// UpdateCategoryBlog(ctx context.Context, updateCategoryBlog *model.UpdateCategoryBlogDto) (err error, statusCode int)
// UpdateStatusCategoryBlog(ctx context.Context, updateStatusCategoryBlog *model.UpdateStatusCategoryBlogDto) (err error, statusCode int)
// FindCategoryBlogById(ctx context.Context, catBlId string) (out *model.CategoryBlogOutput, err error, statusCode int)
// GetAllCategoryBlog(ctx context.Context, Limit int32, Offset int32, isDeleted int32) (out response.ModelPagination[[]*model.CategoryBlogOutput], err error, statusCode int)
// DeleteCategoryBlog(ctx context.Context, catBlId string) (err error, statusCode int)
// RestoreCategoryBlog(ctx context.Context, catBlId string) (err error, statusCode int)

type sCategoryBlog struct {
	r *database.Queries
}

func NewCategoryBlogImpl(r *database.Queries) *sCategoryBlog {
	return &sCategoryBlog{
		r: r,
	}
}

func (s *sCategoryBlog) CreateCategoryBlog(ctx context.Context, createCategoryBlog *model.CreateCategoryBlogDto, User *model.UserManagementProfileOutput) (err error, statusCode int) {
	slug := utils.CreateSlug(createCategoryBlog.CatBlName)
	_, err = s.r.CreateCategoryBlog(ctx, database.CreateCategoryBlogParams{
		CatBlID: uuid.New().String(),
		CatBlName: createCategoryBlog.CatBlName,
		CatBlSlug: slug,
		CatBlDescription: sql.NullString{String: createCategoryBlog.CatBlDesc, Valid: true},
		Createdby: sql.NullString{ String: User.UsID, Valid: true},
	})

	if err != nil {
		return fmt.Errorf("failed to create category blog: %w", err), http.StatusBadRequest
	}

	return nil, http.StatusCreated

}

func (s *sCategoryBlog) UpdateCategoryBlog(ctx context.Context, updateCategoryBlog *model.UpdateCategoryBlogDto, User *model.UserManagementProfileOutput) (err error, statusCode int) {
	_, err = s.r.GetCategoryBlog(ctx, updateCategoryBlog.CatBlId)
	if err != nil {
		return fmt.Errorf("failed to get category blog: %w", err), http.StatusBadRequest
	}

	slug := utils.CreateSlug(updateCategoryBlog.CatBlName)
	 err = s.r.UpdateCategoryBlog(ctx, database.UpdateCategoryBlogParams{
		CatBlName: updateCategoryBlog.CatBlName,
		CatBlSlug: slug,
		CatBlDescription: sql.NullString{String: updateCategoryBlog.CatBlDesc, Valid: true},
		CatBlOrder: sql.NullInt32{Int32: updateCategoryBlog.CatBlOrder, Valid: true},
		Updatedby: sql.NullString{ String: User.UsID, Valid: true},
		CatBlID: updateCategoryBlog.CatBlId,
	})

	if err != nil {
		return fmt.Errorf("failed to update category blog: %w", err), http.StatusBadRequest
	}

	return nil, http.StatusCreated
}

func (s *sCategoryBlog) UpdateStatusCategoryBlog(ctx context.Context, updateStatusCategoryBlog *model.UpdateStatusCategoryBlogDto, User *model.UserManagementProfileOutput) (err error, statusCode int) {
	_, err = s.r.GetCategoryBlog(ctx, updateStatusCategoryBlog.CatBlId)
	if err != nil {
		return fmt.Errorf("failed to get category blog: %w", err), http.StatusBadRequest
	}

	err = s.r.UpdateStatusCategoryBlog(ctx, database.UpdateStatusCategoryBlogParams{
		CatBlStatus: sql.NullInt32{Int32: updateStatusCategoryBlog.CatBlStatus, Valid: true},
		Updatedby: sql.NullString{ String: User.UsID, Valid: true},
		CatBlID: updateStatusCategoryBlog.CatBlId,
	})

	if err != nil {
		return fmt.Errorf("failed to update status category blog: %w", err), http.StatusBadRequest
	}
	return nil, http.StatusCreated
}

func (s *sCategoryBlog) FindCategoryBlogById(ctx context.Context, catBlId string, User *model.UserManagementProfileOutput) (out *model.CategoryBlogOutput, err error, statusCode int) {
	Cat, err := s.r.GetCategoryBlog(ctx, catBlId)
	if err != nil {
		return nil, fmt.Errorf("failed to get category blog: %w", err), http.StatusBadRequest
	}
	return &model.CategoryBlogOutput{
		CatBlId: Cat.CatBlID,
		CatBlName: Cat.CatBlName,
		CatBlSlug: Cat.CatBlSlug,
		CatBlDesc: Cat.CatBlDescription.String,
		CatBlOrder: Cat.CatBlOrder.Int32,
		CatBlStatus: Cat.CatBlStatus.Int32,
	}, nil, http.StatusOK
}

func (s *sCategoryBlog) GetAllCategoryBlog(ctx context.Context, Limit int32, Offset int32, isDeleted int32,CatName string, User *model.UserManagementProfileOutput) (out response.ModelPagination[[]*model.CategoryBlogOutput], err error, statusCode int) {
	Cat, err := s.r.GetListCategoryBlog(ctx, database.GetListCategoryBlogParams{
		Limit: Limit,
		Offset: Offset,
		Isdeleted: sql.NullInt32{Int32: isDeleted, Valid: true},
		CatBlName: "%" + CatName + "%",
		Isdeleted_2: sql.NullInt32{Int32: isDeleted, Valid: false},
		CatBlName_2: "%" + CatName + "%",
		Column3: float64(Limit),

	})
	if err != nil {
		return response.ModelPagination[[]*model.CategoryBlogOutput]{}, fmt.Errorf("failed to get all category blog: %w", err), http.StatusBadRequest
	}
	var catBlog []*model.CategoryBlogOutput
	for _, v := range Cat {
		catBlog = append(catBlog, &model.CategoryBlogOutput{
			CatBlId: v.CatBlID,
			CatBlName: v.CatBlName,
			CatBlSlug: v.CatBlSlug,
			CatBlDesc: v.CatBlDescription.String,
			CatBlOrder: v.CatBlOrder.Int32,
			CatBlStatus: v.CatBlStatus.Int32,
		})
	}
	var totalPages, totalItems int32 = 0, 0
	if len(Cat) > 0 {
		totalPages = Cat[0].TotalPages
		totalItems = int32(Cat[0].TotalItems)
	}
	return response.ModelPagination[[]*model.CategoryBlogOutput]{
		Result: catBlog,
		MetaPagination: response.MetaPagination{
			Current: Offset,
			PageSize: Limit,
			TotalPage: totalPages,
			TotalItems: int64(totalItems),
		},
	}, nil, http.StatusOK
}

func (s *sCategoryBlog) DeleteCategoryBlog(ctx context.Context, catBlId string, User *model.UserManagementProfileOutput) (err error, statusCode int) {
	_, err = s.r.GetCategoryBlog(ctx, catBlId)
	if err != nil {
		return fmt.Errorf("failed to get category blog: %w", err), http.StatusBadRequest
	}

	err = s.r.DeleteCategoryBlog(ctx, database.DeleteCategoryBlogParams{
		CatBlID: catBlId,
		Deletedby: sql.NullString{ String: User.UsID, Valid: true},
	})
	if err != nil {
		return fmt.Errorf("failed to delete category blog: %w", err), http.StatusBadRequest
	}
	return nil, http.StatusCreated
}

func (s *sCategoryBlog) RestoreCategoryBlog(ctx context.Context, catBlId string, User *model.UserManagementProfileOutput) (err error, statusCode int) {
	_, err = s.r.GetCategoryBlog(ctx, catBlId)
	if err != nil {
		return fmt.Errorf("failed to get category blog: %w", err), http.StatusBadRequest
	}

	err = s.r.RestoreCategoryBlog(ctx, catBlId)
	if err != nil {
		return fmt.Errorf("failed to restore category blog: %w", err), http.StatusBadRequest
	}
	return nil, http.StatusCreated
}
