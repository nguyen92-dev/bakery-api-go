package dto

type APIResponse[TResponse any] struct {
	Data  TResponse  `json:"data"`
	Error []APIError `json:"error"`
}

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewAPIResponse[TResponse any](data TResponse, err []APIError) APIResponse[TResponse] {
	return APIResponse[TResponse]{Data: data, Error: err}
}

func NewErrorResponse(err []APIError) APIResponse[any] {
	return NewAPIResponse[any](nil, err)
}

func NewAPIError(code string, message string) APIError {
	return APIError{Code: code, Message: message}
}
