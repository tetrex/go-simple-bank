package util

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DbDriver   string `mapstructure:"DB_DRIVER"`
	DbSource   string `mapstructure:"DB_SOURCE"`
	ServerPort int    `mapstructure:"SERVER_PORT"`
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
	fmt.Printf("config :: %+v", config)
	return config, err
}
