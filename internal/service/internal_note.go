package service

import (
	"context"
	"system-management-pg/internal/model"
	"system-management-pg/pkg/response"
)

type (
	IInternalNote interface {
		CreateInternalNote(ctx context.Context, createInternalNote *model.CreateInternalNoteDto, Account *model.Account) (err error, statusCode int)
		FindInternalNote(ctx context.Context, ItnNoteId string, Account *model.Account) (out *model.InternalNoteOutput, err error, statusCode int)
		UpdateInternalNote(ctx context.Context, updateInternalNote *model.UpdateInternalNoteDto, Account *model.Account) (err error, statusCode int)
		DeleteInternalNote(ctx context.Context, ItnNoteId string, Account *model.Account) (err error, statusCode int)
		RestoreInternalNote(ctx context.Context, ItnNoteId string, Account *model.Account) (err error, statusCode int)
		GetAllInternalNote(ctx context.Context, Limit int32, Offset int32, isDeleted int32, ItnNoteTitle string, Account *model.Account) (out response.ModelPagination[[]*model.InternalNoteOutput], err error, statusCode int)
	}
)

var (
	localInternalNote IInternalNote
)

func InternalNote() IInternalNote {
	if localInternalNote == nil {
		panic("implement.............................................")
	}
	return localInternalNote
}

func InitInternalNote(i IInternalNote) {
	localInternalNote = i
}
