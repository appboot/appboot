package server

// ResponseCode ResponseCode
type ResponseCode int

const (
	RC_OK                        ResponseCode = 0
	RC_SYSTEM_UNKOWN             ResponseCode = 10000
	RC_REQUEST_PARAM             ResponseCode = 10001
	RC_UPDATE_TEMPLATES_ERROR    ResponseCode = 10002
	RC_GET_TEMPLATE_CONFIG_ERROR ResponseCode = 10003
	RC_APP_PARAMS_ERROR          ResponseCode = 10004
	RC_CREATE_APP_ERROR          ResponseCode = 10005
)

//ResponseMessage ResponseMessage
var ResponseMessage = map[ResponseCode]string{
	RC_OK:                        "OK",
	RC_SYSTEM_UNKOWN:             "unknown error",
	RC_REQUEST_PARAM:             "request params error",
	RC_UPDATE_TEMPLATES_ERROR:    "update templates error",
	RC_GET_TEMPLATE_CONFIG_ERROR: "get template config error",
	RC_APP_PARAMS_ERROR:          "application params error",
	RC_CREATE_APP_ERROR:          "create application error",
}
