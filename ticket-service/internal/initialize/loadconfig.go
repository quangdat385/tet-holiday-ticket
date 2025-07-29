package initialize

import (
	"fmt"

	"github.com/quangdat385/holiday-ticket/ticket-service/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	// read config
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("failed to read config: %s", err))
	}
	//read server
	fmt.Println("Server Port::", viper.GetInt("server.port"))
	if err = viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration : %v \n", err)
	}
	fmt.Println("Server Port::", global.Config.Server.Port)
}
