package main

import (
	"backapper/app/appreader"
	"backapper/app/appservice"
	"backapper/app/backupcontroller"
	"log"
	"net/http"
)

func main() {
	appHolder := appreader.Read("apps.json")
	service := appservice.New(appHolder)

	backUpController := backupcontroller.New(service)

	http.Handle("/backup", backUpController)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
