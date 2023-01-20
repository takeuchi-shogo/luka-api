package lib

import (
	"log"

	"github.com/spf13/viper"
)

// As a bridge into /src/infrastructure/config
type Env struct {
	ServerPort  string `mapstructure:"SERVER_PORT"`
	Environment string `mapstructure:"ENV"`
	LogOutput   string `mapstructure:"LOG_OUTPUT"`
	LogLevel    string `mapstructure:"LOG_LEVEL"`
	DBUsername  string `mapstructure:"DB_USER"`
	DBPassword  string `mapstructure:"DB_PASS"`
	DBHost      string `mapstructure:"DB_HOST"`
	DBPort      string `mapstructure:"DB_PORT"`
	DBName      string `mapstructure:"DB_NAME"`
}

func NewEnv() Env {

	env := Env{}
	viper.SetConfigFile("./.env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("error read config", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("env file unmarshal: ", err)
	}

	return env
}
