package server

import (
	"context"
	"net/http"
	"time"

	"github.com/appboot/appboot/configs"
	"github.com/gin-gonic/gin"
	"github.com/go-ecosystem/log"
	"github.com/go-ecosystem/utils/v2/middleware"
)

const (
	maxWaitTimeBeforeShutdown = 10
)

var srv *http.Server

// Run run server with port
func Run(port string) {
	setGinMode()
	router := gin.New()
	router.Use(middleware.Cors())
	registerRouter(router)
	startServer(router, port)
}

// ShutdownServer ShutdownServer
func ShutdownServer() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*maxWaitTimeBeforeShutdown)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Print("shutdown server: ", err)
	}
}

func setGinMode() {
	env := configs.EnvConfig.ProjectEnv
	if env == "dev" {
		gin.SetMode(gin.DebugMode)
	} else if env == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func startServer(router *gin.Engine, port string) {
	srv = &http.Server{
		Addr:              port,
		Handler:           router,
		ReadTimeout:       time.Duration(configs.EnvConfig.ReadTimeout) * time.Second,
		ReadHeaderTimeout: time.Duration(configs.EnvConfig.ReadHeaderTimeout) * time.Second,
		WriteTimeout:      time.Duration(configs.EnvConfig.WriteTimeout) * time.Second,
	}

	go func() {
		log.Println("Start Http Server ", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.FatalE("Failed to serve: ", err)
		}
	}()
}
