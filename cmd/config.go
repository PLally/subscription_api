package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	Database databaseConfig
	HttpPort string
}

type databaseConfig struct {
	Address string
	Port string
	User string
	Password string
	DatabaseName string
}

func readConfig(filename string) (*config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := config{}

	err = decoder.Decode(&conf)
	fmt.Println(conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}