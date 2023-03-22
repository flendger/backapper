package basecontroller

import "github.com/gin-gonic/gin"

type BaseController struct {
}

func (c *BaseController) Info(status int, msg string, context *gin.Context) {
	context.String(status, msg)
	context.Writer.Flush()

	_, err := gin.DefaultWriter.Write([]byte(msg))
	if err != nil {
		return
	}
}
