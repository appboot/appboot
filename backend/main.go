package main

import (
	"fmt"
	"net/http"

	"github.com/appboot/appboot/handler"
	"github.com/appboot/appboot/utils"
	"golang.org/x/net/websocket"
)

func main() {
	fmt.Printf("IP: %v", utils.GetIP())
	http.Handle("/ws", websocket.Handler(handler.Handle))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err)
	}
}
