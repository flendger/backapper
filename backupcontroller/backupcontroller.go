package backupcontroller

import (
	"backapper/app/appservice"
	"log"
	"net/http"
)

type BackupController struct {
	service *appservice.AppService
	logger  *log.Logger
}

func (c *BackupController) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	appName, exists := query["app"]
	if !exists {
		info := "Bad request: no app param"
		c.writeResponse(info, response)
		return
	}

	backUp, err := c.service.BackUp(appName[0])
	if err != nil {
		errInfo := "Back error: " + err.Error()
		c.writeResponse(errInfo, response)
		return
	}

	info := "OK backup: " + backUp
	c.writeResponse(info, response)
}

func (c *BackupController) writeResponse(info string, response http.ResponseWriter) {
	c.logger.Println(info)
	_, err := response.Write([]byte(info))
	if err != nil {
		return
	}
}

func New(service *appservice.AppService, logger *log.Logger) *BackupController {
	controller := &BackupController{service: service, logger: logger}

	http.Handle("/backup", controller)

	return controller
}
