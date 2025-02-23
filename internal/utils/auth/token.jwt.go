package auth

import (
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type PayloadClaimsToken struct {
	jwt.StandardClaims
}

func GenerateToken(payload PayloadClaimsToken, expiration time.Duration, secretKey string) (string, error) {
	payload.ExpiresAt = time.Now().Add(expiration).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(tokenString string, secretKey string) (*PayloadClaimsToken, error) {
	token, err := jwt.ParseWithClaims(tokenString, &PayloadClaimsToken{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*PayloadClaimsToken); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

func GenerateCliTokenUUID(userId string) string {
	newUUID := uuid.New()
	uuidString := strings.ReplaceAll(newUUID.String(), "-", "")
	return userId + "clitoken" + uuidString
}

func ExtractUserIDFromCliToken(cliToken string) (string, error) {
	trimmed := strings.TrimPrefix(cliToken, "clitoken")

	for i := 0; i < len(trimmed); i++ {
		if !unicode.IsDigit(rune(trimmed[i])) {
			userIDStr := trimmed[:i]
			return userIDStr, nil
		}
	}

	return "", fmt.Errorf("invalid token format")
}
