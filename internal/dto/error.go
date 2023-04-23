package dto

type ErrorOutputDTO struct {
	Err     string `json:"error"`
	Message string `json:"message"`
	Code    uint   `json:"code"`
}
