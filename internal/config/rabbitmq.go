package config

type RabbitMQ struct {
	Address string `mapstructure:"address"`
	Queue   string `mapstructure:"queue"`
}
