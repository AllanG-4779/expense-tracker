package types

type UserLoginResponse struct {
	Message string `json:"message"`
	Token string `json:"token"`
	Status int `json:"status"`
	Successful bool `json:"successful"`
}