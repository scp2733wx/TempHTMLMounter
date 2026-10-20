package redisdb

import (
	"context"
	"sync/atomic"
	"temphtmlmounter/config"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	rdb       *redis.Client
	ctx       = context.Background()
	Available atomic.Bool
)

const (
	tokenBlacklistPrefix = "token:blacklist:"
)

func InitRedis() {
	cfg := config.GetConfig().Redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: cfg.Password,
	})
	Refresh()
}

// Refresh 探测 Redis 可用性并更新 Available 状态，返回当前是否可用
func Refresh() bool {
	ok := rdb != nil && rdb.Ping(ctx).Err() == nil
	Available.Store(ok)
	return ok
}

// IsRevokedToken 检查 token 是否已列入黑名单（登出后）
func IsRevokedToken(token string) (bool, error) {
	if rdb == nil {
		return false, nil
	}
	n, err := rdb.Exists(ctx, tokenBlacklistPrefix+token).Result()
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

// RevokeToken 将 token 加入黑名单，expireSeconds 后自动过期
func RevokeToken(token string, expireSeconds int) error {
	if rdb == nil {
		return nil
	}
	return rdb.Set(ctx, tokenBlacklistPrefix+token, 1, time.Duration(expireSeconds)*time.Second).Err()
}