package redis

import (
	"encoding/json"
	"github.com/go-redis/redis"
)

type Client struct {
	client *redis.Client
}

func NewClient(address, password string) *Client {

	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})
	if _, err := client.Ping().Result(); err != nil {
		return nil
	}
	return &Client{client: client}
}
func (c *Client) SaveStruct(key string, value interface{}) error {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	if err = c.client.Set(key, jsonValue, 0).Err(); err != nil {
		return err
	}
	return nil
}
func (c *Client) GetStruct(key string, out interface{}) error {
	val, err := c.client.Get(key).Bytes()
	if err != nil {
		return err
	}
	if err = json.Unmarshal(val, &out); err != nil {
		return err
	}
	return nil
}
