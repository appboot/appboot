package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/appboot/appboot/configs"
	"github.com/appboot/appboot/internal/app/server"
	"github.com/go-ecosystem/utils/v2/net"
)

func main() {
	const port = ":8000"

	url := fmt.Sprintf("://%v%v", net.GetIP(), port)
	log.Printf("API_URL: %v%v%v\n", "ws", url, "/ws/")
	log.Printf("STATIC_URL: %v%v%v\n", "http", url, "/static/")

	configs.InitConfig()

	server.StartServer(port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	server.ShutdownServer()
	log.Println("Shutdown Server ...")
}
