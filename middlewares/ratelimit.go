package middlewares

import (
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type RateLimiter struct {
	tokens      int
	capacity    int
	lastRefill  time.Time
	refillRate  int // tokens per second
	mu          sync.Mutex
}

var clients = sync.Map{} // map[ip]*RateLimiter

func getLimiter(ip string) *RateLimiter {
	limiterIface, exists := clients.Load(ip)
	if !exists {
		limiter := &RateLimiter{
			tokens:     1,   // allow 5 requests
			capacity:   2,
			refillRate: 1,   // refill 1 token/sec
			lastRefill: time.Now(),
		}
		clients.Store(ip, limiter)
		return limiter
	}
	return limiterIface.(*RateLimiter)
}

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		log.Println("[Middlewares(ratelimit.go)] Rate Limit Middleware")
		ip := c.ClientIP()
		limiter := getLimiter(ip)
		limiter.mu.Lock()
		defer limiter.mu.Unlock()

		// Refill
		now := time.Now()
		elapsed := now.Sub(limiter.lastRefill).Seconds()
		refill := int(elapsed * float64(limiter.refillRate))
		if refill > 0 {
			limiter.tokens = min(limiter.capacity, limiter.tokens+refill)
			limiter.lastRefill = now
		}

		// Check
		if limiter.tokens <= 0 {
			c.JSON(429, gin.H{"error": "Too Many Requests"})
			c.Abort()
			return
		}

		limiter.tokens--

		c.Next()
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
