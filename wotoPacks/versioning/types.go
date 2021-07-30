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

import "ltw-server/wotoPacks/serverErrors"

type Version struct {
	Num1 uint8
	Num2 uint8
	Num3 uint8
	Num4 uint8
}

type VersionResp struct {
	Success bool                        `json:"success"`
	Error   *serverErrors.EndPointError `json:"error"`
	Results *VersionResults             `json:"result"`
}

type VersionResults struct {
	IsAcceptable bool `json:"is_acceptable"`

	DataDownloadLink *string `json:"data_download_link"`

	AppDownloadLink *string `json:"app_download_link"`

	ServerTime string `json:"server_time"`
}
