package controllers

import (
	"GoDockerSandbox/services"
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RestController struct {
	//tpl      *template.Template
	ds 	 	 *services.DockerService
	sbox 	 *services.SandboxService
}

type Images struct {
    Images []string `json:"images"`
}

func NewRestController(ds *services.DockerService, sbox *services.SandboxService) *RestController {
	return &RestController {
		ds: ds,
	}
}

func (rc *RestController) CreateSandbox(c *gin.Context) {
	c.HTML(http.StatusOK, "create_sandbox.html", nil)
}

func (rc *RestController) GetImages(c *gin.Context) {
	images := rc.ds.ImagesService.GetImages()
	c.HTML(http.StatusOK, "select_app_images.html", images)
}

func (rc *RestController) GetImagesByName(c *gin.Context) {
	name := c.Param("name")
	result := rc.ds.ImagesService.GetImagesByName(name)
	c.JSON(200, result)
}

func (rc *RestController) SaveImages(c *gin.Context) {
	var data []string
	id := c.Params.ByName("id")
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rc.sbox.SaveAppImages(id, data)
	c.JSON(200, gin.H{"message": "Image saved"})
}

func (rc *RestController) BuildCompose(c *gin.Context) {
	var images Images

	bodyBytes, err := io.ReadAll(c.Request.Body)
    if err != nil {
        log.Printf("Error reading body: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    log.Printf("Building compose file with images, %v", string(bodyBytes))

	// todo can we do any better?
    c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	
	err = c.ShouldBindJSON(&images)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	compose := rc.ds.ComposeBuilder.BuildComposeFile(images.Images, "sandbox_network_"+uuid.New().String())
	c.String(200, compose)
}

func (rc *RestController) StartContainer(c *gin.Context) {
	name := c.Param("name")
	rc.ds.ContainersService.RunContainer(name)
	c.JSON(200, gin.H{"message": "Container started"})
}

func (rc *RestController) StopContainer(c *gin.Context) {
	name := c.Param("name")
	rc.ds.ContainersService.StopContainer(name)
	c.JSON(200, gin.H{"message": "Container stopped"})
}