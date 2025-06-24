package impl

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"system-management-pg/global"
	"system-management-pg/internal/database"
	"system-management-pg/internal/model"
	"system-management-pg/internal/utils/auth"
	"system-management-pg/internal/utils/crypto"
	"system-management-pg/internal/utils/image"
	// "system-management-pg/internal/utils/kafka"
	"system-management-pg/internal/utils/random"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/redis/go-redis/v9"
)

type sUserPatoAccount struct {
	r *database.Queries
}

func NewUserPatoAccountImpl(r *database.Queries) *sUserPatoAccount {
	return &sUserPatoAccount{
		r: r,
	}
}

type KafkaEmailMessage struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Text    string `json:"text"`
	Link    string `json:"link"`
}

func BuildKafkaMessage(email string, otp int, urlRedirect string) (string, error) {
	u, err := url.Parse(urlRedirect)
	if err != nil {
		return "", fmt.Errorf("invalid url_redirect: %w", err)
	}

	query := u.Query()
	query.Set("email", email)
	query.Set("otp", fmt.Sprintf("%d", otp))
	u.RawQuery = query.Encode()

	linkConfirm := u.String()

	// Tạo message
	msg := KafkaEmailMessage{
		To:      email,
		Subject: "Xác nhận đăng ký tài khoản",
		Text:    "Bạn nhận được email này vì bạn đã đăng ký tài khoản tại PATO. Để xác nhận, vui lòng nhấp vào liên kết bên dưới. Nếu bạn không phải là người nhận email này, vui lòng bỏ qua nó.",
		Link:    linkConfirm,
	}

	// Encode sang JSON
	jsonBytes, err := json.Marshal(msg)
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %w", err)
	}

	return string(jsonBytes), nil
}

func (s *sUserPatoAccount) RegisterUserPatoAccount(ctx context.Context, registerUserPatoAccount *model.RegisterUserPatoAccountDto) (registerUserPatoAccountOutput *model.RegisterUserPatoAccountOutput, err error, statusCode int) {
	_, err = s.r.FindUserPatoAccountByEmailAndType(ctx, database.FindUserPatoAccountByEmailAndTypeParams{
		UsaPtEmail: registerUserPatoAccount.UsPtEmail,
		UsaPtType:  sql.NullString{String: "system", Valid: true},
	})

	if err != nil && err != sql.ErrNoRows {
		return nil, err, http.StatusBadRequest
	}

	if err == nil {
		return nil, fmt.Errorf("Email đã tồn tại vui lòng sử dụng email khác"), http.StatusBadRequest
	}

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100) + 1
	salt, err := crypto.GenerateSalt(randomNumber)
	if err != nil {
		return nil, fmt.Errorf("lỗi tạo salt: %v", err), http.StatusInternalServerError
	}
	hashPassword := crypto.HashPassword(registerUserPatoAccount.UsPtPass, salt)

	buildImage, err := image.BuildImageJson(image.ImageData{
		ImageCloud:  "/api/view-image?bucket=default&file=default-avatar-profile-icon-social-media-user-image-gray-avatar-icon-blank-profile-silhouette-illustration-vector.jpg",
		ImageCustom: "/api/view-image?bucket=default&file=default-avatar-profile-icon-social-media-user-image-gray-avatar-icon-blank-profile-silhouette-illustration-vector.jpg",
	})
	if err != nil {
		return nil, fmt.Errorf("lỗi chuyển đổi JSON: %v", err), http.StatusInternalServerError
	}

	patoProfile, err := s.r.CreateUserPatoProfile(ctx, database.CreateUserPatoProfileParams{
		UsPtName:     sql.NullString{String: registerUserPatoAccount.UsPtEmail, Valid: true},
		UsPtAvatar:   sql.NullString{String: buildImage, Valid: true},
		UsPtPhone:    sql.NullString{String: "Không xác định", Valid: true},
		UsPtGender:   sql.NullString{String: "Không xác định", Valid: true},
		UsPtAddress:  sql.NullString{String: "Không xác định", Valid: true},
		UsPtBirthday: sql.NullTime{Time: time.Time{}, Valid: false},
	})

	patoProfileID, err := patoProfile.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("lỗi lấy ID tài khoản: %v", err), http.StatusInternalServerError
	}

	_, err = s.r.CreateUserPatoAccount(ctx, database.CreateUserPatoAccountParams{
		UsaPtID:       int32(patoProfileID),
		UsaPtEmail:    registerUserPatoAccount.UsPtEmail,
		UsaPtPassword: hashPassword,
		UsaPtSalt:     salt,
		UsaPtType:     sql.NullString{String: "system", Valid: true},
		UsaPtActive:   sql.NullInt32{Int32: 1, Valid: true},
		UsaPtLocked:   sql.NullInt32{Int32: 0, Valid: true},
	})

	if err != nil {
		return nil, fmt.Errorf("lỗi tạo tài khoản: %v", err), http.StatusInternalServerError
	}

	if err != nil {
		return nil, fmt.Errorf("lỗi tạo hồ sơ tài khoản: %v", err), http.StatusInternalServerError
	}
	registerUserPatoAccountOutput = &model.RegisterUserPatoAccountOutput{
		UsPtEmail: registerUserPatoAccount.UsPtEmail,
	}
	otp := random.GenerateSixDigitOtp()
	fmt.Printf("Otp is :::%d\n", otp)
	hashOtp := crypto.GetHash(fmt.Sprintf("%d", otp))
	hashEmail := crypto.GetHash(registerUserPatoAccount.UsPtEmail)

	cacheKey := fmt.Sprintf("otp:%s:%s", hashEmail, hashOtp)
	err = global.Rdb.Set(ctx, cacheKey, otp, 10*time.Minute).Err()

	if err != nil {
		return nil, fmt.Errorf("lỗi lưu OTP vào cache: %v", err), http.StatusInternalServerError
	}

	// message, err := BuildKafkaMessage(registerUserPatoAccount.UsPtEmail, otp, registerUserPatoAccount.UrlRedirect)
	// if err != nil {
	// 	return nil, fmt.Errorf("lỗi tạo message JSON gửi Kafka: %v", err), http.StatusInternalServerError
	// }

	// err = kafka.SendMessageToKafka("CONFIRM_REGISTER_USER_ACCOUNT_PATO", message)
	// if err != nil {
	// 	return nil, fmt.Errorf("lỗi gửi OTP đến Kafka: %v", err), http.StatusInternalServerError
	// }

	return registerUserPatoAccountOutput, nil, http.StatusOK
}

// resend otp
func (s *sUserPatoAccount) ResendOtp(ctx context.Context, resendOtp *model.ResendOtpDto) (err error, statusCode int) {
	_, err = s.r.FindUserPatoAccountByEmailAndType(ctx, database.FindUserPatoAccountByEmailAndTypeParams{
		UsaPtEmail: resendOtp.UsPtEmail,
		UsaPtType:  sql.NullString{String: "system", Valid: true},
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("Email không tồn tại vui lòng kiểm tra lại"), http.StatusBadRequest
		}
		return fmt.Errorf("lỗi tìm kiếm tài khoản: %v", err), http.StatusInternalServerError
	}
	otp := random.GenerateSixDigitOtp()
	fmt.Printf("Otp is :::%d\n", otp)
	hashOtp := crypto.GetHash(fmt.Sprintf("%d", otp))
	hashEmail := crypto.GetHash(resendOtp.UsPtEmail)
	cacheKey := fmt.Sprintf("otp:%s:%s", hashEmail, hashOtp)
	err = global.Rdb.Set(ctx, cacheKey, otp, 10*time.Minute).Err()
	if err != nil {
		return fmt.Errorf("lỗi lưu OTP vào cache: %v", err), http.StatusInternalServerError
	}
	// message, err := BuildKafkaMessage(resendOtp.UsPtEmail, otp, resendOtp.UrlRedirect)
	// if err != nil {
	// 	return fmt.Errorf("lỗi tạo message JSON gửi Kafka: %v", err), http.StatusInternalServerError
	// }
	// err = kafka.SendMessageToKafka("CONFIRM_REGISTER_USER_ACCOUNT_PATO", message)
	// if err != nil {
	// 	return fmt.Errorf("lỗi gửi OTP đến Kafka: %v", err), http.StatusInternalServerError
	// }
	return nil, http.StatusOK
}

// activate account
func (s *sUserPatoAccount) ActivateAccount(ctx context.Context, activateAccount *model.ActiveAccountDto) (err error, statusCode int) {
	// 1. check otp in cache
	hashEmail := crypto.GetHash(activateAccount.UsPtEmail)
	// hashOtp := crypto.GetHash(activateAccount.UsPtOtp)
	//chuyển đổi mã OTP thành chuỗi
	hashOtp := crypto.GetHash(fmt.Sprintf("%d", activateAccount.UsPtOtp))
	cacheKey := fmt.Sprintf("otp:%s:%s", hashEmail, hashOtp)
	otp, err := global.Rdb.Get(ctx, cacheKey).Result()
	if err != nil {
		if err == redis.Nil {
			return fmt.Errorf("Mã OTP không tồn tại hoặc đã hết hạn"), http.StatusBadRequest
		}
		return fmt.Errorf("lỗi tìm kiếm mã OTP: %v", err), http.StatusInternalServerError
	}

	if otp != fmt.Sprintf("%d", activateAccount.UsPtOtp) {
		return fmt.Errorf("Mã OTP không đúng"), http.StatusBadRequest
	}

	// 2. check email in db
	userExist, err := s.r.FindUserPatoAccountByEmailAndType(ctx, database.FindUserPatoAccountByEmailAndTypeParams{
		UsaPtEmail: activateAccount.UsPtEmail,
		UsaPtType:  sql.NullString{String: "system", Valid: true},
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("Email không tồn tại vui lòng kiểm tra lại"), http.StatusBadRequest
		}
		return fmt.Errorf("lỗi tìm kiếm tài khoản: %v", err), http.StatusInternalServerError
	}

	//3 ActiveUserPatoAccount
	err = s.r.ActiveUserPatoAccount(ctx, database.ActiveUserPatoAccountParams{
		Updatedby: sql.NullString{String: activateAccount.UsPtEmail, Valid: true},
		UsaPtID:   userExist.UsaPtID,
	})

	if err != nil {
		return fmt.Errorf("lỗi kích hoạt tài khoản: %v", err), http.StatusInternalServerError
	}

	// 4. delete otp in cache
	err = global.Rdb.Del(ctx, cacheKey).Err()
	if err != nil {
		return fmt.Errorf("lỗi xóa mã OTP trong cache: %v", err), http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

// Login
func (s *sUserPatoAccount) Login(ctx context.Context, login *model.LoginUserPatoAccountDto, clientId string) (loginUserPatoAccountOutput *model.LoginUserPatoAccountOutput, err error, statusCode int) {
	// avtive == 0 là kích hoạt rồi
	// locked == 0 là chưa khóa
	user, err := s.r.FindUserPatoAccountByEmailAndType(ctx, database.FindUserPatoAccountByEmailAndTypeParams{
		UsaPtEmail: login.UsPtEmail,
		UsaPtType:  sql.NullString{String: "system", Valid: true},
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Email không tồn tại vui lòng kiểm tra lại"), http.StatusBadRequest
		}
		return nil, fmt.Errorf("lỗi tìm kiếm tài khoản: %v", err), http.StatusInternalServerError
	}
	if user.UsaPtPassword != crypto.HashPassword(login.UsPtPass, user.UsaPtSalt) {
		return nil, fmt.Errorf("Mật khẩu không đúng"), http.StatusBadRequest
	}
	if user.UsaPtActive.Int32 == 1 {
		return nil, fmt.Errorf("Tài khoản chưa được kích hoạt vui lòng kiểm tra email"), http.StatusBadRequest
	}
	if user.UsaPtLocked.Int32 == 1 {
		return nil, fmt.Errorf("Tài khoản đã bị khóa vui lòng liên hệ admin"), http.StatusBadRequest
	}
	accessPayload := auth.PayloadClaimsToken{
		StandardClaims: jwt.StandardClaims{
			Subject:   fmt.Sprintf("%d", user.UsaPtID),
			Audience:  user.UsaPtEmail,
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	}

	secretKey := global.Config.JWT.API_SECRET_KEY
	accessToken, err := auth.GenerateToken(accessPayload, 1000*time.Minute, secretKey)
	if err != nil {
		return nil, fmt.Errorf("lỗi tạo access token: %v", err), http.StatusInternalServerError
	}

	// Tạo payload cho refresh token
	refreshPayload := auth.PayloadClaimsToken{
		StandardClaims: jwt.StandardClaims{
			Subject:   fmt.Sprintf("%d", user.UsaPtID), // ID người dùng
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	}

	refreshToken, err := auth.GenerateToken(refreshPayload, 7*24*time.Hour, secretKey)
	if err != nil {
		return nil, fmt.Errorf("lỗi tạo refresh token: %v", err), http.StatusInternalServerError
	}

	loginUserPatoAccountOutput = &model.LoginUserPatoAccountOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return loginUserPatoAccountOutput, nil, http.StatusOK
}
