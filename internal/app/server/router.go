package server

import "github.com/gin-gonic/gin"

func registerRouter(router *gin.Engine) {
	router.GET("/healthz", healthz)
	router.GET("/templates", getTemplates)
	router.PUT("/templates", updateTemplates)
	router.GET("/configs/:template", getTemplateConfig)
	router.POST("/app", createApp)
}
