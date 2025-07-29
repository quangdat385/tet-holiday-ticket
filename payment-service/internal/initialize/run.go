package initialize

import "github.com/gin-gonic/gin"

func Run() *gin.Engine {
	LoadConfig()
	InitMysqlC()
	InitRedis()
	InitKafka()
	r := InitRouter()
	return r
}
