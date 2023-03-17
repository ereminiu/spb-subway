package main

import (
	"github.com/ereminiu/spb-subway/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/", handlers.HomeHandler)
	r.GET("/showroute", handlers.ShowRouteHandler)
	r.POST("/getroute", handlers.FindRouteHandler)

	r.Run(":8080")
}
