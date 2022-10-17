package response

type LoginResponse struct {
	Name string `json:"name"`
	Token string `json:"token"`
}