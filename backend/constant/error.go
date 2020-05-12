package constant

// ErrCode error code
type ErrCode int32

const (
	// OK code
	OK ErrCode = 0
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
	// ErrUnknownMethod code
	ErrUnknownMethod ErrCode = 505
	// ErrPushCode code
	ErrPushCode ErrCode = 506
)
