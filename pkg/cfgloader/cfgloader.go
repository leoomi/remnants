package cfgloader

import (
	"github.com/spf13/viper"
)

func LoadConfig[T any](path string) (*T, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config T
	viper.Unmarshal(&config)
	return &config, nil
}
