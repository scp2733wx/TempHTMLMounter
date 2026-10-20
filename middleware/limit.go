package middleware

import (
	"fmt"
	"sync"
	"temphtmlmounter/pkg/errcode"
	"time"

	"github.com/gin-gonic/gin"
)

// rateBucket 简单的进程内固定窗口限流器，用于登录/注册等接口防爆破。
type rateBucket struct {
	mu     sync.Mutex
	counts map[string]int
	window time.Duration
	limit  int
	ts     time.Time
}

func newRateBucket(limit int, window time.Duration) *rateBucket {
	return &rateBucket{
		counts: make(map[string]int),
		window: window,
		limit:  limit,
		ts:     time.Now(),
	}
}

func (b *rateBucket) resetIfWindowElapsed() {
	if time.Since(b.ts) >= b.window {
		b.counts = make(map[string]int)
		b.ts = time.Now()
	}
}

func (b *rateBucket) allow(key string) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.resetIfWindowElapsed()
	b.counts[key]++
	return b.counts[key] <= b.limit
}

var authLimit = newRateBucket(10, time.Minute)

// AuthLimit 登录/注册接口限流，按客户端 IP 每窗口最多 limit 次。
func AuthLimit(c *gin.Context) {
	key := fmt.Sprintf("%s|%s", c.Request.URL.Path, c.ClientIP())
	if !authLimit.allow(key) {
		c.Error(errcode.ErrTooManyRequests)
		c.Abort()
		return
	}
	c.Next()
}