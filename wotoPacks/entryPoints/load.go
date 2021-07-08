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

	//router.GET(wotoValues.Slash, unknownHandler)
	//router.POST(wotoValues.Slash, unknownHandler)

	router.GET(wotoValues.GetVersion, versioning.GetVersionHandler)
	return nil
}
