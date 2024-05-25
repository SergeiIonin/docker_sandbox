package main

import (
	"GoDockerSandbox/controllers"
	"GoDockerSandbox/services"
	"GoDockerSandbox/infra"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	ds := services.NewDockerService()
	cache := infra.NewImageCacheInMemImpl()
	sbox := services.NewSandboxService(cache)
	rc := controllers.NewRestController(ds, sbox)
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/create_sandbox")
	})
	router.GET("/create_sandbox", rc.CreateSandbox)
	router.GET("/images/app", rc.GetImages)
	router.GET("/images/:name", rc.GetImagesByName)
	router.POST("/images/build-compose", rc.BuildCompose)
	router.POST("/containers/:name/start", rc.StartContainer)
	router.POST("/containers/:name/stop", rc.StopContainer)

	router.Run("localhost:8082")
}