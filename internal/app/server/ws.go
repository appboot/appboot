package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/appboot/appboot/internal/app/appboot"
	"github.com/appboot/appboot/internal/pkg/common"
	"github.com/appboot/appboot/internal/pkg/zip"
	"github.com/go-ecosystem/utils/v2/response"
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
	GetTemplates      = "getTemplates"
	UpdateTemplates   = "updateTemplates"
	GetTemplateConfig = "getTemplateConfig"
	CreateApp         = "createApp"
)

type Message struct {
	CMD  string         `json:"cmd"`
	Data map[string]any `json:"data"`
}

type ResponseMessage struct {
	Message
	Error response.Error `json:"error"`
}

func wsHandler(w *websocket.Conn) {
	conn = w
	var error error
	for {
		var msg Message
		if error = websocket.JSON.Receive(w, &msg); error != nil {
			log.Println("websocket receive message", error)
			break
		}

		if msg.CMD == GetTemplates {
			send(ResponseMessage{
				Message: Message{
					CMD:  msg.CMD,
					Data: getTemplatesResponseData(),
				},
			})

		} else if msg.CMD == UpdateTemplates {
			if err := appboot.UpdateAllTemplates(); err != nil {
				log.Printf("Failed to update all templates: %v", err)
				send(ResponseMessage{
					Message: Message{
						CMD: msg.CMD,
					},
					Error: common.UpdateTemplatesError(),
				})
				return
			}

			send(ResponseMessage{
				Message: Message{
					CMD:  msg.CMD,
					Data: getTemplatesResponseData(),
				},
			})

		} else if msg.CMD == GetTemplateConfig {
			template := msg.Data["template"].(string)
			config, err := appboot.GetTemplateConfig(template)
			if err != nil {
				log.Printf("Failed to get template config: %v", err)
				send(ResponseMessage{
					Message: Message{
						CMD: msg.CMD,
					},
					Error: common.GetTemplateConfigError(),
				})
				return
			}

			send(ResponseMessage{
				Message: Message{
					CMD: msg.CMD,
					Data: map[string]any{
						"config": config,
					},
				},
			})

		} else if msg.CMD == CreateApp {
			name := msg.Data["name"].(string)
			template := msg.Data["template"].(string)

			if len(name) < 1 || len(template) < 1 || strings.Contains(name, " ") {
				send(ResponseMessage{
					Message: Message{
						CMD: msg.CMD,
					},
					Error: common.CreateAppError(),
				})
				return
			}

			params := msg.Data["params"].(string)

			const t = "true"
			skipBeforeScripts := msg.Data["skipBeforeScripts"].(string) == t
			skipAfterScripts := msg.Data["skipAfterScripts"].(string) == t

			app := appboot.Application{
				Name:       name,
				Template:   template,
				Parameters: params,
				Path:       getSavePath(name),
			}

			_ = os.RemoveAll(app.Path)

			config, err := appboot.GetTemplateConfig(template)
			if err != nil {
				log.Printf("Failed to get template config: %v", err)
				send(ResponseMessage{
					Message: Message{
						CMD: msg.CMD,
					},
					Error: common.GetTemplateConfigError(),
				})
				return
			}

			if err = appboot.Create(app,
				true,
				config.Scripts.Before,
				config.Scripts.After,
				skipBeforeScripts,
				skipAfterScripts); err != nil {
				log.Printf("Failed to create application: %v", err)
				send(ResponseMessage{
					Message: Message{
						CMD: msg.CMD,
					},
					Error: common.CreateAppError(),
				})
				return
			}

			saveName := fmt.Sprintf("%s.zip", app.Name)
			savePath := path.Join(getStaticPath(), saveName)
			err = zip.Zip(app.Path, savePath)
			if err != nil {
				send(ResponseMessage{
					Message: Message{
						CMD: msg.CMD,
					},
					Error: common.ZipAppError(),
				})
				return
			}

			downloadPath := "/static/" + saveName
			send(ResponseMessage{
				Message: Message{
					CMD: msg.CMD,
					Data: map[string]any{
						"path": downloadPath,
					},
				},
			})
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
