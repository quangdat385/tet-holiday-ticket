package initialize

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/quangdat385/holiday-ticket/communications-service/global"
	"go.uber.org/zap"
)

func checkErrorPanicC(err error, errMsg string) {
	if err != nil {
		global.Logger.Error(errMsg, zap.Error(err))
		panic(err)
	}
}

func InitMysqlC() {
	m := global.Config.Mysql
	fmt.Println("Initializing MySQL...", m.Username, m.Password)
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := sql.Open("mysql", s)
	checkErrorPanicC(err, "InitMysql initialization error")
	global.Logger.Info("InitMysql initialized successfully")
	global.Mdb = db
	SetPoolC()
}
func SetPoolC() {
	m := global.Config.Mysql
	global.Mdb.SetMaxIdleConns(m.MaxIdleConns)
	global.Mdb.SetMaxOpenConns(m.MaxOpenConns)
	global.Mdb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime) * time.Second)
}
