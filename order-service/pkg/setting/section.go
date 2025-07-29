package setting

type Config struct {
	Server ServerSetting `mapsstructure:"server"`
	Logger LoggerSetting `mapsstructure:"logger"`
	Mysql  MySQlSetting  `mapsstructure:"mysql"`
	Redis  RedisSetting  `mapsstructure:"redis"`
	Kafka  KafkaSetting  `mapsstructure:"kafka"`
}

type ServerSetting struct {
	Port int    `mapsstructure:"port"`
	Mode string `mapsstructure:"mode"`
	Salt string `mapsstructure:"salt"`
}
type LoggerSetting struct {
	LogLevel    string `mapsstructure:"loglevel"`
	FileLogName string `mapsstructure:"filelogname"`
	MaxSize     int    `mapsstructure:"maxsize"`
	MaxBackups  int    `mapsstructure:"maxbackups"`
	MaxAge      int    `mapsstructure:"maxage"`
	ComPress    bool   `mapsstructure:"compress"`
}
type MySQlSetting struct {
	Host            string `mapsstructure:"host"`
	Port            int    `mapsstructure:"port"`
	Username        string `mapsstructure:"username"`
	Password        string `mapsstructure:"password"`
	Dbname          string `mapsstructure:"Dbname"`
	MaxIdleConns    int    `mapsstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapsstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapsstructure:"connMaxLifetime"`
}
type RedisSetting struct {
	Host     string `mapsstructure:"host"`
	Port     int    `mapsstructure:"port"`
	Password string `mapsstructure:"password"`
	Database int    `mapsstructure:"database"`
}
type KafkaSetting struct {
	Host     string `mapsstructure:"host"`
	Port     int    `mapsstructure:"port"`
	Username string `mapsstructure:"username"`
	Password string `mapsstructure:"password"`
}
