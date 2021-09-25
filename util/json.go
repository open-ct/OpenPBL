package util

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
)

func StructToJson(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(data)
}

func JsonToStruct(data string, v interface{}) error {
	return json.Unmarshal([]byte(data), v)
}


func NewId() string {
	return uuid.NewV4().String()
}