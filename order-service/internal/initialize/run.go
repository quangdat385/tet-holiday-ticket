package initialize

import "github.com/gin-gonic/gin"

func Run() *gin.Engine {
	LoadConfig()
	InitLogger()
	InitMysqlC()
	InitRedis()
	InitKafka()
	InitService()
	r := InitRouter()
	return r
}
