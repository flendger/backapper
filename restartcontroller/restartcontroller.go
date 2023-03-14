package restartcontroller

import (
	"backapper/app/appservice"
	"log"
	"net/http"
)

type RestartController struct {
	service *appservice.AppService
}

func (c *RestartController) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	appParams, exists := query["app"]
	if !exists {
		info := "Bad request: no app param"
		writeResponse(info, response)
		return
	}

	appName := appParams[0]
	err := c.service.Restart(appName)
	if err != nil {
		info := "Couldn't restart app [" + appName + "]: " + err.Error()
		writeResponse(info, response)
		return
	}

	writeResponse("OK restart: "+appName, response)
}

func New(service *appservice.AppService) *RestartController {
	controller := &RestartController{service: service}

	http.Handle("/restart", controller)

	return controller
}

func writeResponse(info string, response http.ResponseWriter) {
	log.Println(info)
	_, err := response.Write([]byte(info))
	if err != nil {
		return
	}
}
