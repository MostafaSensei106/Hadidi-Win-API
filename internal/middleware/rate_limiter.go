package middleware

import (
	"net/http"
	"sync"

	"github.com/MostafaSensei106/Hadidi-Win-API/internal/delivery"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var vistors = make(map[string]*rate.Limiter)

var mu sync.Mutex

func getVistor(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, isHere := vistors[ip]
	if !isHere {
		limiter = rate.NewLimiter(5, 10)
		vistors[ip] = limiter
	}
	return limiter
}

func RateLimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limiter := getVistor(ctx.ClientIP())
		if !limiter.Allow() {
			delivery.Fail(ctx, http.StatusTooManyRequests, "Too many requests", "Please try agin later")
			return
		}
		ctx.Next()
	}
}
