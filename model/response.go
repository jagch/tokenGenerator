package model

type Response struct {
	Status  int8   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
