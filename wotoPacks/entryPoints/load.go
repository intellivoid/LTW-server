package entryPoints

import (
	"errors"

	"ltw-server/wotoPacks/versioning"
	"ltw-server/wotoPacks/wotoValues"

	"github.com/gin-gonic/gin"
)

func LoadEntryPoints(router *gin.Engine) error {
	if router == nil {
		return errors.New("router (gin.Engine) cannot be nil")
	}

	router.GET(wotoValues.Slash, unknownHandler)
	//router.POST(wotoValues.Slash, unknownHandler)

	router.GET(wotoValues.GetVersion, versioning.GetVersionHandler)
	return nil
}
