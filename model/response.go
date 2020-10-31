package models

type Map map[string]interface{}

type ErrorResponse struct {
	Error   bool   `json: error`
	Message string `json: message`
}
