package config

import "github.com/spf13/viper"

type Config struct {
	Server   Server   `mapstructure:"server"`
	RabbitMQ RabbitMQ `mapstructure:"rabbitMQ"`
	Redis    Redis    `mapstructure:"redis"`
}

func LoadConfig(path, name, configType string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(configType)

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
