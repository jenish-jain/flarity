package login

type AuthResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Picture  string `json:"picture"`
	Verified bool   `json:"verified"`
}

type UserInfo struct {
	Name string `json:"name"`
}
