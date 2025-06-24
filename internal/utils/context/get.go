package context

import (
	"context"
	"errors"
	"log"

	"system-management-pg/internal/model"
	"system-management-pg/internal/utils/cache"

	"github.com/gin-gonic/gin"
)

type InfoUserUUID struct {
	UserId      uint64
	UserAccount string
}

func GetSubjectUUID(ctx context.Context) (string, error) {
	sUUID, ok := ctx.Value("subjectUUID").(string)
	if !ok {
		return "", errors.New("failed to get subject UUID")
	}
	return sUUID, nil
}

func GetUserIdFromUUID(ctx context.Context) (uint64, error) {
	sUUID, err := GetSubjectUUID(ctx)
	log.Println("sUUID::", sUUID)
	if err != nil {
		return 0, err
	}
	// get infoUser Redis from uuid
	var inforUser InfoUserUUID
	if err := cache.GetCache(ctx, sUUID, &inforUser); err != nil {
		log.Println("err:::", err)
		return 0, err
	}
	log.Println("inforUser:::", inforUser)
	return inforUser.UserId, nil
}

func GetAccoutFromCtx(c *gin.Context) *model.Account {
    account, exists := c.Get("account")
    if !exists {
        return nil
    }
    if acc, ok := account.(*model.Account); ok {
        return acc
    }
    return nil
}