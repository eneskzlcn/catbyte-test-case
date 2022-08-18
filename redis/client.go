package redis

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"log"
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
		log.Println("marshalling error")
		return err
	}
	if err = c.client.Set(key, jsonValue, 0).Err(); err != nil {
		log.Println("redis set error")
		return err
	}
	return nil
}

//SaveToArrayL adds new elements to the beginning of the array as the task wants to.

func (c *Client) SaveToArrayL(key string, value interface{}) error {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		log.Println("marshalling error")
		return err
	}
	if err = c.client.LPush(key, jsonValue).Err(); err != nil {
		log.Println("redis set error")
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
func (c *Client) GetArray(key string, out *[]string) error {
	val, err := c.client.LRange(key, 0, -1).Result()
	if err != nil {
		return err
	}
	*out = val
	for _, item := range val {
		println(item)
	}
	return nil
}
func (c *Client) Delete(key string) error {
	if err := c.client.Del(key).Err(); err != nil {
		return err
	}
	return nil
}
