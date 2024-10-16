package main

import (
	"GoDockerSandbox/application/services"
	"GoDockerSandbox/infra/mongo/repo"
	"GoDockerSandbox/web/controllers"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	mongoUri := "mongodb://localhost:2717"
	composeRepo, err := repo.NewComposeMongoRepo(mongoUri)
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
	router.POST("/compose/:id/run", rc.RunCompose)
	router.GET("/compose/:id/containers", rc.GetRunningContainers)
	router.POST("/compose/:id/stop", rc.StopCompose)

	err = router.Run("localhost:8082")
	if err != nil {
    	log.Fatalf("Failed to run server: %v", err)
	}
}
