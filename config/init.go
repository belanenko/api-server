package config

import (
	"github.com/spf13/viper"
)

func Init() error {
	viper.AutomaticEnv()
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
