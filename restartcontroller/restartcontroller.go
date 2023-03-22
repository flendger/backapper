package restartcontroller

import (
	"backapper/app/appservice"
	"backapper/basecontroller"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RestartController struct {
	*basecontroller.BaseController
	service *appservice.AppService
}

func (c *RestartController) Handle(context *gin.Context) {
	c.Info(http.StatusOK, "Starting restart...\n", context)

	appName := context.Query("app")
	if appName == "" {
		context.String(http.StatusBadRequest, "Bad request: no App param\n")
		return
	}

	err := c.service.Restart(appName)
	if err != nil {
		info := "Couldn't restart app [" + appName + "]: " + err.Error() + "\n"
		c.Info(http.StatusBadRequest, info, context)
		return
	}

	c.Info(http.StatusOK, "OK restart: "+appName+"\n", context)
}

func New(service *appservice.AppService) *RestartController {
	return &RestartController{service: service}
}
