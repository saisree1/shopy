package domain

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
}

func MapResponse(code int, message string, body interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Body:    body,
	}
}

func MapResponseWithoutBody(code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Body:    nil,
	}
}
