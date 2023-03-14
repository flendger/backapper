package backupcontroller

import (
	"backapper/app/appservice"
	"log"
	"net/http"
)

type BackupController struct {
	service *appservice.AppService
}

func (c *BackupController) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	appName, exists := query["app"]
	if !exists {
		info := "Bad request: no app param"
		writeResponse(info, response)
		return
	}

	backUp, err := c.service.BackUp(appName[0])
	if err != nil {
		errInfo := "Back error: " + err.Error()
		writeResponse(errInfo, response)
		return
	}

	info := "OK backup: " + backUp
	writeResponse(info, response)
}

func writeResponse(info string, response http.ResponseWriter) {
	log.Println(info)
	_, err := response.Write([]byte(info))
	if err != nil {
		return
	}
}

func New(service *appservice.AppService) *BackupController {
	controller := &BackupController{service: service}

	http.Handle("/backup", controller)

	return controller
}
