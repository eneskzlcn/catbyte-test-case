package config

type Redis struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
}
