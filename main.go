package main

import (
	"GoDockerSandbox/application/services"
	"GoDockerSandbox/infra/mongo/repo"
	"GoDockerSandbox/web/controllers"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	composeRepo, err := repo.NewComposeMongoRepo()
	if err != nil {
		panic(err) // fixme
	}
	sbox := services.NewSandboxManager(composeRepo)
	rc := controllers.NewRestController(sbox)
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/create_sandbox")
	})
	router.GET("/create_sandbox", rc.CreateSandbox)
	router.GET("/images/:name", rc.GetImages)
	router.POST("/compose/create", rc.CreateCompose)
	router.POST("/compose/update/:id", rc.UpdateCompose)
	router.GET("/compose/:id", rc.GetCompose)
	//router.POST("/images/build-compose", rc.BuildCompose)
	//router.POST("/containers/:name/start", rc.StartContainer)
	//router.POST("/containers/:name/stop", rc.StopContainer)

	router.Run("localhost:8082")
}
