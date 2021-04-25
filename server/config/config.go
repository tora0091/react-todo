package config

import (
	"log"

	"github.com/spf13/viper"
)

func Configure(s string) interface{} {
	viper.SetConfigFile("config.yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	return viper.Get(s)
}

func ConfigureString(s string) string {
	return Configure(s).(string)
}
