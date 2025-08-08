package dto

type APIResponse[TResponse any] struct {
	Status   int       `json:"status"`
	Success  bool      `json:"success"`
	Data     TResponse `json:"data,omitempty"`
	Error    *APIError `json:"error,omitempty"`
	Metadata any       `json:"metadata,omitempty"`
}

type APIError struct {
	ErrorType string `json:"type"`
	Message   string `json:"message"`
	Detail    any    `json:"detail,omitempty"`
}

func SuccessResponse[TResponse any](status int, data TResponse) APIResponse[TResponse] {
	return APIResponse[TResponse]{Status: status, Data: data, Success: true}
}

func ErrorResponse(status int, err APIError) APIResponse[any] {
	return APIResponse[any]{Status: status, Error: &err, Success: false}
}

func NewAPIError(errType string, message string, details any) APIError {
	return APIError{ErrorType: errType, Message: message, Detail: details}
}
