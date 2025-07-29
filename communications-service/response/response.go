package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/communications-service/global"
	"go.uber.org/zap"
)

type ResponseData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
type ErrorResponseData struct {
	Code   int    `json:"code"`
	Err    string `json:"error"`
	Detail any    `json:"detail"`
}

func SuccessResponse(c *gin.Context, code int, data any) {
	method, _ := c.Get("method")
	path, _ := c.Get("path")
	request_id, _ := c.Get("request_id")
	before, _ := c.Get("before")
	span, _ := c.Get("span")
	ip, _ := c.Get("ip")
	device_id, _ := c.Get("device_id")
	client_id, _ := c.Get("client_id")
	duration := time.Since(before.(time.Time))
	global.Logger.Info(msg[code],
		zap.Any("request_id", request_id),
		zap.String("span", span.(string)),
		zap.String("span_parent", ip.(string)),
		zap.String("ip", ip.(string)),
		zap.String("device_id", device_id.(string)),
		zap.String("client_id", client_id.(string)),
		zap.String("method", method.(string)),
		zap.String("path", path.(string)),
		zap.Duration("duration", time.Duration(duration.Nanoseconds())),
		zap.Int64("timestamp", before.(time.Time).UnixMilli()),
		zap.Int("code", code),
		zap.Any("data", data),
		zap.String("message", msg[code]),
	)
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}
func ErrorResponse(c *gin.Context, code int, detail any) {
	method, _ := c.Get("method")
	path, _ := c.Get("path")
	request_id, _ := c.Get("request_id")
	before, _ := c.Get("before")
	span, _ := c.Get("span")
	ip, _ := c.Get("ip")
	device_id, _ := c.Get("device_id")
	duration := time.Since(before.(time.Time))
	global.Logger.Error(msg[code],
		zap.Any("request_id", request_id),
		zap.String("span", span.(string)),
		zap.String("span_parent", ip.(string)),
		zap.String("ip", ip.(string)),
		zap.String("device_id", device_id.(string)),
		zap.String("client_id", c.GetString("client_id")),
		zap.String("method", method.(string)),
		zap.String("path", path.(string)),
		zap.Duration("duration", time.Duration(duration.Nanoseconds())),
		zap.Int64("timestamp", before.(time.Time).UnixMilli()),
		zap.Int("code", code),
		zap.Any("detail", detail),
		zap.String("err", msg[code]),
	)
	c.JSON(http.StatusBadRequest, ErrorResponseData{
		Code:   code,
		Err:    msg[code],
		Detail: detail,
	})
}
