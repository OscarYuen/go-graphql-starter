package model

type LoginResponse struct {
	*Response
	AccessToken string `json:"access_token,omitempty"`
}
