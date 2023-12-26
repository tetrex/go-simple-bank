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

func LoadConfig() (Config, error) {
	config := Config{}

	viper.AddConfigPath("/")
	viper.SetConfigFile("app.env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		return config, err
	}

	err := viper.Unmarshal(&config)
	fmt.Printf("config :: %+v", config)
	return config, err
}
