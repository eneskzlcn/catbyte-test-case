package config_test

import (
	"github.com/eneskzlcn/catbyte-test-task/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config, err := config.LoadConfig("../../.dev/", "local", "yaml")
	assert.Nil(t, err)
	assert.Equal(t, config.Server.Port, "3234")
	assert.Equal(t, config.Server.TrustedProxies[0], "*")
	assert.Equal(t, config.Redis.Address, "localhost:6379")
	assert.Equal(t, config.Redis.Password, "")
	assert.Equal(t, config.RabbitMQ.Address, "amqp://user:password@localhost:6001/")
	assert.Equal(t, config.RabbitMQ.Queue, "catbyte-messages")
}
