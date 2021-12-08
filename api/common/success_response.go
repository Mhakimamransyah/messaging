package common

import (
	"net/http"
)

type SuccessResponseCode string

//List of success response status
const (
	Success SuccessResponseCode = "success"
)

//SuccessResponse default payload response
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type SuccessGetData struct {
	Message string      `json:"message"`
	Count   int         `json:"total"`
	Data    interface{} `json:"data"`
	Query   interface{} `json:"query"`
}

func NewSuccessResponseGetData(data interface{}, query interface{}, count int) (int, SuccessGetData) {
	return http.StatusOK, SuccessGetData{
		"Success",
		count,
		data,
		query,
	}
}

//NewSuccessResponse create new success payload
func NewSuccessResponse(data interface{}) (int, SuccessResponse) {
	return http.StatusOK, SuccessResponse{
		"Success",
		data,
	}
}

//NewSuccessResponse create new success payload
func NewSuccessCreated() (int, SuccessResponse) {
	return http.StatusCreated, SuccessResponse{
		"Success",
		map[string]interface{}{},
	}
}

//NewSuccessResponse create new success payload
func NewSuccessResponseWithoutData() (int, SuccessResponse) {
	return http.StatusOK, SuccessResponse{
		"Success",
		map[string]interface{}{},
	}
}
