package main

import (
	"backapper/app/appreader"
	"backapper/app/appservice"
	"backapper/backupcontroller"
	"backapper/config"
	"backapper/deploycontroller"
	"backapper/restartcontroller"
	"io"
	"log"
	"net/http"
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
	configuration := config.Load(configPath, appLogger)

	appHolder := appreader.Read(configuration.AppConfigPath, appLogger)
	service := appservice.New(appHolder, appLogger)

	backupcontroller.New(service, appLogger)
	deploycontroller.New(service, appLogger)
	restartcontroller.New(service, appLogger)

	err := http.ListenAndServe(":"+configuration.Port, nil)
	if err != nil {
		appLogger.Fatal(err)
	}
}
