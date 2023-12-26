package config

import (
	"github.com/spf13/viper"
)

func Init() {
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
}
