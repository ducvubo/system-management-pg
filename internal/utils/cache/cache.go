package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"system-management-pg/global"

	"github.com/redis/go-redis/v9"
)

func SetCache(ctx context.Context, key string, value interface{}) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %v", err)
	}

	err = global.Rdb.Set(ctx, key, jsonData, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to set cache: %v", err)
	}
	return nil
}

func SetCacheWithExpiration(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %v", err)
	}

	err = global.Rdb.Set(ctx, key, jsonData, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set cache with expiration: %v", err)
	}
	return nil
}

func DeleteCache(ctx context.Context, key string) error {
	err := global.Rdb.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete cache: %v", err)
	}
	return nil
}

func GetCache(ctx context.Context, key string, obj interface{}) error {
	rs, err := global.Rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return fmt.Errorf("key %s not found", key)
	} else if err != nil {
		return err
	}
	// convert rs json to object
	if err := json.Unmarshal([]byte(rs), obj); err != nil {
		return fmt.Errorf("failed to unmarshal")
	}
	return nil
}
