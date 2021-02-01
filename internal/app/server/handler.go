package server

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/appboot/appboot/configs"
	"github.com/appboot/appboot/internal/app/appboot"
	"github.com/gin-gonic/gin"
)

func healthz(c *gin.Context) {
	c.String(http.StatusOK, "Hello,It works. "+configs.EnvConfig.APIVersion)
}

func getTemplates(c *gin.Context) {
	templates := appboot.GetTemplates()
	c.JSON(http.StatusOK, NewResponse(RC_OK, "get templates success", templates))
}

func updateTemplates(c *gin.Context) {
	templates := appboot.GetTemplates()
	if err := appboot.UpdateAllTemplates(); err != nil {
		log.Printf("Failed to update all templates: %v", err)
		c.JSON(http.StatusOK, NewResponse(RC_UPDATE_TEMPLATES_ERROR, "", templates))
		return
	}

	templates = appboot.GetTemplates()
	c.JSON(http.StatusOK, NewResponse(RC_OK, "update templates success", templates))
}

func getTemplateConfig(c *gin.Context) {
	template := c.Param("template")
	config, err := appboot.GetTemplateConfig(template)
	if err != nil {
		log.Printf("Failed to get template config: %v", err)
		c.JSON(http.StatusOK, NewResponse(RC_GET_TEMPLATE_CONFIG_ERROR, "", nil))
		return
	}
	c.JSON(http.StatusOK, NewResponse(RC_OK, "get template config success", config))
}

func createApp(c *gin.Context) {
	name := c.PostForm("name")
	template := c.PostForm("template")
	params := c.PostForm("params")

	if len(name) < 1 || len(template) < 1 || strings.Contains(name, " ") {
		c.JSON(http.StatusOK, NewResponse(RC_APP_PARAMS_ERROR, "", nil))
		return
	}

	app := appboot.Application{
		Name:       name,
		Template:   template,
		Parameters: params,
		Path:       getSavePath(name),
	}

	_ = os.RemoveAll(app.Path)

	if err := appboot.Create(app, true, false, false); err != nil {
		log.Printf("Failed to create application: %v", err)
		c.JSON(http.StatusOK, NewResponse(RC_CREATE_APP_ERROR, "", nil))
		return
	}

	c.JSON(http.StatusOK, NewResponse(RC_OK, "create application success", app))
}
