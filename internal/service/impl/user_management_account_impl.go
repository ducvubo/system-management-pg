package impl

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"system-management-pg/global"
	"system-management-pg/internal/database"
	"system-management-pg/internal/model"
	"system-management-pg/internal/utils/auth"
	"system-management-pg/internal/utils/crypto"
	"system-management-pg/internal/utils/random"
	"system-management-pg/internal/utils/validator"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type sUserManagementAccount struct {
	r *database.Queries
}

func NewUserManagementAccountImpl(r *database.Queries) *sUserManagementAccount {
	return &sUserManagementAccount{
		r: r,
	}
}

func (s *sUserManagementAccount) CreateUserManagementAccount(ctx context.Context, createUserManagementAccount *model.CreateUserManagementAccountDto, User *model.UserManagementProfileOutput) (err error, statusCode int) {
	if validator.IsStrongPassword(createUserManagementAccount.UsPass) == false {
		return fmt.Errorf("Mật khẩu phải chứa ít nhất 8 ký tự, 1 chữ hoa, 1 chữ thường, 1 số và 1 ký tự đặc biệt"), http.StatusBadRequest
	}

	_, err = s.r.GetUserProfile(ctx, createUserManagementAccount.UsId)
	if err != nil {
		return fmt.Errorf("user không tồn tại"), http.StatusBadRequest
	}

	_, err = s.r.FindUserAccountByEmail(ctx, createUserManagementAccount.UsEmail)
	if err == nil {
		return fmt.Errorf("email đã tồn tại"), http.StatusBadRequest
	}

	_, err = s.r.FindUserAccountById(ctx, createUserManagementAccount.UsId)
	if err == nil {
		return fmt.Errorf("User đã có tài khoản"), http.StatusBadRequest
	}

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100) + 1
	salt, err := crypto.GenerateSalt(randomNumber)
	if err != nil {
		return fmt.Errorf("lỗi tạo salt: %v", err), http.StatusInternalServerError
	}
	hashPassword := crypto.HashPassword(createUserManagementAccount.UsPass, salt)

	newUserAccount := database.CreateUserAccountParams{
		UsaID:         createUserManagementAccount.UsId,
		UsaEmail:      createUserManagementAccount.UsEmail,
		UsaPassword:   hashPassword,
		UsaSalt:       salt,
		UsaActiveTime: sql.NullTime{Time: time.Now(), Valid: true},
		UsaActive:     sql.NullInt32{Int32: 1, Valid: true},
		UsaLocked:     sql.NullInt32{Int32: 0, Valid: true},
		Createdby:     sql.NullString{String: User.UsID, Valid: true},
	}

	_, err = s.r.CreateUserAccount(ctx, newUserAccount)
	if err != nil {
		return fmt.Errorf("lỗi tạo tài khoản: %v", err), http.StatusInternalServerError
	}
	return nil, http.StatusCreated
}

func (s *sUserManagementAccount) FindUserAccountById(ctx context.Context, UsaID string) (out *model.UserManagementAccountOutput, err error, statusCode int) {
	user, err := s.r.FindUserAccountById(ctx, UsaID)
	if err != nil {
		return nil, err, http.StatusBadRequest
	}
	return &model.UserManagementAccountOutput{
		UsEmail: user.UsaEmail,
		UsId:    user.UsaID,
	}, nil, http.StatusOK
}

func (s *sUserManagementAccount) LoginUserManagementAccount(ctx context.Context, loginUserManagementAccount *model.LoginUserManagementAccountDto, clientId string) (loginUserManagementAccountOutput *model.LoginUserManagementAccountOutput, err error, statusCode int) {
	if validator.IsStrongPassword(loginUserManagementAccount.UsPass) == false {
		return nil, fmt.Errorf("Mật khẩu phải chứa ít nhất 8 ký tự, 1 chữ hoa, 1 chữ thường, 1 số và 1 ký tự đặc biệt"), http.StatusBadRequest
	}
	userAccount, err := s.r.FindUserAccountByEmail(ctx, loginUserManagementAccount.UsEmail)
	if err != nil {
		return nil, fmt.Errorf("Email hoặc mật khẩu không chính xác"), http.StatusUnauthorized
	}
	if !crypto.MatchingPassword(userAccount.UsaPassword, loginUserManagementAccount.UsPass, userAccount.UsaSalt) {
		return nil, fmt.Errorf("Email hoặc mật khẩu không chính xác"), http.StatusUnauthorized
	}

	userProfile, err := s.r.GetUserProfile(ctx, userAccount.UsaID)
	if err != nil {
		return nil, fmt.Errorf("Email hoặc mật khẩu không chính xác"), http.StatusUnauthorized
	}

	if userProfile.Isdeleted.Int32 == 1 {
		return nil, fmt.Errorf("Tài khoản đã bị xóa, vui lòng liên hệ admin để được hỗ trợ"), http.StatusUnauthorized
	}

	if userAccount.UsaLocked.Int32 == 1 {
		return nil, fmt.Errorf("Tài khoản đã bị khóa, vui lòng liên hệ admin để được hỗ trợ"), http.StatusUnauthorized
	}

	if userAccount.UsaActive.Int32 == 0 {
		return nil, fmt.Errorf("Tài khoản chưa được kích hoạt, vui lòng liên hệ admin để được hỗ trợ"), http.StatusUnauthorized
	}

	ussId := uuid.New().String()
	subToken := auth.GenerateCliTokenUUID(ussId)

	payloadAccessToken := auth.PayloadClaimsToken{
		StandardClaims: jwt.StandardClaims{
			Subject:   subToken,
			Id:        subToken,
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "system-management-pg",
			Audience:  "system-management-pg",
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	}

	payloadRefreshToken := auth.PayloadClaimsToken{
		StandardClaims: jwt.StandardClaims{
			Subject:   subToken,
			Id:        subToken,
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			Issuer:    "system-management-pg",
			Audience:  "system-management-pg",
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	}

	keyAccessToken := global.Config.JWT.API_SECRET_KEY + clientId + random.GenerateRandomString(10)
	keyRefreshToken := random.GenerateRandomString(10) + global.Config.JWT.API_SECRET_KEY + clientId

	accessToken, err := auth.GenerateToken(payloadAccessToken, time.Hour*24, keyAccessToken)
	if err != nil {
		return nil, fmt.Errorf("lỗi tạo token: %v", err), http.StatusInternalServerError
	}

	refreshToken, err := auth.GenerateToken(payloadRefreshToken, time.Hour*24*30, keyRefreshToken)
	if err != nil {
		return nil, fmt.Errorf("lỗi tạo token: %v", err), http.StatusInternalServerError
	}

	err = s.r.DeleteUserSessionByClientIdAndUsaId(ctx, database.DeleteUserSessionByClientIdAndUsaIdParams{
		UssClientID: clientId,
		UsaID:       userAccount.UsaID,
	})
	if err != nil {
		return nil, fmt.Errorf("Đã có lỗi xảy ra: %v", err), http.StatusUnauthorized
	}

	_, err = s.r.CreateUserSession(ctx, database.CreateUserSessionParams{
		UssID:       ussId,
		UsaID:       userAccount.UsaID,
		UssRf:       refreshToken,
		UssKeyAt:    keyAccessToken,
		UssKeyRf:    keyRefreshToken,
		UssClientID: clientId,
	})
	if err != nil {
		return nil, fmt.Errorf("Đăng nhập không thành công vui lòng thử lại sau !!!"), http.StatusInternalServerError
	}

	return &model.LoginUserManagementAccountOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ClientId:     clientId,
	}, nil, http.StatusCreated
}

func (s *sUserManagementAccount) FindUserSessionBySessionIdAndRefreshToken(ctx context.Context, clientId, refreshToken string) (userSession database.FindUserSessionBySessionIdAndRefreshTokenRow, err error, statusCode int) {
	userSession, err = s.r.FindUserSessionBySessionIdAndRefreshToken(ctx, database.FindUserSessionBySessionIdAndRefreshTokenParams{
		UssClientID: clientId,
		UssRf:       refreshToken,
	})
	if err != nil {
		return userSession, fmt.Errorf("Không tìm thấy phiên đăng nhập"), http.StatusUnauthorized
	}
	return userSession, nil, 0
}

func (s *sUserManagementAccount) RefreshToken(ctx context.Context, refreshTokenInput *model.RefreshTokenInput, clientId string) (*model.LoginUserManagementAccountOutput, error, int) {

	UserSession, err := s.r.FindUserSessionBySessionIdAndRefreshToken(ctx, database.FindUserSessionBySessionIdAndRefreshTokenParams{
		UssClientID: clientId,
		UssRf:       refreshTokenInput.RefreshToken,
	})

	if err != nil {
		return nil, fmt.Errorf("Phiên đăng nhập không tồn tại"), http.StatusUnauthorized
	}

	UserProfile, err := s.r.GetUserProfile(ctx, UserSession.UsaID)

	if err != nil {
		return nil, fmt.Errorf("Không xác định được tài khoan"), http.StatusUnauthorized
	}

	User := model.UserManagementProfileOutput{
		UsID:      UserProfile.UsID,
		Isdeleted: UserProfile.Isdeleted.Int32,
	}

	if User.Isdeleted == 1 {
		return nil, fmt.Errorf("Tài khoản đã bị xóa, vui lòng liên hệ admin để được hỗ trợ"), http.StatusUnauthorized
	}

	userAccount, err := s.r.FindUserAccountById(ctx, User.UsID)

	if err != nil {
		return nil, fmt.Errorf("Không xác định được tài khoan"), http.StatusUnauthorized
	}

	if userAccount.UsaLocked.Int32 == 1 {
		return nil, fmt.Errorf("Tài khoản đã bị khóa, vui lòng liên hệ admin để được hỗ trợ"), http.StatusUnauthorized
	}

	if userAccount.UsaActive.Int32 == 0 {
		return nil, fmt.Errorf("Tài khoản chưa được kích hoạt, vui lòng liên hệ admin để được hỗ trợ"), http.StatusUnauthorized
	}

	ussId := uuid.New().String()
	subToken := auth.GenerateCliTokenUUID(ussId)

	payloadAccessToken := auth.PayloadClaimsToken{
		StandardClaims: jwt.StandardClaims{
			Subject:   subToken,
			Id:        subToken,
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "system-management-pg",
			Audience:  "system-management-pg",
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	}

	payloadRefreshToken := auth.PayloadClaimsToken{
		StandardClaims: jwt.StandardClaims{
			Subject:   subToken,
			Id:        subToken,
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			Issuer:    "system-management-pg",
			Audience:  "system-management-pg",
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	}

	keyAccessToken := global.Config.JWT.API_SECRET_KEY + clientId + random.GenerateRandomString(10)
	keyRefreshToken := random.GenerateRandomString(10) + global.Config.JWT.API_SECRET_KEY + clientId

	accessToken, err := auth.GenerateToken(payloadAccessToken, time.Hour*24, keyAccessToken)
	if err != nil {
		return nil, fmt.Errorf("lỗi tạo token: %v", err), http.StatusInternalServerError
	}

	refreshToken, err := auth.GenerateToken(payloadRefreshToken, time.Hour*24*30, keyRefreshToken)
	if err != nil {
		return nil, fmt.Errorf("lỗi tạo token: %v", err), http.StatusInternalServerError
	}

	err = s.r.DeleteUserSessionByClientIdAndUsaId(ctx, database.DeleteUserSessionByClientIdAndUsaIdParams{
		UssClientID: clientId,
		UsaID:       userAccount.UsaID,
	})
	if err != nil {
		return nil, fmt.Errorf("Đã có lỗi xảy ra: %v", err), http.StatusUnauthorized
	}

	_, err = s.r.CreateUserSession(ctx, database.CreateUserSessionParams{
		UssID:       ussId,
		UsaID:       userAccount.UsaID,
		UssRf:       refreshToken,
		UssKeyAt:    keyAccessToken,
		UssKeyRf:    keyRefreshToken,
		UssClientID: clientId,
	})
	if err != nil {
		return nil, fmt.Errorf("Đăng nhập không thành công vui lòng thử lại sau !!!"), http.StatusInternalServerError
	}

	return &model.LoginUserManagementAccountOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ClientId:     clientId,
	}, nil, http.StatusCreated
}
