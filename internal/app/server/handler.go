package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/appboot/appboot/configs"
	"github.com/appboot/appboot/internal/app/appboot"
	"github.com/appboot/appboot/internal/pkg/common"
	"github.com/appboot/appboot/internal/pkg/zip"
	"github.com/gin-gonic/gin"
	"github.com/go-ecosystem/utils/response"
)

func healthz(c *gin.Context) {
	c.String(http.StatusOK, "Hello,It works. "+configs.EnvConfig.APIVersion)
}

func getTemplates(c *gin.Context) {
	templates := appboot.GetTemplates()
	response.OK(c, "get templates success", templates)
}

func updateTemplates(c *gin.Context) {
	if err := appboot.UpdateAllTemplates(); err != nil {
		log.Printf("Failed to update all templates: %v", err)
		response.Err(c, common.UpdateTemplatesError())
		return
	}

	templates := appboot.GetTemplates()
	response.OK(c, "update templates success", templates)
}

func getTemplatesGitHash(c *gin.Context) {
	hash := appboot.GetTemplatesGitHash()
	response.OK(c, "get templates git hash", hash)
}

func getTemplateConfig(c *gin.Context) {
	template := c.Param("template")
	config, err := appboot.GetTemplateConfig(template)
	if err != nil {
		log.Printf("Failed to get template config: %v", err)
		response.Err(c, common.GetTemplateConfigError())
		return
	}
	response.OK(c, "get template config success", config)
}

func createApp(c *gin.Context) {
	name := c.PostForm("name")
	template := c.PostForm("template")
	params := c.PostForm("params")

	if len(name) < 1 || len(template) < 1 || strings.Contains(name, " ") {
		response.Err(c, common.AppParamsError())
		return
	}

	skipBeforeScripts := c.PostForm("skipBeforeScripts") == "true"
	skipAfterScripts := c.PostForm("skipAfterScripts") == "true"

	app := appboot.Application{
		Name:       name,
		Template:   template,
		Parameters: params,
		Path:       getSavePath(name),
	}

	_ = os.RemoveAll(app.Path)

	config, err := appboot.GetTemplateConfig(template)
	if err != nil {
		log.Printf("Failed to get template config: %v", err)
		response.Err(c, common.GetTemplateConfigError())
		return
	}

	if err := appboot.Create(app,
		true,
		config.Scripts.Before,
		config.Scripts.After,
		skipBeforeScripts,
		skipAfterScripts); err != nil {
		log.Printf("Failed to create application: %v", err)
		response.Err(c, common.CreateAppError())
		return
	}

	saveName := fmt.Sprintf("%s.zip", app.Name)
	savePath := path.Join(getStaticPath(), saveName)
	err = zip.Zip(app.Path, savePath)
	if err != nil {
		response.Err(c, common.ZipAppError())
		return
	}

	downloadPath := "/static/" + saveName
	response.OK(c, "create application success", downloadPath)
}
