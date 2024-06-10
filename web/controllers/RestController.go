package controllers

import (
	"GoDockerSandbox/application/services"
	"GoDockerSandbox/domain/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RestController struct {
	//tpl      *template.Template
	sbox *services.SandboxManager
}

type ImageNames struct {
	ImageNames []string `json:"image_names"`
}

type RawCompose struct {
	Id       string                `json:"id"`
	Services []model.DockerService `json:"docker_services"`
}

func NewRestController(sbox *services.SandboxManager) *RestController {
	return &RestController{
		sbox: sbox,
	}
}

func (rc *RestController) CreateSandbox(c *gin.Context) {
	c.HTML(http.StatusOK, "create_sandbox.html", nil)
}

// get images via docker client
func (rc *RestController) GetImages(c *gin.Context) {
	err, images := rc.sbox.GetImages(c.Param("name"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "select_images.html", images)
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
	id, err := rc.sbox.SaveSandbox(rawCompose.Id, rawCompose.Services)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"id": id})
}

func (rc *RestController) UpdateCompose(c *gin.Context) {
	yamlRaw, err := c.GetRawData()
	yaml := string(yamlRaw)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	if _, err = rc.sbox.UpdateSandbox(id, yaml); err != nil {
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

func (rc *RestController) RunCompose(c *gin.Context) {
	id := c.Param("id")

	yamlRaw, err := c.GetRawData()
	yaml := string(yamlRaw)

	containers, err := rc.sbox.RunSandbox(id, yaml)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"containers": containers})
}

/*func (rc *RestController) GetRunningContainers(c *gin.Context) {
	// todo
}*/

func (rc *RestController) StopCompose(c *gin.Context) {
	id := c.Param("id")

	err := rc.sbox.StopSandbox(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "running_sandbox.html", []string{}) // fixme
}
