package dashboard

type loginRequest struct {
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Password    string `json:"password" validate:"required,min=6"`
}
