package main

import (
	"fmt"
	"os"

	"wechat-miniprogram/application"
	"wechat-miniprogram/utils/database"
	"wechat-miniprogram/utils/server"
)

func main() {
	app := application.App{}

	serverConfig, err := server.ReadConfig("./config/serverConfig.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	dbConfig, err := database.ReadConfig("./config/dbConfig.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = app.InitApp(*dbConfig, *serverConfig)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	app.Run()
	app.Logger.Log(
		application.LogLayerTag, application.LayerApplication,
		application.LogMessageTag, application.MessageHalting,
		application.LogErrorTag, <-app.Errs)
}
