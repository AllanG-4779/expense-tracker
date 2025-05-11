package types

type LoginResponse struct {
	Message    string        `json:"message"`
	Token      TokenResponse `json:"body"`
	Status     int           `json:"status"`
	Successful bool          `json:"successful"`
	User      User		 `json:"user"`
}

type UserRegistration struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	// ignore this
	Password    string `json:"password,omitempty"`
	NewPassword string `json:"new_password,omitempty"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token    string `json:"token"`
	ValidFor int    `json:"valid_for"`
}

type AuthContext struct {
	Email string `json:"email"`
	Sub   string `json:"sub"`
	Iat   string `json:"iat"`
	Iss   string `json:"iss"`
}

type User struct {
	Username string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`	
	Email    string `json:"email"`	
}