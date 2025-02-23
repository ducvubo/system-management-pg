package impl

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"system-management-pg/internal/database"
	"system-management-pg/internal/model"
	"system-management-pg/pkg/response"
	"time"

	"github.com/google/uuid"
)

type sUserManagementProfile struct {
	r *database.Queries
}

func NewUserManagementProfileImpl(r *database.Queries) *sUserManagementProfile {
	return &sUserManagementProfile{
		r: r,
	}
}

func (s *sUserManagementProfile) CreateUserManagementProfile(ctx context.Context, createUserManagementProfile *model.CreateUserManagementProfileDto, User *model.UserManagementProfileOutput) (err error, statusCode int) {
	layout := time.RFC3339 // Định dạng ISO 8601
	parsedTime, err := time.Parse(layout, createUserManagementProfile.UsBirthday)
	if err != nil {
		log.Println("lỗi parse ngày sinh:", err)
		return fmt.Errorf("ngày sinh không đúng định dạng (ISO 8601): %v", err), http.StatusBadRequest
	}

	newUser := database.CreateUserProfileParams{
		UsID:       uuid.New().String(),
		UsName:     sql.NullString{String: createUserManagementProfile.UsName, Valid: true},
		UsAvatar:   sql.NullString{String: createUserManagementProfile.UsAvatar, Valid: true},
		UsPhone:    sql.NullString{String: createUserManagementProfile.UsPhone, Valid: true},
		UsGender:   sql.NullString{String: createUserManagementProfile.UsGender, Valid: true},
		UsAddress:  sql.NullString{String: createUserManagementProfile.UsAddress, Valid: true},
		UsBirthday: sql.NullTime{Time: parsedTime, Valid: true},
		Createdby:  sql.NullString{String: User.UsID, Valid: true},
	}
	_, err = s.r.CreateUserProfile(ctx, newUser)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	return nil, http.StatusCreated
}

func (s *sUserManagementProfile) FindUserManagementProfile(ctx context.Context, UsID string) (out *model.UserManagementProfileOutput, err error, statusCode int) {
	user, err := s.r.GetUserProfile(ctx, UsID)
	if err != nil {
		return nil, err, http.StatusBadRequest
	}
	return &model.UserManagementProfileOutput{
		UsID:       user.UsID,
		UsName:     user.UsName.String,
		UsAvatar:   user.UsAvatar.String,
		UsPhone:    user.UsPhone.String,
		UsGender:   user.UsGender.String,
		UsAddress:  user.UsAddress.String,
		UsBirthday: user.UsBirthday.Time.Format(time.RFC3339),
		Isdeleted:  user.Isdeleted.Int32,
	}, nil, http.StatusOK
}

func (s *sUserManagementProfile) UpdateUserManagementProfile(ctx context.Context, updateUserManagementProfile *model.UpdateUserManagementProfileDto, User *model.UserManagementProfileOutput) (err error, statusCode int) {
	_, err = s.r.GetUserProfile(ctx, updateUserManagementProfile.UsID)
	if err != nil {
		return fmt.Errorf("không tìm thấy user"), http.StatusBadRequest
	}
	layout := time.RFC3339
	parsedTime, err := time.Parse(layout, updateUserManagementProfile.UsBirthday)
	if err != nil {
		log.Println("Lỗi parse ngày sinh:", err)
		return fmt.Errorf("Ngày sinh không đúng định dạng (ISO 8601): %v", err), http.StatusBadRequest
	}

	err = s.r.UpdateUserProfile(ctx, database.UpdateUserProfileParams{
		UsID:       updateUserManagementProfile.UsID,
		UsName:     sql.NullString{String: updateUserManagementProfile.UsAvatar, Valid: true},
		UsAvatar:   sql.NullString{String: updateUserManagementProfile.UsPhone, Valid: true},
		UsPhone:    sql.NullString{String: updateUserManagementProfile.UsGender, Valid: true},
		UsGender:   sql.NullString{String: updateUserManagementProfile.UsAddress, Valid: true},
		UsAddress:  sql.NullString{String: updateUserManagementProfile.UsBirthday, Valid: true},
		UsBirthday: sql.NullTime{Time: parsedTime, Valid: true},
		Updatedby:  sql.NullString{String: User.UsID, Valid: true},
	})
	if err != nil {
		return err, http.StatusInternalServerError
	}
	return nil, http.StatusCreated
}

func (s *sUserManagementProfile) DeleteUserManagementProfile(ctx context.Context, UsID string, User *model.UserManagementProfileOutput) (err error, statusCode int) {
	_, err = s.r.GetUserProfile(ctx, UsID)
	if err != nil {
		return fmt.Errorf("không tìm thấy user"), http.StatusBadRequest
	}
	err = s.r.DeleteUserProfile(ctx, database.DeleteUserProfileParams{
		UsID:      UsID,
		Deletedby: sql.NullString{String: User.UsID, Valid: true},
	})
	if err != nil {
		return err, http.StatusInternalServerError
	}
	return nil, http.StatusCreated
}

func (s *sUserManagementProfile) RestoreUserManagementProfile(ctx context.Context, UsID string, User *model.UserManagementProfileOutput) (err error, statusCode int) {
	_, err = s.r.GetUserProfile(ctx, UsID)
	if err != nil {
		return fmt.Errorf("không tìm thấy user"), http.StatusBadRequest
	}
	
	err = s.r.RestoreUserProfile(ctx, UsID)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	return nil, http.StatusCreated
}

func (s *sUserManagementProfile) GetAllUserManagementProfile(ctx context.Context, Limit int32, Offset int32, isDeleted int32, UsName string) (out response.ModelPagination[[]*model.UserManagementProfileOutput], err error, statusCode int) {
	users, err := s.r.GetListUserProfile(ctx, database.GetListUserProfileParams{
		Isdeleted:   sql.NullInt32{Int32: isDeleted, Valid: true},
		Column3:     float64(Limit),
		Isdeleted_2: sql.NullInt32{Int32: isDeleted, Valid: true},
		Limit:       Limit,
		Offset:      Offset,
		UsName:      sql.NullString{String: "%" + UsName + "%", Valid: true},
		UsName_2:    sql.NullString{String: "%" + UsName + "%", Valid: true},
	})
	if err != nil {
		return response.ModelPagination[[]*model.UserManagementProfileOutput]{}, err, http.StatusInternalServerError
	}
	var userOutputs []*model.UserManagementProfileOutput
	for _, user := range users {
		userOutputs = append(userOutputs, &model.UserManagementProfileOutput{
			UsID:       user.UsID,
			UsName:     user.UsName.String,
			UsAvatar:   user.UsAvatar.String,
			UsPhone:    user.UsPhone.String,
			UsGender:   user.UsGender.String,
			UsAddress:  user.UsAddress.String,
			UsBirthday: user.UsBirthday.Time.Format(time.RFC3339),
		})
	}
	// Nếu không có user nào, đặt TotalPage và TotalItems về 0
	var totalPages, totalItems int32 = 0, 0
	if len(users) > 0 {
		totalPages = users[0].TotalPages
		totalItems = int32(users[0].TotalItems)
	}

	return response.ModelPagination[[]*model.UserManagementProfileOutput]{
		Result: userOutputs,
		MetaPagination: response.MetaPagination{
			Current:    Offset,
			PageSize:   Limit,
			TotalPage:  totalPages,
			TotalItems: int64(totalItems),
		},
	}, nil, http.StatusOK

}
