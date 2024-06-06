package controllers

import (
	"GoDockerSandbox/application/services"
	"GoDockerSandbox/domain/model"
	"fmt"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RestController struct {
	//tpl      *template.Template
	sbox *services.SandboxService
}

type ImageNames struct {
	ImageNames []string `json:"image_names"`
}

type RawCompose struct {
	Id       string                `json:"id"`
	Services []model.DockerService `json:"docker_services"`
}

func NewRestController(sbox *services.SandboxService) *RestController {
	return &RestController{
		sbox: sbox,
	}
}

func (rc *RestController) CreateSandbox(c *gin.Context) {
	c.HTML(http.StatusOK, "create_sandbox.html", nil)
}

// get images via docker client
func (rc *RestController) GetImages(c *gin.Context) {
	images := rc.sbox.GetImages()
	c.HTML(http.StatusOK, "select_images.html", images)
}

// create docker services as they appear in the compose file
func (rc *RestController) CreateDockerServices(c *gin.Context) {
	var imageNames ImageNames

	err := c.ShouldBindJSON(&imageNames)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("Creating docker services for images:")
	for _, imageName := range imageNames.ImageNames {
		log.Println(imageName)
	}
	c.HTML(http.StatusOK, "create_docker_services.html", imageNames.ImageNames)
}

func (rc *RestController) CreateCompose(c *gin.Context) {
	var rawCompose RawCompose

	err := c.ShouldBindJSON(&rawCompose)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Creating compose file for images:")
	for _, image := range rawCompose.Services {
		fmt.Println(image)
	}
	err = rc.sbox.SaveSandbox(rawCompose.Id, rawCompose.Services)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"id": rawCompose.Id})
}

func (rc *RestController) UpdateCompose(c *gin.Context) {
	yamlRaw, err := c.GetRawData()
	yaml := string(yamlRaw)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	if err = rc.sbox.UpdateSandbox(id, yaml); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"id": id})
}

func (rc *RestController) GetCompose(c *gin.Context) {
	id := c.Param("id")
	compose, err := rc.sbox.GetSandbox(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "compose.html", compose.Yaml)
}

/* func (rc *RestController) SaveImages(c *gin.Context) {
	var data []string
	id := c.Params.ByName("id")
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rc.sbox.SaveAppImages(id, data)
	c.JSON(200, gin.H{"message": "Image saved"})
} */

/* func (rc *RestController) StartContainer(c *gin.Context) {
	name := c.Param("name")
	rc.ds.ContainersService.RunContainer(name)
	c.JSON(200, gin.H{"message": "Container started"})
}

func (rc *RestController) StopContainer(c *gin.Context) {
	name := c.Param("name")
	rc.ds.ContainersService.StopContainer(name)
	c.JSON(200, gin.H{"message": "Container stopped"})
} */
