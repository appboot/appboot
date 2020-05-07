package handler

import (
	"errors"
	"fmt"
	"strings"

	"github.com/appboot/appboot/service"

	"github.com/appboot/appbctl/creator"
	"github.com/appboot/appboot/constant"
	"github.com/appboot/appboot/model"
	"golang.org/x/net/websocket"
)

// Handle ws
func Handle(conn *websocket.Conn) {
	defer conn.Close()
	jsonHandler := websocket.JSON
	params := &model.Params{}
	for {
		if err := jsonHandler.Receive(conn, params); err != nil {
			_ = sendMessage(constant.ErrJSONHandler, err.Error(), "", jsonHandler, conn)
			break
		}

		if params.Method == constant.MethodCreateApp {
			app := params.Application
			if err := createApp(app, jsonHandler, conn); err != nil {
				break
			}
			_ = sendData(constant.OK, "create success", constant.MethodCreateApp, "", jsonHandler, conn)
			break

		} else if params.Method == constant.MethodGetTemplates {
			templates := service.GetTemplates()
			_ = sendData(constant.OK, "get templates success", constant.MethodGetTemplates, templates, jsonHandler, conn)

		} else {
			_ = sendMessage(constant.ErrUnknownMethod, "unknown method", "", jsonHandler, conn)
			break
		}
	}
}

func createApp(app model.Application, jsonHandler websocket.Codec, conn *websocket.Conn) error {
	if len(app.Name) < 1 || len(app.Template) < 1 {
		msg := "application name and template can be empty"
		_ = sendMessage(constant.ErrEmpty, msg, constant.MethodCreateApp, jsonHandler, conn)
		return errors.New(msg)
	}

	if strings.Contains(app.Name, " ") {
		msg := "application name can not contain blanks"
		_ = sendMessage(constant.ErrContainBlanks, msg, constant.MethodCreateApp, jsonHandler, conn)
		return errors.New(msg)
	}

	if err := creator.Create(app.Convert(), true, false); err != nil {
		_ = sendMessage(constant.ErrCreate, err.Error(), constant.MethodCreateApp, jsonHandler, conn)
		return err
	}

	return nil
}

func sendMessage(code constant.ErrCode,
	msg string,
	method string,
	jsonHandler websocket.Codec,
	conn *websocket.Conn) (err error) {
	res := &model.Response{
		Code:   code,
		Msg:    msg,
		Method: method,
	}
	err = jsonHandler.Send(conn, res)
	if code > constant.ErrInternal {
		fmt.Println(msg)
	}
	return
}

func sendData(code constant.ErrCode,
	msg string,
	method string,
	data interface{},
	jsonHandler websocket.Codec,
	conn *websocket.Conn) (err error) {
	res := &model.Response{
		Code:   code,
		Msg:    msg,
		Method: method,
		Data:   data,
	}
	err = jsonHandler.Send(conn, res)
	if code > constant.ErrInternal {
		fmt.Println(msg)
	}
	return
}
