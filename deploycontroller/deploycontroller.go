package deploycontroller

import (
	"backapper/app/appservice"
	"backapper/basecontroller"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeployController struct {
	*basecontroller.BaseController
	service *appservice.AppService
}

func (c *DeployController) Handle(context *gin.Context) {
	c.Info(http.StatusOK, "Starting deploy...\n", context)

	appName := context.Query("app")
	if appName == "" {
		context.String(http.StatusBadRequest, "Bad request: no App param\n")
		return
	}

	msgStart := "Upload starting... \n"
	c.Info(http.StatusOK, msgStart, context)

	newFile, _, err := context.Request.FormFile("file")
	if err != nil {
		errInfo := "Deploy error cause couldn't get file: " + err.Error() + "\n"
		c.Info(http.StatusBadRequest, errInfo, context)
		return
	}

	distInfo, err := c.service.Deploy(appName, newFile)
	if err != nil {
		errInfo := "Deploy error cause couldn't save file: " + err.Error() + "\n"
		c.Info(http.StatusBadRequest, errInfo, context)
		return
	}

	msgInfo := "Deploy completed: " + distInfo + "\n"
	c.Info(http.StatusOK, msgInfo, context)
}

func New(service *appservice.AppService) *DeployController {
	return &DeployController{service: service}
}
