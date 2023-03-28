package config

import (
	"encoding/json"
	"io"
	"os"
)

func Init() (*MainConfig, error) {
	jsonFile, err := os.Open("files/secrets/secrets.config.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	conf := new(MainConfig)
	json.Unmarshal(byteValue, conf)
	return conf, nil
}
