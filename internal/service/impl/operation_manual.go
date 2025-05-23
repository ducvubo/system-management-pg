package impl

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"system-management-pg/internal/database"
	"system-management-pg/internal/model"
	"system-management-pg/pkg/response"

	"github.com/google/uuid"
)

type sOperationManual struct {
	r *database.Queries
}

func NewOperationManualImpl(r *database.Queries) *sOperationManual {
	return &sOperationManual{
		r: r,
	}
}

func (s *sOperationManual) CreateOperationManual(ctx context.Context, createOperationManual *model.CreateOperationManualDto,Account *model.Account) (err error, statusCode int) {
	newOperationManual := database.CreateOperationManualParams{
		OperaManualID      : uuid.New().String(),
		OperaManualTitle: sql.NullString{String: createOperationManual.OperaManualTitle, Valid: true},
		OperaManualContent: createOperationManual.OperaManualContent,
		OperaManualType: createOperationManual.OperaManualType,
		OperaManualNote: sql.NullString{String: createOperationManual.OperaManualNote, Valid: true},
		Createdby: sql.NullString{String: Account.ID, Valid: true},
		OperaManuaResID: Account.RestaurantID,
	}
	_, err = s.r.CreateOperationManual(ctx, newOperationManual)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	return nil, http.StatusCreated
}

func (s *sOperationManual) FindOperationManual(ctx context.Context, OperaManualID string, Account *model.Account) (out *model.OperationManualOutput, err error, statusCode int) {
	internalNote, err := s.r.GetOperationManual(ctx,database.GetOperationManualParams{
		OperaManualID: OperaManualID,
		 OperaManuaResID: Account.RestaurantID,
	})
	if err != nil {
		return nil, err, http.StatusBadRequest
	}
	return &model.OperationManualOutput{
		OperaManualID: internalNote.OperaManualID,
		OperaManualTitle: internalNote.OperaManualTitle.String,
		OperaManualContent: internalNote.OperaManualContent,
		OperaManualType: internalNote.OperaManualType,
		OperaManualNote: internalNote.OperaManualNote.String,
		Isdeleted: internalNote.Isdeleted.Int32,
	}, nil, http.StatusOK
}

func (s *sOperationManual) UpdateOperationManual(ctx context.Context, updateOperationManual *model.UpdateOperationManualDto, Account *model.Account) (err error, statusCode int) {
	_, err = s.r.GetOperationManual(ctx,database.GetOperationManualParams{
		OperaManualID: 		updateOperationManual.OperaManualID,
		OperaManuaResID: Account.RestaurantID,
	})
	if err != nil {
		return fmt.Errorf("không tìm thấy đề xuất nội bộ"), http.StatusBadRequest
	}


	err = s.r.UpdateOperationManual(ctx, database.UpdateOperationManualParams{
		OperaManualTitle: sql.NullString{String: updateOperationManual.OperaManualTitle, Valid: true},
		OperaManualContent: updateOperationManual.OperaManualContent,
		OperaManualType: updateOperationManual.OperaManualType,
		OperaManualNote: sql.NullString{String: updateOperationManual.OperaManualNote, Valid: true},
		OperaManualID: updateOperationManual.OperaManualID,
		Updatedby: sql.NullString{String: Account.ID, Valid: true},
		OperaManuaResID:  Account.RestaurantID,
	})
	if err != nil {
		return err, http.StatusInternalServerError
	}
	return nil, http.StatusCreated
}

func (s *sOperationManual) DeleteOperationManual(ctx context.Context, OperaManualID string, Account *model.Account) (err error, statusCode int) {
	_, err = s.r.GetOperationManual(ctx,database.GetOperationManualParams{
		OperaManualID: OperaManualID,
		OperaManuaResID: Account.RestaurantID,
	})
	if err != nil {
		return fmt.Errorf("không tìm thấy đề xuất nội bộ"), http.StatusBadRequest
	}
	err = s.r.DeleteOperationManual(ctx, database.DeleteOperationManualParams{
		OperaManualID: OperaManualID,
		Deletedby: sql.NullString{String: Account.ID, Valid: true},
		OperaManuaResID:  Account.RestaurantID,
	})
	if err != nil {
		return err, http.StatusInternalServerError
	}
	return nil, http.StatusCreated
}

func (s *sOperationManual) RestoreOperationManual(ctx context.Context, OperaManualID string, Account *model.Account) (err error, statusCode int) {
	_, err = s.r.GetOperationManual(ctx,database.GetOperationManualParams{
		OperaManualID: OperaManualID,
		OperaManuaResID: Account.RestaurantID,
	})
	if err != nil {
		return fmt.Errorf("không tìm thấy đề xuất nội bộ"), http.StatusBadRequest
	}
	
	err = s.r.RestoreOperationManual(ctx, database.RestoreOperationManualParams{
		OperaManualID: OperaManualID,
		OperaManuaResID: Account.RestaurantID,
	})
	if err != nil {
		return err, http.StatusInternalServerError
	}
	return nil, http.StatusCreated
}

func (s *sOperationManual) GetAllOperationManual(ctx context.Context, Limit int32, Offset int32, isDeleted int32, OperaManualTitle string, Account *model.Account) (out response.ModelPagination[[]*model.OperationManualOutput], err error, statusCode int) {
	internalNotes, err := s.r.GetListOperationManual(ctx, database.GetListOperationManualParams{
		Isdeleted:   sql.NullInt32{Int32: isDeleted, Valid: true},
		Isdeleted_2: sql.NullInt32{Int32: isDeleted, Valid: true},
		Limit:       Limit,
		Column4: float64(Limit),
		Offset:      Offset,
		OperaManuaResID: Account.RestaurantID,
		OperaManuaResID_2: Account.RestaurantID,
		OperaManualTitle:     sql.NullString{String: "%" + OperaManualTitle + "%", Valid: true},
		OperaManualTitle_2:     sql.NullString{String: "%" + OperaManualTitle + "%", Valid: true},
	})
	if err != nil {
		return response.ModelPagination[[]*model.OperationManualOutput]{}, err, http.StatusInternalServerError
	}
	var internalOutputs []*model.OperationManualOutput
	for _, user := range internalNotes {
		internalOutputs = append(internalOutputs, &model.OperationManualOutput{
			OperaManualID:        user.OperaManualID,
			OperaManualTitle:     user.OperaManualTitle.String,
			OperaManualType:      user.OperaManualType,
			OperaManualStatus:    string(user.OperaManualStatus.OperationManualOperaManualStatus),
			OperaManualNote:      user.OperaManualNote.String,
			OperaManualContent:   user.OperaManualContent,	
		})
	}
	var totalPages, totalItems int32 = 0, 0
	if len(internalNotes) > 0 {
    	totalPages = int32(internalNotes[0].TotalPages.(float64))  
    	totalItems = int32(internalNotes[0].TotalItems)
	}

	return response.ModelPagination[[]*model.OperationManualOutput]{
		Result: internalOutputs,
		MetaPagination: response.MetaPagination{
			Current:    Offset,
			PageSize:   Limit,
			TotalPage:  totalPages,
			TotalItems: int64(totalItems),
		},
	}, nil, http.StatusOK

}

func(s *sOperationManual) UpdateOperationManualStatus(ctx context.Context, updateOperationManualStatus *model.UpdateOperationManualStatusDto, Account *model.Account) (err error, statusCode int) {
	_, err = s.r.GetOperationManual(ctx,database.GetOperationManualParams{
		OperaManualID: updateOperationManualStatus.OperaManualID,
		OperaManuaResID: Account.RestaurantID,
	})
	if err != nil {
		return fmt.Errorf("không tìm thấy đề xuất nội bộ"), http.StatusBadRequest
	}
	status := database.OperationManualOperaManualStatus(updateOperationManualStatus.OperaManualStatus)

	nullStatus := database.NullOperationManualOperaManualStatus{
		OperationManualOperaManualStatus: status,
		Valid:                             true,
	}

	err = s.r.UpdateOperationManualStatus(ctx, database.UpdateOperationManualStatusParams{
		OperaManualID:     updateOperationManualStatus.OperaManualID,
		OperaManualStatus: nullStatus,
		OperaManuaResID:   Account.RestaurantID,
		Updatedby:         sql.NullString{String: Account.ID, Valid: true}, 
	})

	if err != nil {
		return err, http.StatusInternalServerError
	}
	return nil, http.StatusCreated
}