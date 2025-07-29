package middleware

import (
	"time"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/quangdat385/holiday-ticket/order-service/utils/random"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		request_id := c.GetHeader("x-request-id")
		if request_id == "" {
			request_id = uuid.New().String()
		}
		span := c.GetHeader("x-span")
		if span == "" {
			span = "0"
		}
		device_id := c.GetHeader("x-device-id")
		if device_id == "" {
			device_id = strconv.Itoa(random.GenerateSixDigits())
		}
		client_id := c.GetHeader("x-client-id")
		if client_id == "" {
			client_id = "0"
		}
		ip := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path
		// before request
		c.Set("request_id", request_id)
		c.Set("before", t)
		c.Set("span", span)
		c.Set("parent_span", span)
		c.Set("ip", ip)
		c.Set("device_id", device_id)
		c.Set("client_id", client_id)
		c.Set("method", method)
		c.Set("path", path)
		c.Next()
	}
}
