package config

import (
	"github.com/spf13/viper"
)

func Init() (*Config, error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
