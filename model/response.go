package model

type Response struct {
	Code  int    `json:"code"`
	Error string `json:"error,omitempty"`
}
