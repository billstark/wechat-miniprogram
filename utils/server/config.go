package server

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type ServerConfig struct {
	API    apiConfig    `json:"api"`
	Server serverConfig `json:"server"`
}

type apiConfig struct {
	ID          string `json:"id"`
	Secret      string `json:"secret"`
	EnableHTTPS bool   `json:"enableHTTPS"`
	Domain      string `json:"domain"`
}

type serverConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// ListenAddress returns the address and port the application should listen on.
func (c *serverConfig) ListenAddress() string {
	return c.Host + ":" + strconv.Itoa(c.Port)
}

// ReadConfig loads the application's configuration from a JSON file at the given path.
func ReadConfig(path string) (*ServerConfig, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config ServerConfig
	err = json.Unmarshal(raw, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
