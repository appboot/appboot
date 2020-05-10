package main

import (
	"fmt"
	"github.com/appboot/appboot/utils"
	"net/http"

	"github.com/appboot/appboot/handler"
	"golang.org/x/net/websocket"
)

func main() {
	const port = ":8888"
	const pattern = "/appboot"
	url := fmt.Sprintf("ws://%v:%v%v",utils.GetIP(),port,pattern)
	fmt.Printf("WS_URL: %v", url)

	http.Handle(pattern, websocket.Handler(handler.Handle))
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}
}