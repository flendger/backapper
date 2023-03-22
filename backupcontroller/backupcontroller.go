package backupcontroller

import (
	"backapper/app/appservice"
	"backapper/basecontroller"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BackupController struct {
	*basecontroller.BaseController
	service *appservice.AppService
}

func (c *BackupController) Handle(context *gin.Context) {
	c.Info(http.StatusOK, "Starting backup...\n", context)

	appName := context.Query("app")
	if appName == "" {
		c.Info(http.StatusBadRequest, "Bad request: no App param\n", context)
		return
	}

	backUp, err := c.service.BackUp(appName)
	if err != nil {
		errInfo := "Backup error: " + err.Error() + "\n"
		c.Info(http.StatusBadRequest, errInfo, context)
		return
	}

	c.Info(http.StatusOK, "OK backup: "+backUp+"\n", context)
}

func New(service *appservice.AppService) *BackupController {
	return &BackupController{service: service}
}
