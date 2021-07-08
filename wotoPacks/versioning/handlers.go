package versioning

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetVersionHandler(c *gin.Context) {
	userAgent := c.Request.Header.Get(userAgentKey)
	if strings.ToLower(userAgent) != userAgentValue {
		return
	}

	xSameDomain := c.Request.Header.Get(xSameDomainKey)
	if strings.ToLower(xSameDomain) != xSameDomainValue {
		return
	}

	ver := c.Request.Header.Get(ltwVersionKey)
	h := c.Request.Header.Get(ltwVersionHashKey)
	c.JSON(http.StatusOK, &VersionResp{
		IsAcceptable: VersionAcceptable(ver, h),
	})
}
