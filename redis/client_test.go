package redis_test

import (
	"github.com/eneskzlcn/catbyte-test-task/redis"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClient(t *testing.T) {
	//redis image should be up and serve on 6379 to pass that test.
	client := redis.NewClient("localhost:6379", "")
	assert.NotNil(t, client)
}
func TestClient_SaveStruct(t *testing.T) {
	type mockType struct {
		Name    string   `json:"name"`
		Surname string   `json:"surname"`
		Hobbies []string `json:"hobbies"`
	}
	mockData := mockType{
		Name:    "enes",
		Surname: "kzlcn",
		Hobbies: []string{"football", "swimming", "listening music"},
	}
	client := redis.NewClient("localhost:6379", "")
	assert.NotNil(t, client)
	err := client.SaveStruct(mockData.Name, mockData)
	assert.Nil(t, err)
}
func TestClient_GetStruct(t *testing.T) {
	// post a new mockData to get
	type mockType struct {
		Name    string   `json:"name"`
		Surname string   `json:"surname"`
		Hobbies []string `json:"hobbies"`
	}
	mockData := mockType{
		Name:    "enes",
		Surname: "kzlcn",
		Hobbies: []string{"football", "swimming", "listening music"},
	}
	client := redis.NewClient("localhost:6379", "")
	assert.NotNil(t, client)
	err := client.SaveStruct(mockData.Name, mockData)
	assert.Nil(t, err)
	// test getting the data.
	mockDataResult := mockType{}
	err = client.GetStruct(mockData.Name, &mockDataResult)
	assert.Nil(t, err)
	assert.Equal(t, mockData.Name, mockDataResult.Name)
	assert.Equal(t, mockData.Surname, mockDataResult.Surname)
	assert.ElementsMatch(t, mockData.Hobbies, mockDataResult.Hobbies)
}
func TestClient_SaveToArrayL(t *testing.T) {
	// post a new mockData to get
	type mockType struct {
		Name    string   `json:"name"`
		Surname string   `json:"surname"`
		Hobbies []string `json:"hobbies"`
	}
	mockKey := "mock"
	mockDatas := []mockType{
		{
			Name:    "1",
			Surname: "k",
			Hobbies: []string{"basketball"},
		},
		{
			Name:    "2",
			Surname: "k",
			Hobbies: []string{"basketball"},
		},
		{
			Name:    "3",
			Surname: "k",
			Hobbies: []string{"basketball"},
		},
		{
			Name:    "4",
			Surname: "k",
			Hobbies: []string{"basketball"},
		},
	}
	client := redis.NewClient("localhost:6379", "")
	assert.NotNil(t, client)
	for _, mockData := range mockDatas {
		err := client.SaveToArrayL(mockKey, mockData)
		assert.Nil(t, err)
	}

}
func TestClient_GetArray(t *testing.T) {
	// post a new mockData to get
	type mockType struct {
		Name    string   `json:"name"`
		Surname string   `json:"surname"`
		Hobbies []string `json:"hobbies"`
	}
	mockKey := "mock"
	mockDatas := []mockType{
		{
			Name:    "1",
			Surname: "k",
			Hobbies: []string{"basketball"},
		},
		{
			Name:    "2",
			Surname: "k",
			Hobbies: []string{"basketball"},
		},
		{
			Name:    "3",
			Surname: "k",
			Hobbies: []string{"basketball"},
		},
		{
			Name:    "4",
			Surname: "k",
			Hobbies: []string{"basketball"},
		},
	}
	client := redis.NewClient("localhost:6379", "")
	assert.NotNil(t, client)
	//clear mock key
	err := client.Delete(mockKey)
	defer client.Delete(mockKey)
	assert.Nil(t, err)
	for _, mockData := range mockDatas {
		err := client.SaveToArrayL(mockKey, mockData)
		assert.Nil(t, err)
	}
	// try to get data into a type
	resultStr := make([]string, 0)
	err = client.GetArray(mockKey, &resultStr)
	assert.Nil(t, err)

	mockResult := make([]mockType, len(resultStr))
	for index, item := range resultStr {
		err = json.Unmarshal([]byte(item), &mockResult[index])
		assert.Nil(t, err)
	}
	assert.ElementsMatch(t, mockDatas, mockResult)
}
