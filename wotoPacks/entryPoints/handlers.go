package entryPoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func unknownHandler(c *gin.Context) {
	c.String(http.StatusAccepted, "test")
}
