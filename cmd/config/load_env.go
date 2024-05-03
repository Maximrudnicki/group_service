package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	MONGODB_HOST string `mapstructure:"MONGODB_HOST"`
	MONGODB_DB   string `mapstructure:"MONGODB_DB"`
	MONGODB_PORT string `mapstructure:"MONGODB_PORT"`
	GRPCPort     string `mapstructure:"GRPC_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("group")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
