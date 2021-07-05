package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var _count int

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	log.Println("we are here!" + port)
	//wotoPacks.ServerDaytime(port)

	router := gin.New()
	router.Use(gin.Logger())
	//router.LoadHTMLGlob("templates/*.tmpl.html")
	//router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		//c.HTML(http.StatusOK, "index.tmpl.html", nil)
		_s := c.Request.Header.Get("test")
		_count++
		c.String(http.StatusOK, "%v %v; counte: %v", "yes!", _s, _count)
	})

	router.Run(":" + port)
}
