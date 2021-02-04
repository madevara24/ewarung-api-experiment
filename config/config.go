package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBDatabase string `mapstructure:"DB_DATABASE"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	JWTSecret  string `mapstructure:"JWT_SECRET"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Configuration file not found.")
	}

	err = viper.Unmarshal(&config)
	return
}
