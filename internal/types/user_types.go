package types

type LoginResponse struct {
	Message    string `json:"message"`
	Token      string `json:"token"`
	Status     int    `json:"status"`
	Successful bool   `json:"successful"`
}

type UserRegistration struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Dob         string `json:"dob"`
	Residential string `json:"residential"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}