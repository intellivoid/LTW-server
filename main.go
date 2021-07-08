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

package main

import (
	"log"
	"os"

	"ltw-server/wotoPacks/entryPoints"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	log.Println("Port value has been set to: " + port)

	router := gin.New()
	router.Use(gin.Logger())
	//router.LoadHTMLGlob("templates/*.tmpl.html")
	//router.Static("/static", "static")

	err := entryPoints.LoadEntryPoints(router)
	if err != nil {
		log.Fatal(err)
	}

	router.Run(":" + port)
}
