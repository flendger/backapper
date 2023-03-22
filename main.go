package main

import (
	"backapper/app/appreader"
	"backapper/app/appservice"
	"backapper/backupcontroller"
	"backapper/config"
	"backapper/deploycontroller"
	"backapper/restartcontroller"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

const configPath = "backapper.cfg"

var appLogger *log.Logger

func init() {
	file, err := os.OpenFile("backapper.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		appLogger = log.Default()
		return
	}

	appLogger = log.New(io.MultiWriter(file, os.Stdout), "Backapper: ", log.Ldate|log.Ltime)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = appLogger.Writer()

	configuration := config.Load(configPath, appLogger)

	appHolder := appreader.Read(configuration.AppConfigPath, appLogger)
	service := appservice.New(appHolder, appLogger)

	engine := gin.Default()

	engine.GET("/backup", backupcontroller.New(service).Handle)
	engine.POST("/deploy", deploycontroller.New(service).Handle)
	engine.GET("/restart", restartcontroller.New(service).Handle)

	err := engine.Run(":" + configuration.Port)
	if err != nil {
		appLogger.Fatal(err)
	}
}
