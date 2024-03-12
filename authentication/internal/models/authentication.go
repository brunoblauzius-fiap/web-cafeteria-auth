package models

type AuthenticationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticationResponse struct {
	AccessToken string
	ExpiresIn   int32
	TokenType   string
}
