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

	url := fmt.Sprintf("ws://%v%v", net.GetIP(), port)
	log.Printf("API_URL: %v\n", url)

	configs.InitConfig()

	server.StartServer(port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	server.ShutdownServer()
	log.Println("Shutdown Server ...")
}
