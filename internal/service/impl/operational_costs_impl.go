package impl

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"system-management-pg/internal/database"
	"system-management-pg/internal/model"
	"system-management-pg/internal/utils/kafka"
	"system-management-pg/pkg/response"
	"time"

	"github.com/google/uuid"
)

type sOperationalCosts struct {
	r *database.Queries
}

func NewOperationalCostsImpl(r *database.Queries) *sOperationalCosts {
	return &sOperationalCosts{
		r: r,
	}
}

func (s *sOperationalCosts) CreateOperationalCosts(ctx context.Context, createOperationalCosts *model.CreateOperationalCostsDto,Account *model.Account) (err error, statusCode int) {
    operaCostDate, err := time.Parse("2006-01-02", createOperationalCosts.OperaCostDate)
	if err != nil {
		return fmt.Errorf("invalid date format for OperaCostDate"), http.StatusBadRequest
	}

	newOperationalCosts := database.CreateOperationalCostsParams{
		Createdby: sql.NullString{String: Account.ID, Valid: true},
		OperaCostResID: Account.RestaurantID,
		OperaCostID: uuid.New().String(),
		OperaCostType: createOperationalCosts.OperaCostType,
		OperaCostAmount:  strconv.FormatInt(*createOperationalCosts.OperaCostAmount, 10),
		OperaCostDate: operaCostDate,
		OperaCostDescription: sql.NullString{String: createOperationalCosts.OperaCostDescription, Valid: true},
	}
	_, err = s.r.CreateOperationalCosts(ctx, newOperationalCosts)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	kafka.SendMessageToKafka(ctx, "NOTIFICATION_ACCOUNT_CREATE", kafka.NotificationPayload{
		RestaurantID: Account.RestaurantID,
		NotiContent:  fmt.Sprintf("Chi phí vận hành %s vừa được tạo", createOperationalCosts.OperaCostType.String),
		NotiTitle:    "Chi phí vận hành",
		NotiType:     "operational_costs",
		NotiMetadata: `{"text":"new operational costs"}`,
		SendObject:   "all_account",
	})
	return nil, http.StatusCreated
}

func (s *sOperationalCosts) FindOperationalCosts(ctx context.Context, OperaCostID string, Account *model.Account) (out *model.OperationalCostsOutput, err error, statusCode int) {
	operationalCosts, err := s.r.GetOperationalCosts(ctx,database.GetOperationalCostsParams{
		OperaCostID: OperaCostID,
		OperaCostResID: Account.RestaurantID,
	})
	if err != nil {
		return nil, err, http.StatusBadRequest
	}
	output := &model.OperationalCostsOutput{
    	OperaCostID:               operationalCosts.OperaCostID,
    	OperaCostType: 		 string(operationalCosts.OperaCostType),
		OperaCostDescription: operationalCosts.OperaCostDescription.String,
		OperaCostDate: 	operationalCosts.OperaCostDate.Format("2006-01-02"),
	}

	if operationalCosts.OperaCostAmount != "" {
    cost, err := strconv.ParseInt(operationalCosts.OperaCostAmount, 10, 64)
    if err == nil {
        output.OperaCostAmount = &cost
    }
}

	return output, nil, http.StatusOK
}

func (s *sOperationalCosts) UpdateOperationalCosts(ctx context.Context, updateOperationalCosts *model.UpdateOperationalCostsDto, Account *model.Account) (err error, statusCode int) {
	_, err = s.r.GetOperationalCosts(ctx,database.GetOperationalCostsParams{
		OperaCostID: updateOperationalCosts.OperaCostID,
		OperaCostResID: Account.RestaurantID,
	})
	if err != nil {
		return fmt.Errorf("không tìm thấy đề xuất nội bộ"), http.StatusBadRequest
	}
	operaCostDate, err := time.Parse("2006-01-02", updateOperationalCosts.OperaCostDate)
	if err != nil {
		return fmt.Errorf("invalid date format for OperaCostDate"), http.StatusBadRequest
	}

	err = s.r.UpdateOperationalCosts(ctx, database.UpdateOperationalCostsParams{
		OperaCostID: updateOperationalCosts.OperaCostID,
		OperaCostDate: operaCostDate,
		OperaCostResID: Account.RestaurantID,
		OperaCostType: updateOperationalCosts.OperaCostType,
		OperaCostAmount: strconv.FormatInt(*updateOperationalCosts.OperaCostAmount, 10),
		OperaCostDescription: sql.NullString{String: updateOperationalCosts.OperaCostDescription, Valid: true},
		Updatedby: sql.NullString{String: Account.ID, Valid: true},
	})
	if err != nil {
		return err, http.StatusInternalServerError
	}
	kafka.SendMessageToKafka(ctx, "NOTIFICATION_ACCOUNT_CREATE", kafka.NotificationPayload{
		RestaurantID: Account.RestaurantID,
		NotiContent:  fmt.Sprintf("Chi phí vận hành %s vừa được cập nhật", updateOperationalCosts.OperaCostType.String),
		NotiTitle:    "Chi phí vận hành",
		NotiType:     "operational_costs",
		NotiMetadata: `{"text":"update operational costs"}`,
		SendObject:   "all_account",
	})
	return nil, http.StatusCreated
}

func (s *sOperationalCosts) DeleteOperationalCosts(ctx context.Context, OperaCostID string, Account *model.Account) (err error, statusCode int) {
	operationalCost, err := s.r.GetOperationalCosts(ctx,database.GetOperationalCostsParams{
		OperaCostID: OperaCostID,
		OperaCostResID: Account.RestaurantID,
	})
	if err != nil {
		return fmt.Errorf("không tìm thấy đề xuất nội bộ"), http.StatusBadRequest
	}
	err = s.r.DeleteOperationalCosts(ctx, database.DeleteOperationalCostsParams{
		OperaCostID: OperaCostID,
		Deletedby: sql.NullString{String: Account.ID, Valid: true},
		OperaCostResID: Account.RestaurantID,
	})
	if err != nil {
		return err, http.StatusInternalServerError
	}
	kafka.SendMessageToKafka(ctx, "NOTIFICATION_ACCOUNT_CREATE", kafka.NotificationPayload{
		RestaurantID: Account.RestaurantID,
		NotiContent:  fmt.Sprintf("Chi phí vận hành %s vừa được xóa", operationalCost.OperaCostType.String),
		NotiTitle:    "Chi phí vận hành",
		NotiType:     "operational_costs",
		NotiMetadata: `{"text":"delete operational costs"}`,
		SendObject:   "all_account",
	})
	return nil, http.StatusCreated
}

func (s *sOperationalCosts) RestoreOperationalCosts(ctx context.Context, OperaCostID string, Account *model.Account) (err error, statusCode int) {
	operationalCostsParams, err := s.r.GetOperationalCosts(ctx,database.GetOperationalCostsParams{
		OperaCostID: OperaCostID,
		OperaCostResID: Account.RestaurantID,
	})
	if err != nil {
		return fmt.Errorf("không tìm thấy đề xuất nội bộ"), http.StatusBadRequest
	}
	
	err = s.r.RestoreOperationalCosts(ctx, database.RestoreOperationalCostsParams{
		OperaCostID: OperaCostID,
		OperaCostResID: Account.RestaurantID,
	})
	if err != nil {
		return err, http.StatusInternalServerError
	}
	kafka.SendMessageToKafka(ctx, "NOTIFICATION_ACCOUNT_CREATE", kafka.NotificationPayload{
		RestaurantID: Account.RestaurantID,
		NotiContent:  fmt.Sprintf("Chi phí vận hành %s vừa được khôi phục", operationalCostsParams.OperaCostType.String),
		NotiTitle:    "Chi phí vận hành",
		NotiType:     "operational_costs",
		NotiMetadata: `{"text":"restore operational costs"}`,
		SendObject:   "all_account",
	})
	return nil, http.StatusCreated
}

func (s *sOperationalCosts) GetAllOperationalCosts(ctx context.Context, Limit int32, Offset int32, isDeleted int32, OperaCostType string, Account *model.Account) (out response.ModelPagination[[]*model.OperationalCostsOutput], err error, statusCode int) {
	operationalCostss, err := s.r.GetListOperationalCostss(ctx, database.GetListOperationalCostssParams{
		Isdeleted:   sql.NullInt32{Int32: isDeleted, Valid: true},
		Isdeleted_2: sql.NullInt32{Int32: isDeleted, Valid: true},
		Limit:       Limit,
		Offset:      Offset,
		Column4: float64(Limit),
		OperaCostResID: Account.RestaurantID,
		OperaCostResID_2: Account.RestaurantID,
		OperaCostType:      "%" + OperaCostType + "%",
		OperaCostType_2:    "%" + OperaCostType + "%",
	})
	if err != nil {
		return response.ModelPagination[[]*model.OperationalCostsOutput]{}, err, http.StatusInternalServerError
	}
	var internalOutputs []*model.OperationalCostsOutput
	for _, user := range operationalCostss {
		var amount *int64
	if user.OperaCostAmount != "" {
    cost, err := strconv.ParseInt(user.OperaCostAmount, 10, 64)
    if err == nil {
        amount = &cost
    }
}
		internalOutputs = append(internalOutputs, &model.OperationalCostsOutput{
			OperaCostID:               user.OperaCostID,
			OperaCostType: 		   string(user.OperaCostType),
			OperaCostDate: 	user.OperaCostDate.Format("2006-01-02"),
			OperaCostStatus: string(user.OperaCostStatus.OperationalCostsOperaCostStatus),
			OperaCostDescription: user.OperaCostDescription.String,
			OperaCostAmount:amount ,
		})
	}
	var totalPages, totalItems int32 = 0, 0
	if len(operationalCostss) > 0 {
    	totalPages = int32(operationalCostss[0].TotalPages.(float64))  
    	totalItems = int32(operationalCostss[0].TotalItems)
	}

	return response.ModelPagination[[]*model.OperationalCostsOutput]{
		Result: internalOutputs,
		MetaPagination: response.MetaPagination{
			Current:    Offset,
			PageSize:   Limit,
			TotalPage:  totalPages,
			TotalItems: int64(totalItems),
		},
	}, nil, http.StatusOK

}

func(s *sOperationalCosts) UpdateOperationalCostsStatus(ctx context.Context, updateOperationalCostsStatus *model.UpdateOperationalCostsStatusDto, Account *model.Account) (err error, statusCode int) {
	operationalCostsStatus, err := s.r.GetOperationalCosts(ctx,database.GetOperationalCostsParams{
		OperaCostID: updateOperationalCostsStatus.OperaCostID,
		OperaCostResID: Account.RestaurantID,
	})
	if err != nil {
		return fmt.Errorf("không tìm thấy đề xuất nội bộ"), http.StatusBadRequest
	}
	status := database.OperationalCostsOperaCostStatus(updateOperationalCostsStatus.OperaCostStatus)

	nullStatus := database.NullOperationalCostsOperaCostStatus{
		OperationalCostsOperaCostStatus: status,
		Valid:                             true,
	}

	err = s.r.UpdateOperationalCostsStatus(ctx, database.UpdateOperationalCostsStatusParams{
		OperaCostID:     updateOperationalCostsStatus.OperaCostID,
		OperaCostStatus:  nullStatus,
		OperaCostResID:  Account.RestaurantID,
		Updatedby:         sql.NullString{String: Account.ID, Valid: true}, // nếu có
	})

	if err != nil {
		return err, http.StatusInternalServerError
	}
	kafka.SendMessageToKafka(ctx, "NOTIFICATION_ACCOUNT_CREATE", kafka.NotificationPayload{
		RestaurantID: Account.RestaurantID,
		NotiContent:  fmt.Sprintf("Chi phí vận hành %s vừa được cập nhật trạng thái", operationalCostsStatus.OperaCostType.String),
		NotiTitle:    "Chi phí vận hành",
		NotiType:     "operational_costs",
		NotiMetadata: `{"text":"update operational costs status"}`,
		SendObject:   "all_account",
	})
	return nil, http.StatusCreated
}