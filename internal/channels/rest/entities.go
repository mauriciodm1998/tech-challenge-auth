package rest

type LoginRequest struct {
	Document string `json:"document"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type Response struct {
	Message string `json:"message"`
}
