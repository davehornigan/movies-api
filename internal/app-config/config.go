package app_config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	ServerPort      string `mapstructure:"SERVER_PORT"`
	Environment     string `mapstructure:"APP_ENV"`
	TmdbServer      string `mapstructure:"TMDB_API_SERVER"`
	TmdbAccessToken string `mapstructure:"TMDB_API_ACCESS_TOKEN"`
}

func LoadConfig(configFile string) (config Config, err error) {
	SetDefaults()
	viper.SetConfigFile(configFile)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		log.Fatalf("read app-config error: %s", err.Error())
	}

	err = viper.Unmarshal(&config)
	return
}

func SetDefaults() {
	viper.SetDefault("APP_ENV", "development")
}
