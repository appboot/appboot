package handler

// ErrCode error code
type ErrCode int32

const (
	// ErrInternal code
	ErrInternal ErrCode = 500
	// ErrJSONHandler code
	ErrJSONHandler ErrCode = 501
	// ErrEmpty code
	ErrEmpty ErrCode = 502
	// ErrContainBlanks code
	ErrContainBlanks ErrCode = 503
	// ErrCreate code
	ErrCreate ErrCode = 504
)
