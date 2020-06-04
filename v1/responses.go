package v1

import "encoding/json"

// response defines an API response
type response struct {
	Code    int         `json:"code"`
	Type    string      `json:"type"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// successResponse generates a success response
func successResponse(code int, message string, data interface{}) []byte {
	bytes, _ := json.Marshal(&response{
		Code:    code,
		Type:    "success",
		Message: message,
		Data:    data,
	})
	return bytes
}

// errorResponse generates an error response
func errorResponse(code int, message string, data interface{}) []byte {
	bytes, _ := json.Marshal(&response{
		Code:    code,
		Type:    "error",
		Message: message,
		Data:    data,
	})
	return bytes
}
