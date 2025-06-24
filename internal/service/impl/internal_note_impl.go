package impl

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"system-management-pg/internal/database"
	"system-management-pg/internal/model"
	"system-management-pg/internal/utils/kafka"
	"system-management-pg/pkg/response"

	"github.com/google/uuid"
)

type sInternalNote struct {
	r *database.Queries
}

func NewInternalNoteImpl(r *database.Queries) *sInternalNote {
	return &sInternalNote{
		r: r,
	}
}

func (s *sInternalNote) CreateInternalNote(ctx context.Context, createInternalNote *model.CreateInternalNoteDto,Account *model.Account) (err error, statusCode int) {
	newInternalNote := database.CreateInternalNoteParams{
		ItnNoteID: uuid.New().String(),
		ItnNoteTitle: sql.NullString{String: createInternalNote.ItnNoteTitle,Valid:  true},
		ItnNoteContent: sql.NullString{String: createInternalNote.ItnNoteContent,Valid:  true},
		ItnNoteType: sql.NullString{String: createInternalNote.ItnNoteType,Valid:  true},
		Createdby: sql.NullString{String: Account.ID, Valid: true},
		ItnNoteResID: Account.RestaurantID,
	}
	_, err = s.r.CreateInternalNote(ctx, newInternalNote)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	kafka.SendMessageToKafka(ctx, "NOTIFICATION_ACCOUNT_CREATE", kafka.NotificationPayload{
		RestaurantID: Account.RestaurantID,
		NotiContent:  fmt.Sprintf("Ghi chú nội bộ %s vừa được tạo", createInternalNote.ItnNoteTitle),
		NotiTitle:    "Ghi chú nội bộ",
		NotiType:     "internal_note",
		NotiMetadata: `{"text":"new internal note"}`,
		SendObject:   "all_account",
	})

	return nil, http.StatusCreated
}

func (s *sInternalNote) FindInternalNote(ctx context.Context, ItnNoteId string, Account *model.Account) (out *model.InternalNoteOutput, err error, statusCode int) {
	internalNote, err := s.r.GetInternalNote(ctx,database.GetInternalNoteParams{
		ItnNoteID: ItnNoteId,
		ItnNoteResID: Account.RestaurantID,
	})
	if err != nil {
		return nil, err, http.StatusBadRequest
	}
	return &model.InternalNoteOutput{
		ItnNoteId: internalNote.ItnNoteID,
		ItnNoteTitle: internalNote.ItnNoteTitle.String,
		ItnNoteContent: internalNote.ItnNoteContent.String,
		ItnNoteType: internalNote.ItnNoteType.String,
		Isdeleted: internalNote.Isdeleted.Int32,
	}, nil, http.StatusOK
}

func (s *sInternalNote) UpdateInternalNote(ctx context.Context, updateInternalNote *model.UpdateInternalNoteDto, Account *model.Account) (err error, statusCode int) {
	_, err = s.r.GetInternalNote(ctx,database.GetInternalNoteParams{
		ItnNoteID: updateInternalNote.ItnNoteId,
		ItnNoteResID: Account.RestaurantID,
	})
	if err != nil {
		return fmt.Errorf("không tìm thấy ghi chú nội bộ"), http.StatusBadRequest
	}


	err = s.r.UpdateInternalNote(ctx, database.UpdateInternalNoteParams{
		ItnNoteTitle: sql.NullString{String: updateInternalNote.ItnNoteTitle, Valid: true},
		ItnNoteContent: sql.NullString{String: updateInternalNote.ItnNoteContent, Valid: true},
		ItnNoteType: sql.NullString{String: updateInternalNote.ItnNoteType, Valid: true},
		ItnNoteID: updateInternalNote.ItnNoteId,
		Updatedby: sql.NullString{String: Account.ID, Valid: true},
		ItnNoteResID: Account.RestaurantID,
	})
	if err != nil {
		return err, http.StatusInternalServerError
	}
	kafka.SendMessageToKafka(ctx, "NOTIFICATION_ACCOUNT_CREATE", kafka.NotificationPayload{
		RestaurantID: Account.RestaurantID,
		NotiContent:  fmt.Sprintf("Ghi chú nội bộ %s vừa được cập nhật", updateInternalNote.ItnNoteTitle),
		NotiTitle:    "Ghi chú nội bộ",
		NotiType:     "internal_note",
		NotiMetadata: `{"text":"update internal note"}`,
		SendObject:   "all_account",
	})
	return nil, http.StatusCreated
}

func (s *sInternalNote) DeleteInternalNote(ctx context.Context, ItnNoteId string, Account *model.Account) (err error, statusCode int) {
	internalNote, err := s.r.GetInternalNote(ctx,database.GetInternalNoteParams{
		ItnNoteID: ItnNoteId,
		ItnNoteResID: Account.RestaurantID,
	})
	if err != nil {
		return fmt.Errorf("không tìm thấy ghi chú nội bộ"), http.StatusBadRequest
	}
	err = s.r.DeleteInternalNote(ctx, database.DeleteInternalNoteParams{
		ItnNoteID: ItnNoteId,
		Deletedby: sql.NullString{String: Account.ID, Valid: true},
		ItnNoteResID: Account.RestaurantID,
	})
	if err != nil {
		return err, http.StatusInternalServerError
	}
	kafka.SendMessageToKafka(ctx, "NOTIFICATION_ACCOUNT_CREATE", kafka.NotificationPayload{
		RestaurantID: Account.RestaurantID,
		NotiContent:  fmt.Sprintf("Ghi chú nội bộ %s vừa được xóa", internalNote.ItnNoteTitle.String),
		NotiTitle:    "Ghi chú nội bộ",
		NotiType:     "internal_note",
		NotiMetadata: `{"text":"delete internal note"}`,
		SendObject:   "all_account",
	})
	return nil, http.StatusCreated
}

func (s *sInternalNote) RestoreInternalNote(ctx context.Context, ItnNoteId string, Account *model.Account) (err error, statusCode int) {
	internalNote, err := s.r.GetInternalNote(ctx,database.GetInternalNoteParams{
		ItnNoteID: ItnNoteId,
		ItnNoteResID: Account.RestaurantID,
	})
	if err != nil {
		return fmt.Errorf("không tìm thấy ghi chú nội bộ"), http.StatusBadRequest
	}
	
	err = s.r.RestoreInternalNote(ctx, database.RestoreInternalNoteParams{
		ItnNoteID: ItnNoteId,
		ItnNoteResID: Account.RestaurantID,
	})
	if err != nil {
		return err, http.StatusInternalServerError
	}
	kafka.SendMessageToKafka(ctx, "NOTIFICATION_ACCOUNT_CREATE", kafka.NotificationPayload{
		RestaurantID: Account.RestaurantID,
		NotiContent:  fmt.Sprintf("Ghi chú nội bộ %s vừa được khôi phục", internalNote.ItnNoteTitle.String),
		NotiTitle:    "Ghi chú nội bộ",
		NotiType:     "internal_note",
		NotiMetadata: `{"text":"restore internal note"}`,
		SendObject:   "all_account",
	})
		return nil, http.StatusCreated
}

func (s *sInternalNote) GetAllInternalNote(ctx context.Context, Limit int32, Offset int32, isDeleted int32, ItnNoteTitle string, Account *model.Account) (out response.ModelPagination[[]*model.InternalNoteOutput], err error, statusCode int) {
	internalNotes, err := s.r.GetListInternalNote(ctx, database.GetListInternalNoteParams{
		Isdeleted:   sql.NullInt32{Int32: isDeleted, Valid: true},
		Isdeleted_2: sql.NullInt32{Int32: isDeleted, Valid: true},
		Limit:       Limit,
		Offset:      Offset,
		Column4: float64(Limit),
		ItnNoteResID: Account.RestaurantID,
		ItnNoteResID_2: Account.RestaurantID,
		ItnNoteTitle:      sql.NullString{String: "%" + ItnNoteTitle + "%", Valid: true},
		ItnNoteTitle_2:     sql.NullString{String: "%" + ItnNoteTitle + "%", Valid: true},
	})
	if err != nil {
		return response.ModelPagination[[]*model.InternalNoteOutput]{}, err, http.StatusInternalServerError
	}
	var internalOutputs []*model.InternalNoteOutput
	for _, user := range internalNotes {
		internalOutputs = append(internalOutputs, &model.InternalNoteOutput{
			ItnNoteId:        user.ItnNoteID,
			ItnNoteTitle:     user.ItnNoteTitle.String,
			ItnNoteContent:   user.ItnNoteContent.String,
			ItnNoteType:      user.ItnNoteType.String,
		})
	}
	// Nếu không có user nào, đặt TotalPage và TotalItems về 0
	var totalPages, totalItems int32 = 0, 0
	if len(internalNotes) > 0 {
    	totalPages = int32(internalNotes[0].TotalPages.(float64))  // Convert từ interface{} -> float64 -> int32
    	totalItems = int32(internalNotes[0].TotalItems)
	}

	return response.ModelPagination[[]*model.InternalNoteOutput]{
		Result: internalOutputs,
		MetaPagination: response.MetaPagination{
			Current:    Offset,
			PageSize:   Limit,
			TotalPage:  totalPages,
			TotalItems: int64(totalItems),
		},
	}, nil, http.StatusOK

}
