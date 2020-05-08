package handler

import (
	"fmt"

	"github.com/appboot/appboot/service"

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

			if code, err := service.CreateApp(app); err != nil {
				_ = sendMessage(code, err.Error(), constant.MethodCreateApp, jsonHandler, conn)
				break
			}

			if err := service.PushCode(app); err != nil {
				_ = sendMessage(constant.ErrPushCode, "push code failed", constant.MethodCreateApp, jsonHandler, conn)
				break
			}

			_ = sendMessage(constant.OK, "create success", constant.MethodCreateApp, jsonHandler, conn)
			break

		} else if params.Method == constant.MethodGetTemplates {
			templates := service.GetTemplates()
			_ = sendData(constant.OK, "get templates success", constant.MethodGetTemplates, templates, jsonHandler, conn)

		} else if params.Method == constant.MethodGetParams {
			app := params.Application
			result := service.GetParams(app.Template)
			_ = sendData(constant.OK, "get params success", constant.MethodGetParams, result, jsonHandler, conn)

		} else {
			_ = sendMessage(constant.ErrUnknownMethod, "unknown method", "", jsonHandler, conn)
			break
		}
	}
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
