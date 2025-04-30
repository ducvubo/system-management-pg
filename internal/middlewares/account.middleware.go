package middlewares

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"strings"

	"system-management-pg/global"
	"system-management-pg/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

const (
	EmployeeIndex         = "employees"
	RestaurantIndex       = "restaurants"
	AccountIndex          = "accounts"
	RefreshTokenIndex     = "refresh-tokens-account"
	UnauthorizedErrorMsg  = "Token không hợp lệ"
	UnauthorizedErrorCode = -10
)

// TokenData represents refresh token data
type TokenData struct {
	RefreshTokenPublicKey string `json:"rf_public_key_refresh_token"`
	AccessTokenPublicKey  string `json:"rf_public_key_access_token"`
}

// Custom error type for unauthorized access
type UnauthorizedError struct {
	Message string
	Code    int
}

func (e *UnauthorizedError) Error() string {
	return e.Message
}

// FindRefreshToken searches for refresh token in Elasticsearch
func FindRefreshToken(ctx context.Context, refreshToken string) (*TokenData, error) {
	if !indexExists(ctx, RefreshTokenIndex) {
		return nil, &UnauthorizedError{Message: UnauthorizedErrorMsg + ": refresh-tokens-account index not found", Code: UnauthorizedErrorCode}
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"rf_refresh_token": refreshToken,
			},
		},
	}

	queryBytes, err := json.Marshal(query)
	if err != nil {
		global.Logger.Error("Failed to marshal query", zap.Error(err))
		return nil, err
	}

	var result map[string]interface{}
	res, err := global.EsClient.Search(
		global.EsClient.Search.WithContext(ctx),
		global.EsClient.Search.WithIndex(RefreshTokenIndex),
		global.EsClient.Search.WithBody(bytes.NewReader(queryBytes)),
	)
	if err != nil {
		global.Logger.Error("Elasticsearch search error", zap.Error(err))
		return nil, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		global.Logger.Error("Failed to decode Elasticsearch response", zap.Error(err))
		return nil, err
	}

	hits, ok := result["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok || len(hits) == 0 {
		global.Logger.Warn("No refresh token found", zap.String("refreshToken", refreshToken))
		return nil, nil
	}

	source, ok := hits[0].(map[string]interface{})["_source"].(map[string]interface{})
	if !ok {
		global.Logger.Warn("Invalid source format in Elasticsearch response")
		return nil, nil
	}

	tokenData := &TokenData{}
	if val, ok := source["rf_public_key_refresh_token"].(string); ok {
		tokenData.RefreshTokenPublicKey = val
	} else {
		global.Logger.Warn("rf_public_key_refresh_token is missing or not a string")
		return nil, nil
	}
	if val, ok := source["rf_public_key_access_token"].(string); ok {
		tokenData.AccessTokenPublicKey = val
	} else {
		global.Logger.Warn("rf_public_key_access_token is missing or not a string")
		return nil, nil
	}

	return tokenData, nil
}

// VerifyToken verifies JWT token with public key
func VerifyToken(tokenString, publicKey string) (string, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		global.Logger.Error("Failed to parse RSA public key", zap.Error(err))
		return "", &UnauthorizedError{Message: UnauthorizedErrorMsg + ": invalid public key", Code: UnauthorizedErrorCode}
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return key, nil
	})

	if err != nil {
		global.Logger.Error("Failed to parse token", zap.Error(err))
		return "", &UnauthorizedError{Message: UnauthorizedErrorMsg + ": token parsing failed", Code: UnauthorizedErrorCode}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if id, ok := claims["_id"].(string); ok {
			return id, nil
		}
		global.Logger.Warn("Token claims missing _id or invalid")
	}

	return "", &UnauthorizedError{Message: UnauthorizedErrorMsg + ": invalid token claims", Code: UnauthorizedErrorCode}
}

// FindAccountByID searches for account by ID
func FindAccountByID(ctx context.Context, id string) (*model.Account, error) {
	if !indexExists(ctx, AccountIndex) {
		return nil, &UnauthorizedError{Message: UnauthorizedErrorMsg + ": accounts index not found", Code: UnauthorizedErrorCode}
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"_id": id,
			},
		},
	}

	queryBytes, err := json.Marshal(query)
	if err != nil {
		global.Logger.Error("Failed to marshal query", zap.Error(err))
		return nil, err
	}

	var result map[string]interface{}
	res, err := global.EsClient.Search(
		global.EsClient.Search.WithContext(ctx),
		global.EsClient.Search.WithIndex(AccountIndex),
		global.EsClient.Search.WithBody(bytes.NewReader(queryBytes)),
	)
	if err != nil {
		global.Logger.Error("Elasticsearch search error", zap.Error(err))
		return nil, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		global.Logger.Error("Failed to decode Elasticsearch response", zap.Error(err))
		return nil, err
	}

	hits, ok := result["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok || len(hits) == 0 {
		global.Logger.Warn("No account found", zap.String("id", id))
		return nil, &UnauthorizedError{Message: UnauthorizedErrorMsg + ": account not found", Code: UnauthorizedErrorCode}
	}

	hit, ok := hits[0].(map[string]interface{})
	if !ok {
		global.Logger.Warn("Invalid hit format in Elasticsearch response")
		return nil, &UnauthorizedError{Message: UnauthorizedErrorMsg + ": invalid account data", Code: UnauthorizedErrorCode}
	}

	source, ok := hit["_source"].(map[string]interface{})
	if !ok {
		global.Logger.Warn("Invalid source format in Elasticsearch response")
		return nil, &UnauthorizedError{Message: UnauthorizedErrorMsg + ": invalid account data", Code: UnauthorizedErrorCode}
	}

	// Lấy _id từ metadata của hit
	accountID, ok := hit["_id"].(string)
	if !ok || accountID == "" {
		global.Logger.Error("Account _id is missing or not a string in metadata", zap.Any("hit", hit))
		return nil, &UnauthorizedError{Message: UnauthorizedErrorMsg + ": invalid account _id", Code: UnauthorizedErrorCode}
	}

	// Kiểm tra các trường trong _source
	account := &model.Account{
		ID: accountID, // Sử dụng _id từ metadata
	}
	if val, ok := source["account_email"].(string); ok {
		account.Email = val
	} else {
		global.Logger.Warn("account_email is missing or not a string", zap.Any("source", source))
	}
	if val, ok := source["account_password"].(string); ok {
		account.Password = val
	} else {
		global.Logger.Warn("account_password is missing or not a string", zap.Any("source", source))
	}
	if val, ok := source["account_type"].(string); ok {
		account.Type = val
	} else {
		global.Logger.Warn("account_type is missing or not a string", zap.Any("source", source))
	}
	if val, ok := source["account_role"].(string); ok {
		account.Role = val
	} else {
		global.Logger.Warn("account_role is missing or not a string", zap.Any("source", source))
	}
	if val, ok := source["account_restaurant_id"].(string); ok {
		account.RestaurantID = val
	} else {
		global.Logger.Warn("account_restaurant_id is missing or not a string", zap.Any("source", source))
	}
	if val, ok := source["account_employee_id"].(string); ok {
		account.EmployeeID = val
	} else {
		global.Logger.Warn("account_employee_id is missing or not a string", zap.Any("source", source))
	}

	return account, nil
}

// FindRestaurantByID searches for restaurant by ID
func FindRestaurantByID(ctx context.Context, id string) (map[string]interface{}, error) {
	if !indexExists(ctx, RestaurantIndex) {
		return nil, &UnauthorizedError{Message: UnauthorizedErrorMsg + ": restaurants index not found", Code: UnauthorizedErrorCode}
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []interface{}{
					map[string]interface{}{
						"match": map[string]interface{}{
							"_id": id,
						},
					},
					map[string]interface{}{
						"match": map[string]interface{}{
							"isDeleted": false,
						},
					},
					map[string]interface{}{
						"match": map[string]interface{}{
							"restaurant_status": "inactive",
						},
					},
				},
			},
		},
	}

	queryBytes, err := json.Marshal(query)
	if err != nil {
		global.Logger.Error("Failed to marshal query", zap.Error(err))
		return nil, err
	}

	var result map[string]interface{}
	res, err := global.EsClient.Search(
		global.EsClient.Search.WithContext(ctx),
		global.EsClient.Search.WithIndex(RestaurantIndex),
		global.EsClient.Search.WithBody(bytes.NewReader(queryBytes)),
	)
	if err != nil {
		global.Logger.Error("Elasticsearch search error", zap.Error(err))
		return nil, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		global.Logger.Error("Failed to decode Elasticsearch response", zap.Error(err))
		return nil, err
	}

	hits, ok := result["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok || len(hits) == 0 {
		global.Logger.Warn("No restaurant found", zap.String("id", id))
		return nil, &UnauthorizedError{Message: UnauthorizedErrorMsg + ": restaurant not found", Code: UnauthorizedErrorCode}
	}

	source, ok := hits[0].(map[string]interface{})["_source"].(map[string]interface{})
	if !ok {
		global.Logger.Warn("Invalid source format in Elasticsearch response")
		return nil, &UnauthorizedError{Message: UnauthorizedErrorMsg + ": invalid restaurant data", Code: UnauthorizedErrorCode}
	}
	return source, nil
}

// FindEmployeeByID searches for employee by ID
func FindEmployeeByID(ctx context.Context, id string) (map[string]interface{}, error) {
	if !indexExists(ctx, EmployeeIndex) {
		return nil, &UnauthorizedError{Message: UnauthorizedErrorMsg + ": employees index not found", Code: UnauthorizedErrorCode}
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []interface{}{
					map[string]interface{}{
						"match": map[string]interface{}{
							"_id": id,
						},
					},
					map[string]interface{}{
						"match": map[string]interface{}{
							"isDeleted": false,
						},
					},
					map[string]interface{}{
						"match": map[string]interface{}{
							"epl_status": "enable",
						},
					},
				},
			},
		},
	}

	queryBytes, err := json.Marshal(query)
	if err != nil {
		global.Logger.Error("Failed to marshal query", zap.Error(err))
		return nil, err
	}

	var result map[string]interface{}
	res, err := global.EsClient.Search(
		global.EsClient.Search.WithContext(ctx),
		global.EsClient.Search.WithIndex(EmployeeIndex),
		global.EsClient.Search.WithBody(bytes.NewReader(queryBytes)),
	)
	if err != nil {
		global.Logger.Error("Elasticsearch search error", zap.Error(err))
		return nil, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		global.Logger.Error("Failed to decode Elasticsearch response", zap.Error(err))
		return nil, err
	}

	hits, ok := result["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok || len(hits) == 0 {
		global.Logger.Warn("No employee found", zap.String("id", id))
		return nil, &UnauthorizedError{Message: UnauthorizedErrorMsg + ": employee not found", Code: UnauthorizedErrorCode}
	}

	source, ok := hits[0].(map[string]interface{})["_source"].(map[string]interface{})
	if !ok {
		global.Logger.Warn("Invalid source format in Elasticsearch response")
		return nil, &UnauthorizedError{Message: UnauthorizedErrorMsg + ": invalid employee data", Code: UnauthorizedErrorCode}
	}
	return source, nil
}

// indexExists checks if an Elasticsearch index exists
func indexExists(ctx context.Context, index string) bool {
	res, err := global.EsClient.Indices.Exists([]string{index}, global.EsClient.Indices.Exists.WithContext(ctx))
	if err != nil {
		global.Logger.Error("Error checking index existence", zap.String("index", index), zap.Error(err))
		return false
	}
	defer res.Body.Close()
	return res.StatusCode == 200
}

// AuthGuard implements the authentication logic
func AuthGuard(ctx context.Context, headers map[string]string) (*model.Account, error) {
	accessTokenRTR := getTokenFromHeader(headers["x-at-rtr"])
	refreshTokenRTR := getTokenFromHeader(headers["x-rf-rtr"])
	accessTokenEPL := getTokenFromHeader(headers["x-at-epl"])
	refreshTokenEPL := getTokenFromHeader(headers["x-rf-epl"])

	accessToken := accessTokenRTR
	if accessToken == "" {
		accessToken = accessTokenEPL
	}
	refreshToken := refreshTokenRTR
	if refreshToken == "" {
		refreshToken = refreshTokenEPL
	}

	if accessToken == "" || refreshToken == "" {
		global.Logger.Warn("Missing tokens", zap.String("accessToken", accessToken), zap.String("refreshToken", refreshToken))
		return nil, &UnauthorizedError{Message: UnauthorizedErrorMsg + ": missing tokens", Code: UnauthorizedErrorCode}
	}

	tokenData, err := FindRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	if tokenData == nil || tokenData.RefreshTokenPublicKey == "" || tokenData.AccessTokenPublicKey == "" {
		global.Logger.Warn("Invalid token data", zap.Any("tokenData", tokenData))
		return nil, &UnauthorizedError{Message: UnauthorizedErrorMsg + ": invalid refresh token", Code: UnauthorizedErrorCode}
	}

	accessTokenID, err := VerifyToken(accessToken, tokenData.AccessTokenPublicKey)
	if err != nil {
		return nil, err
	}

	refreshTokenID, err := VerifyToken(refreshToken, tokenData.RefreshTokenPublicKey)
	if err != nil {
		return nil, err
	}

	if accessTokenID == "" || refreshTokenID == "" {
		global.Logger.Warn("Empty token IDs", zap.String("accessTokenID", accessTokenID), zap.String("refreshTokenID", refreshTokenID))
		return nil, &UnauthorizedError{Message: UnauthorizedErrorMsg + ": invalid token IDs", Code: UnauthorizedErrorCode}
	}

	account, err := FindAccountByID(ctx, accessTokenID)
	if err != nil {
		return nil, err
	}

	if account.Type == "restaurant" {
		_, err = FindRestaurantByID(ctx, account.RestaurantID)
		if err != nil {
			return nil, err
		}
	} else if account.Type == "employee" {
		_, err = FindEmployeeByID(ctx, account.EmployeeID)
		if err != nil {
			return nil, err
		}
	}

	account.Password = ""
	return account, nil
}

// AuthenMiddlewareAccount is a Gin middleware that authenticates the request and sets Account in context
func AuthenMiddlewareAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		headers := make(map[string]string)
		for k, v := range c.Request.Header {
			if len(v) > 0 {
				headers[strings.ToLower(k)] = v[0]
			}
		}

		account, err := AuthGuard(c.Request.Context(), headers)
		if err != nil {
			global.Logger.Error("Authentication failed", zap.Error(err), zap.Any("headers", headers))
			c.JSON(401, gin.H{
				"code":    UnauthorizedErrorCode,
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		// Gán Account vào gin.Context
		c.Set("account", account)
		global.Logger.Info("Authentication successful", zap.String("accountID", account.ID), zap.String("email", account.Email))

		// Tiếp tục xử lý request
		c.Next()
	}
}

// getTokenFromHeader extracts token from header
func getTokenFromHeader(header string) string {
	if header == "" {
		return ""
	}
	parts := strings.Split(header, " ")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}