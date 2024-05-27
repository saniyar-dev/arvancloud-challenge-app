package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBSource      string `mapstructure:"dbsource"`
	ServerAddress string `mapstructure:"serveraddress"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
