package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type AllConfig struct {
	Server Server
	Redis Redis 
}

type Server struct {
	Port  string
	Level string
}
type Redis struct {
	Host string 
	Port string 
	Password string 
	DataBase int `mapstructure:"data_base"` 
}

func ConfigInit() *AllConfig {
	config := viper.New() 
	config.AddConfigPath("./config")
	config.SetConfigName("application-dev")
	config.SetConfigType("yaml")

	var configs *AllConfig
	
	err := config.ReadInConfig() 
	if err != nil {
		panic(fmt.Errorf("ReadInConfig() error: %s", err))
	}
	err = config.Unmarshal(&configs)
	if err != nil {
		panic(fmt.Errorf("Unmarshal() error: %s", err))
	}
	fmt.Println("configuration: ", configs)
	return configs 
}