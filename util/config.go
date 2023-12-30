package util

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DbDriver            string        `mapstructure:"DB_DRIVER"`
	DbSource            string        `mapstructure:"DB_SOURCE"`
	ServerPort          int           `mapstructure:"SERVER_PORT"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (Config, error) {
	config := Config{}

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		return config, err
	}

	err := viper.Unmarshal(&config)
	fmt.Printf("config :: %+v\n", config)
	return config, err
}
