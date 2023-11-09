package uti

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DbDriver     string `mapstructure:"DB_DRIVER"`
	Uri_db       string `mapstructure:"URI_DB"`
	ServerSource string `mapstructure:"SERVER_SOURCE"`
}

func LoadConfig(stringPath string) {
	viper.AddConfigPath(stringPath)
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	fmt.Println(viper.)
	viper.AutomaticEnv()
}

func GetConfig() (config Config, err error) {
	err = viper.Unmarshal(&config)
	return
}
