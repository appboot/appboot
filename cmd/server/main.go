package main

import (
	"fmt"
	"github.com/CatchZeng/gutils/net"
	"github.com/appboot/appboot/configs"
	"github.com/appboot/appboot/internal/app/server"
	"log"
	"os"
	"os/signal"
)

func main() {
	const port = ":8000"

	url := fmt.Sprintf("http://%v%v", net.GetIP(), port)
	log.Printf("API_HOST: %v\n", url)

	configs.InitConfig()

	server.Run(port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
}
