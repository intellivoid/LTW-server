/*
 * This file is part of LTW-server project (https://github.com/intellivoid/LTW-server).
 * Copyright (c) 2021 WotoTeam.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, version 2.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package versioning

import (
	"net/http"
	"strings"
	"time"

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
		Success: true,
		Results: &VersionResults{
			IsAcceptable: VersionAcceptable(ver, h),
			ServerTime:   time.Now().String(),
		},
	})
}
