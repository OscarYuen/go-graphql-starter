package model

type LoginResponse struct {
	*Response
	JWT string `json:"jwt,omitempty"`
}
