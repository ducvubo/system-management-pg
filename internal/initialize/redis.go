// package initialize

// import (
// 	"context"
// 	"fmt"

// 	"system-management-pg/global"

// 	"github.com/redis/go-redis/v9"
// 	"go.uber.org/zap"
// )

// var ctx = context.Background()

// func InitRedis() {
// 	r := global.Config.Redis
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
// 		Password: r.Password, // no password set
// 		DB:       r.Database, // use default DB
// 		Username : r.User,
// 		PoolSize: 1000,         //
// 	})

// 	_, err := rdb.Ping(ctx).Result()
// 	if err != nil {
// 		global.Logger.Error("Redis initialization Error:", zap.Error(err))
// 	}

// 	// fmt.Println("Initializing Redis Successfully")
// 	global.Logger.Info("Initializing Redis Successfully")
// 	global.Rdb = rdb
// 	// redisExample()
// }

// func redisExample() {
// 	err := global.Rdb.Set(ctx, "score", 100, 0).Err()
// 	if err != nil {
// 		fmt.Println("Error redis setting:", zap.Error(err))
// 		return
// 	}

// 	value, err := global.Rdb.Get(ctx, "score").Result()
// 	if err != nil {
// 		fmt.Println("Error redis setting:", zap.Error(err))
// 		return
// 	}

// 	global.Logger.Info("value score is::", zap.String("score", value))
// }

package initialize

import (
	"context"
	"fmt"
	"sync"

	"system-management-pg/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// RedisSingleton quản lý kết nối Redis theo mô hình Singleton
type RedisSingleton struct {
	client *redis.Client
}

var (
	redisInstance *RedisSingleton
	redisOnce     sync.Once
	ctx           = context.Background()
)

// checkErrorPanicRedis ghi log và panic nếu có lỗi
func checkErrorPanicRedis(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

// GetRedisInstance trả về instance duy nhất của Redis client
func GetRedisInstance() *redis.Client {
	redisOnce.Do(func() {
		redisInstance = &RedisSingleton{}
		redisInstance.initRedis()
	})
	return redisInstance.client
}

// initRedis khởi tạo kết nối Redis
func (r *RedisSingleton) initRedis() {
	redisConfig := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		Username: redisConfig.User,
		DB:       redisConfig.Database,
		PoolSize: 1000,
	})

	// Kiểm tra kết nối
	_, err := rdb.Ping(ctx).Result()
	checkErrorPanicRedis(err, "Redis initialization error")

	r.client = rdb
	global.Rdb = rdb
	global.Logger.Info("Initialized Redis successfully with Singleton pattern")
}

// Close đóng kết nối Redis
func (r *RedisSingleton) Close() error {
	if r.client != nil {
		return r.client.Close()
	}
	return nil
}

// InitRedis khởi tạo Redis với Singleton pattern
func InitRedis() {
	global.Rdb = GetRedisInstance()
}

// RedisExample ví dụ sử dụng Redis
func RedisExample() {
	err := global.Rdb.Set(ctx, "score", 100, 0).Err()
	if err != nil {
		global.Logger.Error("Redis set error", zap.Error(err))
		return
	}

	value, err := global.Rdb.Get(ctx, "score").Result()
	if err != nil {
		global.Logger.Error("Redis get error", zap.Error(err))
		return
	}

	global.Logger.Info("Value score is", zap.String("score", value))
}