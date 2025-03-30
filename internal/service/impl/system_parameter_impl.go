package impl

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	// consts "system-management-pg/internal/const"
	"system-management-pg/global"
	"system-management-pg/internal/database"
	"system-management-pg/internal/model"

	"github.com/redis/go-redis/v9"
)

type sSystemParameter struct {
	r *database.Queries
}

func NewSystemParameterImpl(r *database.Queries) *sSystemParameter {
	return &sSystemParameter{
		r: r,
	}
}

// func (s *sSystemParameter) SaveSystemParameter(ctx context.Context, saveSystemParameterDto *model.SaveSystemParameterDto, User *model.UserManagementProfileOutput) (err error, statusCode int) {

// 	// _, exists := consts.GetSystemParameter(saveSystemParameterDto.SysParaID)

// 	// if exists {
// 	// 	return fmt.Errorf("Tham số không tồn tại"), http.StatusNotFound
// 	// }

// 	_, err = s.r.SaveSystemParameter(ctx, database.SaveSystemParameterParams{
// 		SysParaID:            saveSystemParameterDto.SysParaID,
// 		SysParaDescription:   sql.NullString{Valid: true, String: saveSystemParameterDto.SysParaDescription},
// 		SysParaValue:         saveSystemParameterDto.SysParaValue,
// 		Createdby:            sql.NullString{Valid: true, String: User.UsID},
// 		Updatedby:            sql.NullString{Valid: true, String: User.UsID},
// 		SysParaValue_2:       saveSystemParameterDto.SysParaValue,
// 		Updatedby_2:          sql.NullString{Valid: true, String: User.UsID},
// 		SysParaDescription_2: sql.NullString{Valid: true, String: saveSystemParameterDto.SysParaDescription},
// 	})

// 	if err != nil {
// 		return fmt.Errorf("Lưu tham số không thành công"), http.StatusInternalServerError
// 	}

// 	return nil, http.StatusCreated
// }

func (s *sSystemParameter) SaveSystemParameter(ctx context.Context, saveSystemParameterDto *model.SaveSystemParameterDto, User *model.UserManagementProfileOutput) (err error, statusCode int) {
	_, err = s.r.SaveSystemParameter(ctx, database.SaveSystemParameterParams{
		SysParaID:            saveSystemParameterDto.SysParaID,
		SysParaDescription:   sql.NullString{Valid: true, String: saveSystemParameterDto.SysParaDescription},
		SysParaValue:         saveSystemParameterDto.SysParaValue,
		Createdby:            sql.NullString{Valid: true, String: User.UsID},
		Updatedby:            sql.NullString{Valid: true, String: User.UsID},
		SysParaValue_2:       saveSystemParameterDto.SysParaValue,
		Updatedby_2:          sql.NullString{Valid: true, String: User.UsID},
		SysParaDescription_2: sql.NullString{Valid: true, String: saveSystemParameterDto.SysParaDescription},
	})

	if err != nil {
		return fmt.Errorf("Lưu tham số không thành công"), http.StatusInternalServerError
	}

	// Cập nhật lại cache sau khi lưu DB thành công
	go func() {
		cacheKey := "system_parameters"
		systemParameters, err := s.r.GetAllSystemParameters(ctx)
		if err != nil {
			return // Nếu lỗi, bỏ qua việc cập nhật cache
		}

		var out []*model.SaveSystemParameterDto
		for _, systemParameter := range systemParameters {
			out = append(out, &model.SaveSystemParameterDto{
				SysParaID:          systemParameter.SysParaID,
				SysParaDescription: systemParameter.SysParaDescription.String,
				SysParaValue:       systemParameter.SysParaValue,
			})
		}

		cacheBytes, err := json.Marshal(out)
		if err == nil {
			_ = global.Rdb.Set(ctx, cacheKey, cacheBytes, 10*time.Minute).Err()
		}
	}()

	return nil, http.StatusCreated
}


func (s *sSystemParameter) GetSystemParameter(ctx context.Context, sysParaID string) (out *model.SaveSystemParameterDto, err error, statusCode int) {
	// _, exists := consts.GetSystemParameter(sysParaID)

	// if !exists {
	// 	return nil, fmt.Errorf("Tham số không tồn tại"), http.StatusBadRequest
	// }

	systemParameter, err := s.r.GetSystemParameter(ctx, sysParaID)
	if err != nil {
		return nil, fmt.Errorf("Lấy tham số không thành công"), http.StatusBadRequest
	}

	out = &model.SaveSystemParameterDto{
		SysParaID:          systemParameter.SysParaID,
		SysParaDescription: systemParameter.SysParaDescription.String,
		SysParaValue:       systemParameter.SysParaValue,
	}

	return out, nil, http.StatusOK
}

func (s *sSystemParameter) GetAllSystemParameter(ctx context.Context) (out []*model.SaveSystemParameterDto, err error, statusCode int) {
	cacheKey := "system_parameters"

	cacheData, err := global.Rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(cacheData), &out); err == nil {
			return out, nil, http.StatusOK
		}
	} else if err != redis.Nil {
		return nil, fmt.Errorf("lỗi khi truy vấn cache: %v", err), http.StatusInternalServerError
	}

	// Nếu không có cache, lấy từ DB
	systemParameters, err := s.r.GetAllSystemParameters(ctx)
	if err != nil {
		return nil, fmt.Errorf("Lấy tham số không thành công"), http.StatusBadRequest
	}

	for _, systemParameter := range systemParameters {
		out = append(out, &model.SaveSystemParameterDto{
			SysParaID:          systemParameter.SysParaID,
			SysParaDescription: systemParameter.SysParaDescription.String,
			SysParaValue:       systemParameter.SysParaValue,
		})
	}

	cacheBytes, err := json.Marshal(out)
	if err == nil {
		_ = global.Rdb.Set(ctx, cacheKey, cacheBytes, 10*time.Minute).Err()
	}

	return out, nil, http.StatusOK
}
