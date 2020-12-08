package main

import (
	"github.com/Ardesco/credit-card-generator/api/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/ping", controllers.Ping)
	router.StaticFS("./css", http.Dir("css"))
	router.StaticFS("./fonts", http.Dir("fonts"))
	router.StaticFS("./images", http.Dir("images"))
	router.StaticFS("./js", http.Dir("js"))
	router.GET("/", func(c *gin.Context) {
		c.File("./index.html")
	})

	router.GET("/api/list", controllers.ListTypes)
	router.GET("/api/generate/cards", controllers.GenerateCards)
	router.GET("/api/generate/card", controllers.GenerateCard)

	router.Run()
}
