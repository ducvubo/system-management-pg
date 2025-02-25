package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func ExtractBearerToken(c *gin.Context) (string, bool) {

	// Authorization: Bearer token
	// token
	authHeader := c.GetHeader("Authorization")
	if strings.HasPrefix(authHeader, "Bearer") {
		return strings.TrimPrefix(authHeader, "Bearer "), true
	}
	return "", false
}

func ExtractTokenFromKeyHeader(c *gin.Context, key string) (string, bool) {

	// Authorization: Bearer token
	// token
	authHeader := c.GetHeader(key)
	if strings.HasPrefix(authHeader, "Bearer") {
		return strings.TrimPrefix(authHeader, "Bearer "), true
	}
	return "", false
}

func GetClientId(c *gin.Context) string {
	authHeader := c.Request.Header.Get("id_user_guest")
	return authHeader
}
