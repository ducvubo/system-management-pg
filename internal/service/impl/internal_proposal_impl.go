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

type sInternalProposal struct {
	r *database.Queries
}

func NewInternalProposalImpl(r *database.Queries) *sInternalProposal {
	return &sInternalProposal{
		r: r,
	}
}

func (s *sInternalProposal) CreateInternalProposal(ctx context.Context, createInternalProposal *model.CreateInternalProposalDto,Account *model.Account) (err error, statusCode int) {
	newInternalProposal := database.CreateInternalProposalParams{
		ItnProposalID: uuid.New().String(),
		ItnProposalTitle: sql.NullString{String: createInternalProposal.ItnProposalTitle,Valid:  true},
		ItnProposalContent: sql.NullString{String: createInternalProposal.ItnProposalContent,Valid:  true},
		ItnProposalType: sql.NullString{String: createInternalProposal.ItnProposalType,Valid:  true},
		Createdby: sql.NullString{String: Account.ID, Valid: true},
		ItnProposalResID: Account.RestaurantID,
	}
	_, err = s.r.CreateInternalProposal(ctx, newInternalProposal)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	kafka.SendMessageToKafka(ctx, "NOTIFICATION_ACCOUNT_CREATE", kafka.NotificationPayload{
		RestaurantID: Account.RestaurantID,
		NotiContent:  fmt.Sprintf("Đề xuất nội bộ %s vừa được tạo", createInternalProposal.ItnProposalTitle),
		NotiTitle:    "Đề xuất nội bộ",
		NotiType:     "internal_proposal",
		NotiMetadata: `{"text":"new internal proposal"}`,
		SendObject:   "all_account",
	})
	return nil, http.StatusCreated
}

func (s *sInternalProposal) FindInternalProposal(ctx context.Context, ItnProposalId string, Account *model.Account) (out *model.InternalProposalOutput, err error, statusCode int) {
	internalNote, err := s.r.GetInternalProposal(ctx,database.GetInternalProposalParams{
		ItnProposalID: ItnProposalId,
		ItnProposalResID: Account.RestaurantID,
	})
	if err != nil {
		return nil, err, http.StatusBadRequest
	}
	return &model.InternalProposalOutput{
		ItnProposalId: internalNote.ItnProposalID,
		ItnProposalTitle: internalNote.ItnProposalTitle.String,
		ItnProposalContent: internalNote.ItnProposalContent.String,
		ItnProposalType: internalNote.ItnProposalType.String,
		Isdeleted: internalNote.Isdeleted.Int32,
	}, nil, http.StatusOK
}

func (s *sInternalProposal) UpdateInternalProposal(ctx context.Context, updateInternalProposal *model.UpdateInternalProposalDto, Account *model.Account) (err error, statusCode int) {
	_, err = s.r.GetInternalProposal(ctx,database.GetInternalProposalParams{
		ItnProposalID: updateInternalProposal.ItnProposalId,
		ItnProposalResID: Account.RestaurantID,
	})
	if err != nil {
		return fmt.Errorf("không tìm thấy đề xuất nội bộ"), http.StatusBadRequest
	}


	err = s.r.UpdateInternalProposal(ctx, database.UpdateInternalProposalParams{
		ItnProposalTitle: sql.NullString{String: updateInternalProposal.ItnProposalTitle, Valid: true},
		ItnProposalContent: sql.NullString{String: updateInternalProposal.ItnProposalContent, Valid: true},
		ItnProposalType: sql.NullString{String: updateInternalProposal.ItnProposalType, Valid: true},
		ItnProposalID: updateInternalProposal.ItnProposalId,
		Updatedby: sql.NullString{String: Account.ID, Valid: true},
		ItnProposalResID: Account.RestaurantID,
	})
	if err != nil {
		return err, http.StatusInternalServerError
	}
	kafka.SendMessageToKafka(ctx, "NOTIFICATION_ACCOUNT_CREATE", kafka.NotificationPayload{
		RestaurantID: Account.RestaurantID,
		NotiContent:  fmt.Sprintf("Đề xuất nội bộ %s vừa được cập nhật", updateInternalProposal.ItnProposalTitle),
		NotiTitle:    "Đề xuất nội bộ",
		NotiType:     "internal_proposal",
		NotiMetadata: `{"text":"update internal proposal"}`,
		SendObject:   "all_account",
	})
	return nil, http.StatusCreated
}

func (s *sInternalProposal) DeleteInternalProposal(ctx context.Context, ItnProposalId string, Account *model.Account) (err error, statusCode int) {
	internalNote, err := s.r.GetInternalProposal(ctx,database.GetInternalProposalParams{
		ItnProposalID: ItnProposalId,
		ItnProposalResID: Account.RestaurantID,
	})
	if err != nil {
		return fmt.Errorf("không tìm thấy đề xuất nội bộ"), http.StatusBadRequest
	}
	err = s.r.DeleteInternalProposal(ctx, database.DeleteInternalProposalParams{
		ItnProposalID: ItnProposalId,
		Deletedby: sql.NullString{String: Account.ID, Valid: true},
		ItnProposalResID: Account.RestaurantID,
	})
	if err != nil {
		return err, http.StatusInternalServerError
	}
	kafka.SendMessageToKafka(ctx, "NOTIFICATION_ACCOUNT_CREATE", kafka.NotificationPayload{
		RestaurantID: Account.RestaurantID,
		NotiContent:  fmt.Sprintf("Đề xuất nội bộ %s vừa được xóa", internalNote.ItnProposalTitle.String),
		NotiTitle:    "Đề xuất nội bộ",
		NotiType:     "internal_proposal",
		NotiMetadata: `{"text":"delete internal proposal"}`,
		SendObject:   "all_account",
	})
	return nil, http.StatusCreated
}

func (s *sInternalProposal) RestoreInternalProposal(ctx context.Context, ItnProposalId string, Account *model.Account) (err error, statusCode int) {
	internalNote, err := s.r.GetInternalProposal(ctx,database.GetInternalProposalParams{
		ItnProposalID: ItnProposalId,
		ItnProposalResID: Account.RestaurantID,
	})
	if err != nil {
		return fmt.Errorf("không tìm thấy đề xuất nội bộ"), http.StatusBadRequest
	}
	
	err = s.r.RestoreInternalProposal(ctx, database.RestoreInternalProposalParams{
		ItnProposalID: ItnProposalId,
		ItnProposalResID: Account.RestaurantID,
	})
	if err != nil {
		return err, http.StatusInternalServerError
	}
	kafka.SendMessageToKafka(ctx, "NOTIFICATION_ACCOUNT_CREATE", kafka.NotificationPayload{
		RestaurantID: Account.RestaurantID,
		NotiContent:  fmt.Sprintf("Đề xuất nội bộ %s vừa được khôi phục", internalNote.ItnProposalTitle.String),
		NotiTitle:    "Đề xuất nội bộ",
		NotiType:     "internal_proposal",
		NotiMetadata: `{"text":"restore internal proposal"}`,
		SendObject:   "all_account",
	})
		return nil, http.StatusCreated
}

func (s *sInternalProposal) GetAllInternalProposal(ctx context.Context, Limit int32, Offset int32, isDeleted int32, ItnProposalTitle string, Account *model.Account) (out response.ModelPagination[[]*model.InternalProposalOutput], err error, statusCode int) {
	internalNotes, err := s.r.GetListInternalProposal(ctx, database.GetListInternalProposalParams{
		Isdeleted:   sql.NullInt32{Int32: isDeleted, Valid: true},
		Isdeleted_2: sql.NullInt32{Int32: isDeleted, Valid: true},
		Limit:       Limit,
		Column4: float64(Limit),
		Offset:      Offset,
		ItnProposalResID: Account.RestaurantID,
		ItnProposalResID_2: Account.RestaurantID,
		ItnProposalTitle:      sql.NullString{String: "%" + ItnProposalTitle + "%", Valid: true},
		ItnProposalTitle_2:     sql.NullString{String: "%" + ItnProposalTitle + "%", Valid: true},
	})
	if err != nil {
		return response.ModelPagination[[]*model.InternalProposalOutput]{}, err, http.StatusInternalServerError
	}
	var internalOutputs []*model.InternalProposalOutput
	for _, user := range internalNotes {
		internalOutputs = append(internalOutputs, &model.InternalProposalOutput{
			ItnProposalId:        user.ItnProposalID,
			ItnProposalTitle:     user.ItnProposalTitle.String,
			ItnProposalContent:   user.ItnProposalContent.String,
			ItnProposalType:      user.ItnProposalType.String,
			ItnProposalStatus:   string(user.ItnProposalStatus.InternalProposalItnProposalStatus),
		})
	}
	// Nếu không có user nào, đặt TotalPage và TotalItems về 0
	var totalPages, totalItems int32 = 0, 0
	if len(internalNotes) > 0 {
    	totalPages = int32(internalNotes[0].TotalPages.(float64))  // Convert từ interface{} -> float64 -> int32
    	totalItems = int32(internalNotes[0].TotalItems)
	}

	return response.ModelPagination[[]*model.InternalProposalOutput]{
		Result: internalOutputs,
		MetaPagination: response.MetaPagination{
			Current:    Offset,
			PageSize:   Limit,
			TotalPage:  totalPages,
			TotalItems: int64(totalItems),
		},
	}, nil, http.StatusOK

}

func(s *sInternalProposal) UpdateInternalProposalStatus(ctx context.Context, updateInternalProposalStatus *model.UpdateInternalProposalStatusDto, Account *model.Account) (err error, statusCode int) {
	internalNote, err := s.r.GetInternalProposal(ctx,database.GetInternalProposalParams{
		ItnProposalID: updateInternalProposalStatus.ItnProposalId,
		ItnProposalResID: Account.RestaurantID,
	})
	if err != nil {
		return fmt.Errorf("không tìm thấy đề xuất nội bộ"), http.StatusBadRequest
	}
	status := database.InternalProposalItnProposalStatus(updateInternalProposalStatus.ItnProposalStatus)

	nullStatus := database.NullInternalProposalItnProposalStatus{
		InternalProposalItnProposalStatus: status,
		Valid:                             true,
	}

	err = s.r.UpdateInternalProposalStatus(ctx, database.UpdateInternalProposalStatusParams{
		ItnProposalID:     updateInternalProposalStatus.ItnProposalId,
		ItnProposalStatus: nullStatus,
		ItnProposalResID:  Account.RestaurantID,
		Updatedby:         sql.NullString{String: Account.ID, Valid: true}, // nếu có
	})

	if err != nil {
		return err, http.StatusInternalServerError
	}
	kafka.SendMessageToKafka(ctx, "NOTIFICATION_ACCOUNT_CREATE", kafka.NotificationPayload{
		RestaurantID: Account.RestaurantID,
		NotiContent:  fmt.Sprintf("Đề xuất nội bộ %s vừa được cập nhật trạng thái", internalNote.ItnProposalTitle.String),
		NotiTitle:    "Đề xuất nội bộ",
		NotiType:     "internal_proposal",
		NotiMetadata: `{"text":"update internal proposal status"}`,
		SendObject:   "all_account",
	})
	return nil, http.StatusCreated
}