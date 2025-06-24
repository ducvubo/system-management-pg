package impl

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"system-management-pg/internal/database"
	"system-management-pg/internal/model"
	"system-management-pg/pkg/response"
	"time"
	"system-management-pg/internal/utils/kafka"
	//json
	"encoding/json"


	"github.com/google/uuid"
)

type sEquipmentMaintenance struct {
	r *database.Queries
}

func NewEquipmentMaintenanceImpl(r *database.Queries) *sEquipmentMaintenance {
	return &sEquipmentMaintenance{
		r: r,
	}
}

func (s *sEquipmentMaintenance) CreateEquipmentMaintenance(ctx context.Context, createEquipmentMaintenance *model.CreateEquipmentMaintenanceDto,Account *model.Account) (err error, statusCode int) {
	dateReported, err := time.Parse("2006-01-02", createEquipmentMaintenance.EqpMtnDateReported)
	if err != nil {
		return fmt.Errorf("invalid date format for EqpMtnDateReported"), http.StatusBadRequest
	}

	var dateFixed sql.NullTime
	if createEquipmentMaintenance.EqpMtnDateFixed != nil {
		parsedDate, err := time.Parse("2006-01-02", *createEquipmentMaintenance.EqpMtnDateFixed)
		if err != nil {
			return fmt.Errorf("invalid date format for EqpMtnDateFixed"), http.StatusBadRequest
		}
		dateFixed = sql.NullTime{Time: parsedDate, Valid: true}
	}
	newEquipmentMaintenance := database.CreateEquipmentMaintenanceParams{
		Createdby: sql.NullString{String: Account.ID, Valid: true},
		EqpMtnResID: Account.RestaurantID,
		EqpMtnID: uuid.New().String(),
		EqpMtnName: sql.NullString{String: createEquipmentMaintenance.EqpMtnName, Valid: true},
		EqpMtnLocation: sql.NullString{String: createEquipmentMaintenance.EqpMtnLocation, Valid: true},
		EqpMtnIssueDescription: sql.NullString{String: createEquipmentMaintenance.EqpMtnIssueDescription, Valid: true},
		EqpMtnReportedBy: sql.NullString{String: createEquipmentMaintenance.EqpMtnReportedBy, Valid: true},
		EqpMtnCost: sql.NullString{
			String: strconv.FormatInt(*createEquipmentMaintenance.EqpMtnCost, 10),
			Valid: createEquipmentMaintenance.EqpMtnCost != nil,
		},
		EqpMtnPerformedBy: sql.NullString{String: createEquipmentMaintenance.EqpMtnPerformedBy, Valid: true},
		EqpMtnDateReported: dateReported,
		EqpMtnDateFixed: dateFixed,
		EqpMtnNote: sql.NullString{String: createEquipmentMaintenance.EqpMtnNote, Valid: true},
	}
	_, err = s.r.CreateEquipmentMaintenance(ctx, newEquipmentMaintenance)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	kafka.SendMessageToKafka(ctx, "NOTIFICATION_ACCOUNT_CREATE", kafka.NotificationPayload{
		RestaurantID: Account.RestaurantID,
		NotiContent:  fmt.Sprintf("Thiết bị %s vừa được tạo yêu cầu bảo trì", createEquipmentMaintenance.EqpMtnName),
		NotiTitle:    "Bảo trì thiết bị",
		NotiType:     "maintenance",
		NotiMetadata: `{"text":"new maintenance request"}`,
		SendObject:   "all_account",
	})

	return nil, http.StatusCreated
}

func (s *sEquipmentMaintenance) FindEquipmentMaintenance(ctx context.Context, EqpMtnID string, Account *model.Account) (out *model.EquipmentMaintenanceOutput, err error, statusCode int) {
	equipmentMaintenance, err := s.r.GetEquipmentMaintenance(ctx,database.GetEquipmentMaintenanceParams{
		EqpMtnID: EqpMtnID,
		EqpMtnResID: Account.RestaurantID,
	})
	if err != nil {
		return nil, err, http.StatusBadRequest
	}
	output := &model.EquipmentMaintenanceOutput{
    	EqpMtnID:               equipmentMaintenance.EqpMtnID,
    	EqpMtnName:             equipmentMaintenance.EqpMtnName.String,
    	EqpMtnLocation:         equipmentMaintenance.EqpMtnLocation.String,
    	EqpMtnIssueDescription: equipmentMaintenance.EqpMtnIssueDescription.String,
    	EqpMtnReportedBy:       equipmentMaintenance.EqpMtnReportedBy.String,
    	EqpMtnPerformedBy:      equipmentMaintenance.EqpMtnPerformedBy.String,
    	EqpMtnDateReported:     equipmentMaintenance.EqpMtnDateReported.Format("2006-01-02"),
	}

	if equipmentMaintenance.EqpMtnDateFixed.Valid {
    	dateFixed := equipmentMaintenance.EqpMtnDateFixed.Time.Format("2006-01-02")
    		output.EqpMtnDateFixed = &dateFixed
	}

	if equipmentMaintenance.EqpMtnCost.Valid {
    	cost, err := strconv.ParseInt(equipmentMaintenance.EqpMtnCost.String, 10, 64)
    	if err == nil {
        	output.EqpMtnCost = &cost
    	}	
	}

	output.EqpMtnNote = equipmentMaintenance.EqpMtnNote.String

	return output, nil, http.StatusOK
}

func (s *sEquipmentMaintenance) UpdateEquipmentMaintenance(ctx context.Context, updateEquipmentMaintenance *model.UpdateEquipmentMaintenanceDto, Account *model.Account) (err error, statusCode int) {
	_, err = s.r.GetEquipmentMaintenance(ctx,database.GetEquipmentMaintenanceParams{
		EqpMtnID: updateEquipmentMaintenance.EqpMtnID,
		EqpMtnResID: Account.RestaurantID,
	})
	if err != nil {
		return fmt.Errorf("không tìm thấy đề xuất nội bộ"), http.StatusBadRequest
	}
	dateReported, err := time.Parse("2006-01-02", updateEquipmentMaintenance.EqpMtnDateReported)
	if err != nil {
		return fmt.Errorf("invalid date format for EqpMtnDateReported"), http.StatusBadRequest
	}

	var dateFixed sql.NullTime
	if updateEquipmentMaintenance.EqpMtnDateFixed != nil {
		parsedDate, err := time.Parse("2006-01-02", *updateEquipmentMaintenance.EqpMtnDateFixed)
		if err != nil {
			return fmt.Errorf("invalid date format for EqpMtnDateFixed"), http.StatusBadRequest
		}
		dateFixed = sql.NullTime{Time: parsedDate, Valid: true}
	}

	err = s.r.UpdateEquipmentMaintenance(ctx, database.UpdateEquipmentMaintenanceParams{
		EqpMtnID: updateEquipmentMaintenance.EqpMtnID,
		EqpMtnName: sql.NullString{String: updateEquipmentMaintenance.EqpMtnName, Valid: true},
		EqpMtnLocation: sql.NullString{String: updateEquipmentMaintenance.EqpMtnLocation, Valid: true},
		EqpMtnIssueDescription: sql.NullString{String: updateEquipmentMaintenance.EqpMtnIssueDescription, Valid: true},
		EqpMtnReportedBy: sql.NullString{String: updateEquipmentMaintenance.EqpMtnReportedBy, Valid: true},
		EqpMtnCost: sql.NullString{
			String: strconv.FormatInt(*updateEquipmentMaintenance.EqpMtnCost, 10),
			Valid: updateEquipmentMaintenance.EqpMtnCost != nil,
		},
		EqpMtnPerformedBy: sql.NullString{String: updateEquipmentMaintenance.EqpMtnPerformedBy, Valid: true},
		EqpMtnDateReported: dateReported,
		EqpMtnDateFixed: dateFixed,
		EqpMtnNote: sql.NullString{String: updateEquipmentMaintenance.EqpMtnNote, Valid: true},
		Updatedby: sql.NullString{String: Account.ID, Valid: true},
		EqpMtnResID: Account.RestaurantID,
	})
	if err != nil {
		return err, http.StatusInternalServerError
	}

	kafka.SendMessageToKafka(ctx, "NOTIFICATION_ACCOUNT_CREATE", kafka.NotificationPayload{
		RestaurantID: Account.RestaurantID,
		NotiContent:  fmt.Sprintf("Thiết bị %s vừa được cập nhật thông tin bảo trì", updateEquipmentMaintenance.EqpMtnName),
		NotiTitle:    "Cập nhật bảo trì thiết bị",
		NotiType:     "maintenance",
		NotiMetadata: `{"text":"maintenance updated"}`,
		SendObject:   "all_account",
	})

	return nil, http.StatusCreated
}

func (s *sEquipmentMaintenance) DeleteEquipmentMaintenance(ctx context.Context, EqpMtnID string, Account *model.Account) (err error, statusCode int) {
	equipmentMaintenance, err := s.r.GetEquipmentMaintenance(ctx,database.GetEquipmentMaintenanceParams{
		EqpMtnID: EqpMtnID,
		EqpMtnResID: Account.RestaurantID,
	})
	if err != nil {
		return fmt.Errorf("không tìm thấy đề xuất nội bộ"), http.StatusBadRequest
	}
	err = s.r.DeleteEquipmentMaintenance(ctx, database.DeleteEquipmentMaintenanceParams{
		EqpMtnID: EqpMtnID,
		Deletedby: sql.NullString{String: Account.ID, Valid: true},
		EqpMtnResID: Account.RestaurantID,
	})
	if err != nil {
		return err, http.StatusInternalServerError
	}

	kafka.SendMessageToKafka(ctx, "NOTIFICATION_ACCOUNT_CREATE", kafka.NotificationPayload{
		RestaurantID: Account.RestaurantID,
		NotiContent:  fmt.Sprintf("Thiết bị %s vừa bị xóa yêu cầu bảo trì", equipmentMaintenance.EqpMtnName.String),
		NotiTitle:    "Xóa bảo trì thiết bị",
		NotiType:     "maintenance",
		NotiMetadata: `{"text":"maintenance deleted"}`,
		SendObject:   "all_account",
	})



	return nil, http.StatusCreated
}

func (s *sEquipmentMaintenance) RestoreEquipmentMaintenance(ctx context.Context, EqpMtnID string, Account *model.Account) (err error, statusCode int) {
	equipmentMaintenance, err := s.r.GetEquipmentMaintenance(ctx,database.GetEquipmentMaintenanceParams{
		EqpMtnID: EqpMtnID,
		EqpMtnResID: Account.RestaurantID,
	})
	if err != nil {
		return fmt.Errorf("không tìm thấy đề xuất nội bộ"), http.StatusBadRequest
	}
	
	err = s.r.RestoreEquipmentMaintenance(ctx, database.RestoreEquipmentMaintenanceParams{
		EqpMtnID: EqpMtnID,
		EqpMtnResID: Account.RestaurantID,
	})
	if err != nil {
		return err, http.StatusInternalServerError
	}

	kafka.SendMessageToKafka(ctx, "NOTIFICATION_ACCOUNT_CREATE", kafka.NotificationPayload{
		RestaurantID: Account.RestaurantID,
		NotiContent:  fmt.Sprintf("Thiết bị %s vừa được khôi phục yêu cầu bảo trì", equipmentMaintenance.EqpMtnName.String),
		NotiTitle:    "Khôi phục bảo trì thiết bị",
		NotiType:     "maintenance",
		NotiMetadata: `{"text":"maintenance restored"}`,
		SendObject:   "all_account",
	})

	return nil, http.StatusCreated
}

func (s *sEquipmentMaintenance) GetAllEquipmentMaintenance(ctx context.Context, Limit int32, Offset int32, isDeleted int32, EqpMtnName string, Account *model.Account) (out response.ModelPagination[[]*model.EquipmentMaintenanceOutput], err error, statusCode int) {
	equipmentMaintenances, err := s.r.GetListEquipmentMaintenance(ctx, database.GetListEquipmentMaintenanceParams{
		Isdeleted:   sql.NullInt32{Int32: isDeleted, Valid: true},
		Isdeleted_2: sql.NullInt32{Int32: isDeleted, Valid: true},
		Limit:       Limit,
		Offset:      Offset,
		Column4: float64(Limit),
		EqpMtnResID: Account.RestaurantID,
		EqpMtnResID_2: Account.RestaurantID,
		EqpMtnName:      sql.NullString{String: "%" + EqpMtnName + "%", Valid: true},
		EqpMtnName_2:     sql.NullString{String: "%" + EqpMtnName + "%", Valid: true},
	})
	if err != nil {
		return response.ModelPagination[[]*model.EquipmentMaintenanceOutput]{}, err, http.StatusInternalServerError
	}
	var internalOutputs []*model.EquipmentMaintenanceOutput
	for _, user := range equipmentMaintenances {
		internalOutputs = append(internalOutputs, &model.EquipmentMaintenanceOutput{
			EqpMtnID:               user.EqpMtnID,
        	EqpMtnName:             user.EqpMtnName.String,
        	EqpMtnLocation:         user.EqpMtnLocation.String,
        	EqpMtnIssueDescription: user.EqpMtnIssueDescription.String,
        	EqpMtnStatus:           string(user.EqpMtnStatus.EquipmentMaintenanceEqpMtnStatus),
        	EqpMtnReportedBy:       user.EqpMtnReportedBy.String,
        	EqpMtnPerformedBy:      user.EqpMtnPerformedBy.String,
        	EqpMtnDateReported:     user.EqpMtnDateReported.Format("2006-01-02"),
        	EqpMtnDateFixed:        formatNullTime(user.EqpMtnDateFixed),
        	EqpMtnCost:             parseNullCost(user.EqpMtnCost),
        	EqpMtnNote:             user.EqpMtnNote.String,
		})
	}
	var totalPages, totalItems int32 = 0, 0
	if len(equipmentMaintenances) > 0 {
    	totalPages = int32(equipmentMaintenances[0].TotalPages.(float64))  
    	totalItems = int32(equipmentMaintenances[0].TotalItems)
	}

	return response.ModelPagination[[]*model.EquipmentMaintenanceOutput]{
		Result: internalOutputs,
		MetaPagination: response.MetaPagination{
			Current:    Offset,
			PageSize:   Limit,
			TotalPage:  totalPages,
			TotalItems: int64(totalItems),
		},
	}, nil, http.StatusOK

}

func(s *sEquipmentMaintenance) UpdateEquipmentMaintenanceStatus(ctx context.Context, updateEquipmentMaintenanceStatus *model.UpdateEquipmentMaintenanceStatusDto, Account *model.Account) (err error, statusCode int) {
	equipmentMaintenance, err := s.r.GetEquipmentMaintenance(ctx,database.GetEquipmentMaintenanceParams{
		EqpMtnID: updateEquipmentMaintenanceStatus.EqpMtnID,
		EqpMtnResID: Account.RestaurantID,
	})
	if err != nil {
		return fmt.Errorf("không tìm thấy đề xuất nội bộ"), http.StatusBadRequest
	}
	status := database.EquipmentMaintenanceEqpMtnStatus(updateEquipmentMaintenanceStatus.EqpMtnStatus)

	nullStatus := database.NullEquipmentMaintenanceEqpMtnStatus{
		EquipmentMaintenanceEqpMtnStatus: status,
		Valid:                             true,
	}

	err = s.r.UpdateEquipmentMaintenanceStatus(ctx, database.UpdateEquipmentMaintenanceStatusParams{
		EqpMtnID:     updateEquipmentMaintenanceStatus.EqpMtnID,
		EqpMtnStatus: nullStatus,
		EqpMtnResID:  Account.RestaurantID,
		Updatedby:         sql.NullString{String: Account.ID, Valid: true}, // nếu có
	})

	if err != nil {
		return err, http.StatusInternalServerError
	}

	kafka.SendMessageToKafka(ctx, "NOTIFICATION_ACCOUNT_CREATE", kafka.NotificationPayload{
		RestaurantID: Account.RestaurantID,
		NotiContent:  fmt.Sprintf("Thiết bị %s vừa được cập nhật trạng thái bảo trì thành %s", equipmentMaintenance.EqpMtnName.String, updateEquipmentMaintenanceStatus.EqpMtnStatus),
		NotiTitle:    "Cập nhật trạng thái bảo trì",
		NotiType:     "maintenance",
		NotiMetadata: `{"text":"status updated"}`,
		SendObject:   "all_account",
	})

	return nil, http.StatusCreated
}


func formatNullTime(t sql.NullTime) *string {
    if !t.Valid {
        return nil
    }
    formatted := t.Time.Format("2006-01-02")
    return &formatted
}

// Helper function để parse cost từ string sang int64
func parseNullCost(cost sql.NullString) *int64 {
    if !cost.Valid {
        return nil
    }
    costInt, err := strconv.ParseInt(cost.String, 10, 64)
    if err != nil {
        return nil
    }
    return &costInt
}