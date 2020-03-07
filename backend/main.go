package main

import (
	"fmt"
	"net/http"

	"github.com/appboot/appboot/handler"
	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/ws", websocket.Handler(handler.Handle))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err)
	}
}
