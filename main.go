package main

import (
	"backapper/app/appreader"
	"backapper/app/appservice"
	"backapper/app/backupcontroller"
	"backapper/app/deploycontroller"
	"log"
	"net/http"
)

func main() {
	appHolder := appreader.Read("apps.json")
	service := appservice.New(appHolder)

	backupcontroller.New(service)
	deploycontroller.New(service)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
