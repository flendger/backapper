package deploycontroller

import (
	"backapper/app/appservice"
	"log"
	"net/http"
)

type DeployController struct {
	service *appservice.AppService
}

func (d *DeployController) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		writeResponse("Method not allowed", response)
		return
	}

	query := request.URL.Query()
	appParams, exists := query["app"]
	if !exists {
		info := "Bad request: no app param"
		writeResponse(info, response)
		return
	}

	newFile, _, err := request.FormFile("file")
	if err != nil {
		return
	}

	distInfo, err := d.service.Deploy(appParams[0], newFile)
	if err != nil {
		errInfo := "Back error: " + err.Error()
		writeResponse(errInfo, response)
		return
	}

	writeResponse("OK deploy: "+distInfo, response)
}

func New(service *appservice.AppService) *DeployController {
	controller := &DeployController{service: service}

	http.Handle("/deploy", controller)

	return controller
}

func writeResponse(info string, response http.ResponseWriter) {
	log.Println(info)
	_, err := response.Write([]byte(info))
	if err != nil {
		return
	}
}
