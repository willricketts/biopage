package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	router := gin.Default()
	router.Static("/assets", "./bower_components")
	router.Static("/styles", "./public/styles")
	router.LoadHTMLGlob("public/*.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8080")
}
