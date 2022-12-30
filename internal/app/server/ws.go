package server

import (
	"log"
	"net/http"

	"github.com/appboot/appboot/internal/app/appboot"
	"golang.org/x/net/websocket"
)

var conn *websocket.Conn

func StartServer(addr string) {
	http.Handle("/", websocket.Handler(wsHandler))

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func ShutdownServer() {
	conn.Close()
}

const (
	GetTemplates = "getTemplates"
)

type ReceiveMessage struct {
	CMD  string         `json:"cmd"`
	Data map[string]any `json:"data"`
}

func wsHandler(w *websocket.Conn) {
	conn = w
	var error error
	for {
		var reply ReceiveMessage
		if error = websocket.JSON.Receive(w, &reply); error != nil {
			log.Println("websocket receive message", error)
			break
		}

		if reply.CMD == GetTemplates {
			groups := appboot.GetTemplateGroups()
			hash := appboot.GetTemplatesGitHash()
			data := map[string]any{"cmd": GetTemplates, "groups": groups, "hash": hash}
			send(data)
		}
	}
}

func send(v any) {
	if conn == nil {
		return
	}

	if err := websocket.JSON.Send(conn, v); err != nil {
		log.Println("websocket send message", err)
	}
}
