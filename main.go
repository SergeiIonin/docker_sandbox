package main

import (
	"GoDockerSandbox/application/services"
	"GoDockerSandbox/infra/clients/docker"
	"GoDockerSandbox/infra/mongo/repo"
	"GoDockerSandbox/web/controllers"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	mongoUri := viper.GetString("mongo.uri")
	db := viper.GetString("mongo.database")
	collName := viper.GetString("mongo.collection")

	composeRepo, err := repo.NewComposeMongoRepo(mongoUri, db, collName)
	if err != nil {
		log.Fatalf("Failed to initialize ComposeMongoRepo: %v", err)
	}

	err = docker.CreateDockerApiClient("1.45")
	if err != nil {
		log.Fatalf("Failed to create docker client: %v", err)
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
