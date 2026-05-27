package restartcontroller

import (
	"backapper/app/appservice"
	"backapper/basecontroller"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

type RestartController struct {
	*basecontroller.BaseController
	service *appservice.AppService
}

func (c *RestartController) Handle(context *gin.Context) {
	appName := context.Query("app")
	if appName == "" {
		context.String(http.StatusOK, "Bad request: no App param\n")
		return
	}

	output, err := c.service.Restart(appName)
	if err != nil {
		status := http.StatusBadRequest
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			status = http.StatusInternalServerError
		}
		info := "Couldn't restart app [" + appName + "]: " + err.Error() + "\n"
		c.Info(status, info, context)
		return
	}

	c.Info(http.StatusOK, output+"\nOK restart: "+appName+"\n", context)
}

func New(service *appservice.AppService) *RestartController {
	return &RestartController{service: service}
}
