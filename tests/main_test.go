package tests

import (
	"fmt"
	"os"
	"testing"
	"wechat-miniprogram/application"
	"wechat-miniprogram/utils/database"
	"wechat-miniprogram/utils/server"
)

var app application.App

func TestMain(m *testing.M) {
	app = application.App{}

	serverConfig, err := server.ReadConfig("../config/serverConfig.json")
	dbConfig, err := database.ReadConfig("../config/testDBConfig.json")
	if err != nil {
		os.Exit(1)
	}

	err = app.InitApp(*dbConfig, *serverConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	app.Run()
	fmt.Println(app.DB)
	code := m.Run()
	os.Exit(code)
}
