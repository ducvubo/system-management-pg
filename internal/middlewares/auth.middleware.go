package middlewares

import (
	"context"
	"log"

	"system-management-pg/internal/model"
	"system-management-pg/internal/service"
	"system-management-pg/internal/utils/auth"

	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the request url path
		uri := c.Request.URL.Path
		log.Println(" uri request: ", uri)
		// check headers authorization
		jwtToken, valid := auth.ExtractBearerToken(c)
		if !valid {
			c.AbortWithStatusJSON(401, gin.H{"code": 40001, "err": "Unauthorized", "description": ""})
			return
		}

		// validate jwt token by subject
		claims, err := auth.VerifyTokenSubject(jwtToken)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"code": 40001, "err": "invalid token", "description": ""})
			return
		}
		// update claims to context
		log.Println("claims::: UUID::", claims.Subject) // 11clitoken....
		ctx := context.WithValue(c.Request.Context(), "subjectUUID", claims.Subject)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func AuthenMiddlewareUserManagement() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uri := ctx.Request.URL.Path
		log.Println(" uri request: ", uri)
		accessToken, valid := auth.ExtractTokenFromKeyHeader(ctx, "x-at-tk")
		if !valid {
			ctx.AbortWithStatusJSON(401, gin.H{"code": 40001, "err": "Unauthorized1", "description": ""})
			return
		}
		refreshToken, valid := auth.ExtractTokenFromKeyHeader(ctx, "x-rf-tk")
		if !valid {
			ctx.AbortWithStatusJSON(401, gin.H{"code": 40001, "err": "Unauthorized2", "description": ""})
			return
		}

		clientId := ctx.GetHeader("id_user_guest")
		if clientId == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"code": 40001, "err": "invalid token3", "description": ""})
			return
		}

		userSession, err, statusCode := service.UserManagementAccount().FindUserSessionBySessionIdAndRefreshToken(ctx, clientId, refreshToken)
		if err != nil {
			ctx.AbortWithStatusJSON(statusCode, gin.H{"code": 40001, "err": "invalid token4", "description": ""})
			return
		}

		_, err = auth.VerifyToken(accessToken, userSession.UssKeyAt)
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"code": 40001, "err": "invalid token5", "description": ""})
			return
		}
		_, err = auth.VerifyToken(refreshToken, userSession.UssKeyRf)
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"code": 40001, "err": "invalid token6", "description": ""})
			return
		}

		userProfile, err, statusCode := service.UserManagementProfile().FindUserManagementProfile(ctx, userSession.UsaID)
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"code": 40001, "err": "invalid token7", "description": ""})
			return
		}

		userAccount, err, statusCode := service.UserManagementAccount().FindUserAccountById(ctx, userSession.UsaID)

		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"code": 40001, "err": "invalid token8", "description": ""})
			return
		}

		ctx.Set("userProfile", &model.UserManagementProfileOutput{
			UsID:       userProfile.UsID,
			UsName:     userProfile.UsName,
			UsAvatar:   userProfile.UsAvatar,
			UsGender:   userProfile.UsGender,
			UsPhone:    userProfile.UsPhone,
			UsAddress:  userProfile.UsAddress,
			UsBirthday: userProfile.UsBirthday,
			Isdeleted:  userProfile.Isdeleted,
			UsEmail:    userAccount.UsEmail,
		})
		ctx.Next()
	}
}
