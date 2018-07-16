package tests

import (
	"os"
	"testing"
	"wechat-miniprogram/application"
	"wechat-miniprogram/utils/database"
	"wechat-miniprogram/utils/server"
)

var app application.App

func TestMain(m *testing.M) {
	app = application.App{}

	serverConfig, err := server.ReadConfig("../serverConfig.json")
	dbConfig, err := database.ReadConfig("../testDBConfig.json")
	if err != nil {
		os.Exit(1)
	}

	app.InitApp(*dbConfig, *serverConfig)
	app.Run()
	code := m.Run()
	os.Exit(code)
}
