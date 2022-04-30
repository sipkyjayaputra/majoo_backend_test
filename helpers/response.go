package helpers

import "majoo/model/dto"

func SendResponse(status_code int, message string, err interface{}, data interface{}) dto.Response {
	return dto.Response{
		StatusCode: status_code,
		Status: message,
		Error: err,
		Data: data,
	}
}