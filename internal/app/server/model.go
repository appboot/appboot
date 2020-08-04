package server

import (
	"github.com/appboot/appboot/internal/app/appboot"
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
func (app *Application) Convert() appboot.Application {
	result := appboot.Application{}
	result.Name = app.Name
	result.Template = app.Template
	result.Parameters = app.Params
	result.Path = getSavePath(app.Name)
	return result
}
