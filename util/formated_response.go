package util

type ErrorResponse struct {
	Error string
}

type OkResponse struct {
	Message string
	Data    interface{}
}
