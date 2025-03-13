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

	for _, note := range createBlog.BlogNote {
		_, err = s.r.CreateBlogNote(ctx, database.CreateBlogNoteParams{
			BlID: BlID,
			BlNoteID: uuid.New().String(),
			BlContent: note.BlContent,
		})
		if err != nil {
			return fmt.Errorf("failed to create blog note: %w", err), http.StatusBadRequest
		}
	}

	for _, related := range createBlog.BlogRelated {
		_, err = s.r.CreateRelatedBlog(ctx, database.CreateRelatedBlogParams{
			BlID: BlID,
			BlRltID: related.BlRltID,
		})
		if err != nil {
			return fmt.Errorf("failed to create related blog: %w", err), http.StatusBadRequest
		}
	}

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

	err = s.r.DeleteBlogNote(ctx, updateBlog.BlID)

	err = s.r.DeleteRelatedBlog(ctx, updateBlog.BlID)

	for _, note := range updateBlog.BlogNote {
		_, err = s.r.CreateBlogNote(ctx, database.CreateBlogNoteParams{
			BlID: updateBlog.BlID,
			BlNoteID: uuid.New().String(),
			BlContent: note.BlContent,
		})
		if err != nil {
			return fmt.Errorf("failed to create blog note: %w", err), http.StatusBadRequest
		}
	}

	for _, related := range updateBlog.BlogRelated {
		_, err = s.r.CreateRelatedBlog(ctx, database.CreateRelatedBlogParams{
			BlID: updateBlog.BlID,
			BlRltID: related.BlRltID,
		})
		if err != nil {
			return fmt.Errorf("failed to create related blog: %w", err), http.StatusBadRequest
		}
	}

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

func (s *sBlog) FindBlogById(ctx context.Context, blId string) (out *model.BlogOutput, err error, statusCode int) {
	blog, err := s.r.GetBlogByID(ctx, blId)
	if err != nil {
		return nil, fmt.Errorf("failed to get blog: %w", err), http.StatusBadRequest
	}
	return &model.BlogOutput{
		BlID: blog.BlID,
		CatBlID: blog.CatBlID,
		BlTitle: blog.BlTitle,
		BlDescription: blog.BlDescription.String,
		BlSlug: blog.BlSlug,
		BlStatus: blog.BlStatus.Int32,
		BlImage: blog.BlImage.String,
		BlContent: blog.BlContent,
		BlType: blog.BlType.Int32,
		BlView: blog.BlView.Int32,
		BlPublishedTime: blog.BlPublishedTime.Time.String(),
		BlPublishedSchedule: blog.BlPublishedSchedule.Time.String(),
	}, nil, http.StatusOK
}

func (s *sBlog) GetAllBlog(ctx context.Context, Limit int32, Offset int32, isDeleted int32, BlTitle string) (out response.ModelPagination[[]*model.BlogOutput], err error, statusCode int) {
	blogs, err := s.r.GetListBlog(ctx, database.GetListBlogParams{
		Isdeleted: sql.NullInt32{Int32: int32(isDeleted), Valid: true},
		Limit: Limit,
		Offset: Offset,
		BlTitle: "%" + BlTitle + "%",
		Isdeleted_2: sql.NullInt32{Int32: int32(isDeleted), Valid: false},
		BlTitle_2: "%" + BlTitle + "%",
		Column3: float64(Limit),
	})
	if err != nil {
		return response.ModelPagination[[]*model.BlogOutput]{}, fmt.Errorf("failed to get list blog: %w", err), http.StatusBadRequest
	}
	var outData []*model.BlogOutput
	for _, blog := range blogs {
		outData = append(outData, &model.BlogOutput{
			BlID: blog.BlID,
			CatBlID: blog.CatBlID,
			BlTitle: blog.BlTitle,
			BlDescription: blog.BlDescription.String,
			BlSlug: blog.BlSlug,
			BlStatus: blog.BlStatus.Int32,
			BlImage: blog.BlImage.String,
			BlContent: blog.BlContent,
			BlType: blog.BlType.Int32,
			BlView: blog.BlView.Int32,
			BlPublishedTime: blog.BlPublishedTime.Time.String(),
			BlPublishedSchedule: blog.BlPublishedSchedule.Time.String(),
		})
	}
	return response.ModelPagination[[]*model.BlogOutput]{
		Result: outData,
		MetaPagination: response.MetaPagination{
			TotalItems: blogs[0].TotalItems,
			TotalPage: blogs[0].TotalPages,
			Current: Offset,
			PageSize: Limit,
		},
	}, nil, http.StatusOK
}

func (s *sBlog) DeleteBlog(ctx context.Context, blId string, User *model.UserManagementProfileOutput) (err error, statusCode int) {
	_, err = s.r.GetBlogByID(ctx, blId)
	if err != nil {
		return fmt.Errorf("failed to get blog: %w", err), http.StatusBadRequest
	}
	err = s.r.DeleteBlog(ctx, database.DeleteBlogParams{
		Deletedby: sql.NullString{String: User.UsID, Valid: true},
		BlID: blId,
	})
	if err != nil {
		return fmt.Errorf("failed to delete blog: %w", err), http.StatusBadRequest
	}
	return nil, http.StatusOK
}

func (s *sBlog) RestoreBlog(ctx context.Context, blId string, User *model.UserManagementProfileOutput) (err error, statusCode int) {
	_, err = s.r.GetBlogByID(ctx, blId)
	if err != nil {
		return fmt.Errorf("failed to get blog: %w", err), http.StatusBadRequest
	}
	err = s.r.RestoreBlog(ctx,blId)
	if err != nil {
		return fmt.Errorf("failed to restore blog: %w", err), http.StatusBadRequest
	}
	return nil, http.StatusOK
}

