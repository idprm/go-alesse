package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func ViperEnv(key string) string {

	viper.SetConfigName("app")

	viper.SetConfigType("env")

	viper.AddConfigPath(".")
	viper.AddConfigPath("/root/app/go-alesse/")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.UnsupportedConfigError); ok {
			log.Printf("[ERROR] No Rigfile exists.")
			os.Exit(1)
		} else {
			log.Printf("[ERROR] %s", err)
		}
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}
