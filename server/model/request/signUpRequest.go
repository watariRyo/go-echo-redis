package request

type SignUpRequest struct {
	Name string `json:"name"`
	Password string `json:"password"`
	Message string `json:"message"`
}
