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
	"strconv"
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
			ServerTime:   GenerateDateTime(),
		},
	})
}

func GenerateDateTime() string {
	// dd/MM/yyyy HH:mm:ss
	makeSure := func(i, count int) string {
		s := strconv.Itoa(i)
		final := count - len(s)
		for ; final > 0; final-- {
			s = "0" + s
		}

		return s
	}
	t := time.Now()

	str := makeSure(t.Day(), 2) + "/"
	str += makeSure(int(t.Month()), 2) + "/"
	str += makeSure(t.Year(), 4) + " "
	str += makeSure(t.Hour(), 2) + ":"
	str += makeSure(t.Minute(), 2) + ":"
	str += makeSure(t.Second(), 2)

	return str
}
