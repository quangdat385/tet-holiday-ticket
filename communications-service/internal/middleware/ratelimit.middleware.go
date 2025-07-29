package middleware

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/communications-service/global"
	"github.com/quangdat385/holiday-ticket/communications-service/response"
	"github.com/quangdat385/holiday-ticket/communications-service/utils/context"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/redis"
)

type RateLimit struct {
	globalkRateLimit    *limiter.Limiter
	publicAPIRateLimit  *limiter.Limiter
	privateAPIRateLimit *limiter.Limiter
}

func NewRateLimit() *RateLimit {
	rateLimit := &RateLimit{
		globalkRateLimit:    rateLimit("100-S"),
		publicAPIRateLimit:  rateLimit("80-S"),
		privateAPIRateLimit: rateLimit("60-S"),
	}

	return rateLimit
}

func rateLimit(interval string) *limiter.Limiter {
	store, err := redis.NewStoreWithOptions(global.Rdb, limiter.StoreOptions{
		Prefix:          "rate-limiter",
		MaxRetry:        3,
		CleanUpInterval: time.Hour,
	})
	if err != nil {
		return nil
	}
	rate, err := limiter.NewRateFromFormatted(interval)
	if err != nil {
		panic(err)
	}
	instance := limiter.New(store, rate)
	return instance
}

func (rl *RateLimit) GlobalRalimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := "global"
		log.Println("Global Rate Limiter")
		limiterContext, err := rl.globalkRateLimit.Get(c, key)
		if err != nil {
			fmt.Println("failed to get limiter Global", err)
			c.Next()
			return
		}
		if limiterContext.Reached {
			log.Printf("Global Rate Limiter reached:%s", key)
			c.JSON(429, gin.H{"error": "Too many requests. Please try again later."})
			return
		}

		c.Next()
	}
}
func (rl *RateLimit) PublicRalimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		urlPath := c.Request.URL.Path
		ratelimitPath := rl.publicAPIRateLimit
		log.Println("Client IP--->", c.ClientIP())
		key := fmt.Sprintf("%s:%s", c.ClientIP(), urlPath)
		limiterContext, err := ratelimitPath.Get(c, key)
		if err != nil {
			fmt.Println("failed to check rate limit", err)
			c.Next()
			return
		}
		if limiterContext.Reached {
			log.Printf("Rate Limiter reached:%s", key)
			c.JSON(429, gin.H{"error": "Too many requests. Please try again later."})
			return
		}

		c.Next()
	}
}

func (rl *RateLimit) PrivateRalimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		urlPath := c.Request.URL.Path
		ratelimitPath := rl.privateAPIRateLimit
		userRoles, err := context.GetUserBase(c.Request.Context())
		fmt.Println("userRoles", userRoles)
		if err != nil {
			response.ErrorResponse(c, response.UnauthorizedCodeStatus, err.Error())
			c.Abort()
			return
		}
		key := fmt.Sprintf("%s:%s", userRoles.UserID, urlPath)
		limiterContext, err := ratelimitPath.Get(c, key)
		if err != nil {
			fmt.Println("failed to check rate limit", err)
			c.Next()
			return
		}
		if limiterContext.Reached {
			log.Printf("Rate Limiter reached:%s", key)
			c.JSON(429, gin.H{"error": "Too many requests. Please try again later."})
			return
		}

		c.Next()
	}
}
