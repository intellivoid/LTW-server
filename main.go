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
	log.Println("we are here!" + port)

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
