package config

type Server struct {
	Port           string   `mapstructure:"port"`
	TrustedProxies []string `mapstructure:"trustedProxies"`
}
