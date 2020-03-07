package handler

import (
	"github.com/appboot/appbctl/application"
)

// Application struct
type Application struct {
	Name     string
	Template string
	params   string
}

// Convert to appboot Application
func (app *Application) Convert() application.Application {
	result := application.Application{}
	result.Name = app.Name
	result.Template = app.Template
	result.Values = app.params
	result.Path = getSavePath(app.Name)
	return result
}

// Response struct
type Response struct {
	Code ErrCode
	Msg  string
	Data interface{}
}
