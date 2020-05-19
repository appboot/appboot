package main

import (
	"fmt"
	"net/http"

	"github.com/appboot/appboot/utils"

	"github.com/appboot/appboot/service"

	"github.com/appboot/appboot/handler"
	"golang.org/x/net/websocket"
)

func main() {
	const port = ":8888"
	const pattern = "/appboot"
	url := fmt.Sprintf("ws://%v%v%v", utils.GetIP(), port, pattern)
	fmt.Printf("WS_URL: %v\n", url)

	service.InitAppbctlConfig()

	http.Handle(pattern, websocket.Handler(handler.Handle))
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
