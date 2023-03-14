package deploycontroller

import (
	"backapper/app/appservice"
	"log"
	"net/http"
)

type DeployController struct {
	service *appservice.AppService
	logger  *log.Logger
}

func (c *DeployController) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		c.writeResponse("Method not allowed", response)
		return
	}

	query := request.URL.Query()
	appParams, exists := query["app"]
	if !exists {
		info := "Bad request: no app param"
		c.writeResponse(info, response)
		return
	}

	newFile, _, err := request.FormFile("file")
	if err != nil {
		return
	}

	distInfo, err := c.service.Deploy(appParams[0], newFile)
	if err != nil {
		errInfo := "Back error: " + err.Error()
		c.writeResponse(errInfo, response)
		return
	}

	c.writeResponse("OK deploy: "+distInfo, response)
}

func New(service *appservice.AppService, logger *log.Logger) *DeployController {
	controller := &DeployController{service: service, logger: logger}

	http.Handle("/deploy", controller)

	return controller
}

func (c *DeployController) writeResponse(info string, response http.ResponseWriter) {
	c.logger.Println(info)
	_, err := response.Write([]byte(info))
	if err != nil {
		return
	}
}
