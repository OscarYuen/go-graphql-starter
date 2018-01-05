package model

type LoginResponse struct {
	Response
	JWT string
}

func NewLoginResponse() *LoginResponse {
	response := Response{}
	loginResponse := LoginResponse{}
	loginResponse.Response = response
	return &loginResponse
}
