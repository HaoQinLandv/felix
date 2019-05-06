package models

type JsonResponse struct {
	HasError bool        `json:"has_error"`
	Message  interface{} `json:"message"`
	Data     interface{} `json:"data"`
}
