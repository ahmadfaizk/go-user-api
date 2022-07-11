package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort       int    `mapstructure:"SERVER_PORT"`
	JWTAccessSecret  string `mapstructure:"JWT_ACCESS_SECRET"`
	JWTRefreshSecret string `mapstructure:"JWT_REFRESH_SECRET"`
	DatabaseHost     string `mapstructure:"DB_HOST"`
	DatabasePort     int    `mapstructure:"DB_PORT"`
	DatabaseName     string `mapstructure:"DB_NAME"`
	DatabaseUser     string `mapstructure:"DB_USER"`
	DatabasePassword string `mapstructure:"DB_PASSWORD"`
}

func Load(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
