package main

import (
	"backapper/app/appreader"
	"backapper/app/appservice"
	"backapper/app/backupcontroller"
	"backapper/app/deploycontroller"
	"backapper/config"
	"log"
	"net/http"
)

const configPath = "backapper.cfg"

func main() {
	configuration := config.Load(configPath)

	appHolder := appreader.Read(configuration.AppConfigPath)
	service := appservice.New(appHolder)

	backupcontroller.New(service)
	deploycontroller.New(service)

	err := http.ListenAndServe(":"+configuration.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
