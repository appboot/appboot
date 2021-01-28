package server

import (
	"log"
	"net/http"
	"time"

	"github.com/appboot/appboot/configs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	maxAgeInMinutes  = 10
	timeoutInSeconds = 30
)

//Run run server with port
func Run(port string) {
	setGinMode()
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTION"},
		AllowHeaders:     []string{"utoken,x-auth-token,x-request-id,Content-Type,Accept,Origin,Access-Control-Allow-Origin", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           maxAgeInMinutes * time.Minute,
	}))
	registerRouter(router)
	startHTTPServer(router, port)
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

func startHTTPServer(router *gin.Engine, port string) {
	srv := &http.Server{
		Addr:         port,
		Handler:      router,
		ReadTimeout:  timeoutInSeconds * time.Second,
		WriteTimeout: timeoutInSeconds * time.Second,
	}

	go func() {
		log.Println("Start Http Server ", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to serve: ", err)
		}
	}()
}
