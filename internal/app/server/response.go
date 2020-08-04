package server

import "log"

//APIResponse APIResponse
type APIResponse struct {
	Code    ResponseCode `json:"code" binding:"required"`
	Data    interface{}  `json:"data"`
	Message string       `json:"message"`
}

//NewMapResponse NewMapResponse
func NewMapResponse() *APIResponse {
	return &APIResponse{
		Code:    RC_OK,
		Message: "",
		Data:    make(map[string]interface{}),
	}
}

//NewArrayResponse NewArrayResponse
func NewArrayResponse() *APIResponse {
	return &APIResponse{
		Code:    RC_OK,
		Message: "",
		Data:    make([]interface{}, 0),
	}
}

//NewResponse NewResponse
func NewResponse(code ResponseCode, message string, data interface{}) *APIResponse {
	if message == "" {
		message = ResponseMessage[code]
	}
	res := &APIResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return res
}

//EmptyMapData EmptyMapData
var EmptyMapData = make(map[string]interface{})

//EmptyArrayData EmptyArrayData
var EmptyArrayData = make([]interface{}, 0)

//Response Response
func Response(code ResponseCode, message string, data interface{}) {
	res := NewResponse(code, message, data)
	panic(res)
}

//ResponseIfError ResponseIfError
func ResponseIfError(err error) {
	if err != nil {
		log.Println(err)
		Response(RC_SYSTEM_UNKOWN, err.Error(), EmptyMapData)
	}
}
