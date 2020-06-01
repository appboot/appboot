package model

import (
	"github.com/appboot/appbctl/application"
	"github.com/appboot/appboot/constant"
	"github.com/appboot/appboot/utils"
)

// Params params struct
type Params struct {
	Method      string      `json:"method"`
	Application Application `json:"application,omitempty"`
}

// Application struct
type Application struct {
	Name     string `json:"name"`
	Template string `json:"template"`
	Params   string `json:"params"`
}

// Convert to appboot Application
func (app *Application) Convert() application.Application {
	result := application.Application{}
	result.Name = app.Name
	result.Template = app.Template
	result.Parameters = app.Params
	result.Path = utils.GetSavePath(app.Name)
	return result
}

// Response struct
type Response struct {
	Code   constant.ErrCode `json:"code"`
	Method string           `json:"method"`
	Msg    string           `json:"msg"`
	Data   interface{}      `json:"data"`
}
