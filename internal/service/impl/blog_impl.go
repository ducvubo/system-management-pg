package impl

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"system-management-pg/internal/database"
	"system-management-pg/internal/model"
	"system-management-pg/internal/utils"

	"github.com/google/uuid"
)

// CreateCategoryBlog(ctx context.Context, createCategoryBlog *model.CreateCategoryBlogDto, User *model.UserManagementProfileOutput) (err error, statusCode int)
// 	UpdateCategoryBlog(ctx context.Context, updateCategoryBlog *model.UpdateCategoryBlogDto, User *model.UserManagementProfileOutput) (err error, statusCode int)
// 	UpdateStatusCategoryBlog(ctx context.Context, updateStatusCategoryBlog *model.UpdateStatusCategoryBlogDto, User *model.UserManagementProfileOutput) (err error, statusCode int)
// 	FindCategoryBlogById(ctx context.Context, catBlId string, User *model.UserManagementProfileOutput) (out *model.CategoryBlogOutput, err error, statusCode int)
// 	GetAllCategoryBlog(ctx context.Context, Limit int32, Offset int32, isDeleted int32,CatName string, User *model.UserManagementProfileOutput) (out response.ModelPagination[[]*model.CategoryBlogOutput], err error, statusCode int)
// 	DeleteCategoryBlog(ctx context.Context, catBlId string, User *model.UserManagementProfileOutput) (err error, statusCode int)
// 	RestoreCategoryBlog(ctx context.Context, catBlId string, User *model.UserManagementProfileOutput) (err error, statusCode int)



type sBlog struct {
	r *database.Queries
}

func NewBlogImpl(r *database.Queries) *sBlog {
	return &sBlog{
		r: r,
	}
}

func (s *sBlog) CreateBlog(ctx context.Context, createBlog *model.CreateBlogDto, User *model.UserManagementProfileOutput) (err error, statusCode int) {
	slug := utils.CreateSlug(createBlog.BlTitle)
	BlID := uuid.New().String()
	_, err = s.r.CreateBlog(ctx, database.CreateBlogParams{
		BlID: BlID,
		CatBlID: createBlog.CatBlID,
		BlTitle: createBlog.BlTitle,
		BlDescription: sql.NullString{String: createBlog.BlDescription, Valid: true},
		BlSlug: slug,
		BlImage: sql.NullString{String: createBlog.BlImage, Valid: true},
		BlContent: createBlog.BlContent,
		BlType: sql.NullInt32{Int32: createBlog.BlType, Valid: true},
		BlView: sql.NullInt32{Int32: 0, Valid: true},
		BlStatus: sql.NullInt32{Int32: 1, Valid: true},
		Updatedby: sql.NullString{String: User.UsID, Valid: true},
		Createdby: sql.NullString{ String: User.UsID, Valid: true},
	})

	// _,err = s.r.CreateBlogNote(ctx, database.CreateBlogNoteParams{
	// 	BlID: BlID,
	// 	BlNoteID: uuid.New().String(),
	// 	BlNoteContent: ,
	// })
	 

	if err != nil {
		return fmt.Errorf("failed to create blog: %w", err), http.StatusBadRequest
	}
	return nil, http.StatusCreated
}

func (s *sBlog) UpdateBlog(ctx context.Context, updateBlog *model.UpdateBlogDto, User *model.UserManagementProfileOutput) (err error, statusCode int) {
	_, err = s.r.GetBlogByID(ctx, updateBlog.BlID)
	if err != nil {
		return fmt.Errorf("failed to get blog: %w", err), http.StatusBadRequest
	}
	slug := utils.CreateSlug(updateBlog.BlTitle)
	err = s.r.UpdateBlog(ctx, database.UpdateBlogParams{
		BlID: updateBlog.BlID,
		CatBlID: updateBlog.CatBlID,
		BlTitle: updateBlog.BlTitle,
		BlDescription: sql.NullString{String: updateBlog.BlDescription, Valid: true},
		BlSlug: slug,
		BlImage: sql.NullString{String: updateBlog.BlImage, Valid: true},
		BlContent: updateBlog.BlContent,
		Updatedby: sql.NullString{String: User.UsID, Valid: true},
	})
	if err != nil {
		return fmt.Errorf("failed to update blog: %w", err), http.StatusBadRequest
	}
	return nil, http.StatusOK
}

func (s *sBlog) UpdateStatusBlog(ctx context.Context, updateStatusBlog *model.UpdateStatusBlogDto, User *model.UserManagementProfileOutput) (err error, statusCode int) {
	_, err = s.r.GetBlogByID(ctx, updateStatusBlog.BlID)
	if err != nil {
		return fmt.Errorf("failed to get blog: %w", err), http.StatusBadRequest
	}
	err = s.r.UpdateStatusBlog(ctx, database.UpdateStatusBlogParams{
		BlStatus: sql.NullInt32{Int32: updateStatusBlog.BlStatus, Valid: true},
		Updatedby: sql.NullString{String: User.UsID, Valid: true},
		BlID: updateStatusBlog.BlID,
	})
	if err != nil {
		return fmt.Errorf("failed to update status blog: %w", err), http.StatusBadRequest
	}
	return nil, http.StatusOK
}