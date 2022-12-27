package common

import "github.com/go-ecosystem/utils/v2/response"

// response.Code
const (
	CodeUpdateTemplates   response.Code = 10002
	CodeGetTemplateConfig response.Code = 10003
	CodeAppParams         response.Code = 10004
	CodeCreateApp         response.Code = 10005
	CodeZipApp            response.Code = 10006
)

// System error
var (
	UpdateTemplatesError = func() response.Error {
		return response.Error{Code: CodeUpdateTemplates, Msg: "update templates error", Detail: ""}
	}

	GetTemplateConfigError = func() response.Error {
		return response.Error{Code: CodeGetTemplateConfig, Msg: "get template config error", Detail: ""}
	}

	AppParamsError = func() response.Error {
		return response.Error{Code: CodeAppParams, Msg: "application params error", Detail: ""}
	}

	CreateAppError = func() response.Error {
		return response.Error{Code: CodeCreateApp, Msg: "create application error", Detail: ""}
	}

	ZipAppError = func() response.Error {
		return response.Error{Code: CodeZipApp, Msg: "zip application error", Detail: ""}
	}
)
