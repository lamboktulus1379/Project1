package Formatters

import (
	"mygra.tech/project1/Utils/Constants"
	"mygra.tech/project1/Utils/Responses"
)


func Format(data interface{}, statusCode string, statusMessage string) Responses.ResponseApi {

	response := Responses.ResponseApi{}

	response.StatusCode = statusCode
	response.StatusMessage = statusMessage
	response.Data = data

	if statusCode == "" {
		response.StatusCode = Constants.SUCCESS_RC200
	}
	if statusMessage == "" {
		response.StatusMessage = Constants.SUCCESS_RM200
	}

	return response
}