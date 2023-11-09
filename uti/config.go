package uti

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB_DRIVER     string `mapstructure:"DB_DRIVER"`
	URI_DB        string `mapstructure:"URI_DB"`
	SERVER_SOURCE string `mapstructure:"SERVER_SOURCE"`
}

func LoadConfig(stringPath string) {
	viper.AddConfigPath(stringPath)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

}

func GetConfig() (config Config, err error) {
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
