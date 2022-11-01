package config

type Config struct {
	App         App    `mapstructure:"app"`
	Environment string `mapstructure:"environment"`
	LogLevel    string `mapstructure:"log_level"`
}
type App struct {
	Port int `mapstructure:"port"`
}
