package initialize

import "github.com/gin-gonic/gin"

func Run() *gin.Engine {
	LoadConfig()
	InitLogger()
	InitMysqlC()
	InitRedis()
	InitKafka()
	RunHUb()
	InitSocketService()
	InitMessageService()
	r := InitRouter()
	return r
}
