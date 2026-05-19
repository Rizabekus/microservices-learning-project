package http

type UserCreateDTO struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	MobileNumber string `json:"mobile_number"`
}

type RegisterResponseDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginResponseDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type RefreshDTO struct {
	RefreshToken string `json:"refresh_token"`
}
type RefreshResponseDTO struct {
	AccessToken string `json:"access_token"`
}
