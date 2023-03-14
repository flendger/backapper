package restartcontroller

import (
	"backapper/app/appservice"
	"log"
	"net/http"
)

type RestartController struct {
	service *appservice.AppService
	logger  *log.Logger
}

func (c *RestartController) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	appParams, exists := query["app"]
	if !exists {
		info := "Bad request: no app param"
		c.writeResponse(info, response)
		return
	}

	appName := appParams[0]
	err := c.service.Restart(appName)
	if err != nil {
		info := "Couldn't restart app [" + appName + "]: " + err.Error()
		c.writeResponse(info, response)
		return
	}

	c.writeResponse("OK restart: "+appName, response)
}

func New(service *appservice.AppService, logger *log.Logger) *RestartController {
	controller := &RestartController{service: service, logger: logger}

	http.Handle("/restart", controller)

	return controller
}

func (c *RestartController) writeResponse(info string, response http.ResponseWriter) {
	c.logger.Println(info)
	_, err := response.Write([]byte(info))
	if err != nil {
		return
	}
}
