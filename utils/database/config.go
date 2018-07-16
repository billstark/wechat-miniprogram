package database

import (
	"encoding/json"
	"io/ioutil"
)

// DBConfig defines basic db config model
type DBConfig struct {
	Driver             string `json:"driver"`
	Host               string `json:"host"`
	Port               int    `json:"port"`
	EnableSSL          bool   `json:"enable_ssl"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	DBName             string `json:"db_name"`
	MaxIdleConnections int    `json:"max_idle_connections"`
	MaxOpenConnections int    `json:"max_open_connections"`
}

// ReadConfig defines method to read config file
func ReadConfig(path string) (*DBConfig, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config DBConfig
	err = json.Unmarshal(raw, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
