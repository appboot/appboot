package server

import "github.com/gin-gonic/gin"

func registerRouter(router *gin.Engine) {
	router.GET("/healthz", healthz)
	router.GET("/templates", getTemplates)
	router.GET("/templates/git_hash", getTemplatesGitHash)
	router.PUT("/templates", updateTemplates)
	router.GET("/configs/:template", getTemplateConfig)
	router.POST("/app", createApp)
}
