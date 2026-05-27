package healthcheckcontroller

import (
	"backapper/app/appservice"
	"backapper/basecontroller"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

type HealthCheckController struct {
	*basecontroller.BaseController
	service *appservice.AppService
}

func (c *HealthCheckController) Handle(context *gin.Context) {
	appName := context.Query("app")
	if appName == "" {
		context.String(http.StatusBadRequest, "Bad request: no App param\n")
		return
	}

	output, err := c.service.HealthCheck(appName)
	if err != nil {
		status := http.StatusBadRequest
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			status = http.StatusInternalServerError
		}
		info := "Couldn't healthcheck app [" + appName + "]: " + err.Error() + "\n"
		c.Info(status, info, context)
		return
	}

	c.Info(http.StatusOK, output+"\nOK healthcheck: "+appName+"\n", context)
}

func New(service *appservice.AppService) *HealthCheckController {
	return &HealthCheckController{service: service}
}
