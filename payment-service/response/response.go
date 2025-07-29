package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}
func ErrorResponse(c *gin.Context, code int, detail interface{}) {
	c.JSON(http.StatusBadRequest, ErrorResponseData{
		Code:   code,
		Err:    msg[code],
		Detail: detail,
	})
}
