package Responses

type ResponseApi struct {
	StatusCode    string      `json:"statusCode"`
	StatusMessage string      `json:"statusMessage"`
	Data          interface{} `json:"data"`
}