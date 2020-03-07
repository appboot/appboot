package handler

import (
	"fmt"
	"strings"

	"github.com/appboot/appbctl/creator"
	"golang.org/x/net/websocket"
)

// Handle ws
func Handle(conn *websocket.Conn) {
	defer conn.Close()
	jsonHandler := websocket.JSON
	app := &Application{}
	for {
		if err := jsonHandler.Receive(conn, app); err != nil {
			send(ErrJSONHandler, err.Error(), jsonHandler, conn)
			break
		}

		if len(app.Name) < 1 || len(app.Template) < 1 {
			send(ErrEmpty, "application name and template can be empty", jsonHandler, conn)
			break
		}
		if strings.Contains(app.Name, " ") {
			send(ErrContainBlanks, "application name can not contain blanks", jsonHandler, conn)
			break
		}

		if err := creator.Create(app.Convert(), true, false); err != nil {
			send(ErrCreate, err.Error(), jsonHandler, conn)
			break
		}
	}
}

func send(code ErrCode,
	msg string,
	jsonHandler websocket.Codec,
	conn *websocket.Conn) (err error) {
	res := &Response{
		Code: code,
		Msg:  msg,
	}
	err = jsonHandler.Send(conn, res)
	if code > ErrInternal {
		fmt.Println(msg)
	}
	return
}
