package dto

type APIResponse[TResponse any] struct {
	Status   int       `json:"status"`
	Sucess   bool      `json:"sucess"`
	Data     TResponse `json:"data"`
	Error    APIError  `json:"error"`
	Metadata any       `json:"metadata"`
}

type APIError struct {
	ErrorType string `json:"type"`
	Message   string `json:"message"`
	Detail    any    `json:"detail"`
}

func SuccessResponse[TResponse any](status int, data TResponse) APIResponse[TResponse] {
	return APIResponse[TResponse]{Status: status, Data: data}
}

func ErrorResponse(status int, err APIError) APIResponse[any] {
	return APIResponse[any]{Status: status, Error: err}
}

func NewAPIError(errType string, message string, details any) APIError {
	return APIError{ErrorType: errType, Message: message, Detail: details}
}
