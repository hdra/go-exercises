package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config map[string]interface{}

func (c Config) Name() string {
	return c["name"].(string)
}

func LoadConfig(path string) (Config, error) {
	var m map[string]interface{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return m, err
	}

	err = json.Unmarshal(data, &m)
	return m, err
}

func main() {
	json, err := LoadConfig("config.json")

	if err != nil {
		panic(err)
	}

	fmt.Println(json)
	fmt.Println(json.Name())
}
