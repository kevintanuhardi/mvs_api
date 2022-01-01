package dto

type Error struct {
	IsError           bool   `json:"is_error"`
	IsShowPopup       bool   `json:"is_show_popup"`
	ErrorType         string `json:"error_type"`
	HumanErrorTitle   string `json:"human_error_title"`
	HumanErrorMessage string `json:"human_error_message"`
	ServerMessage     string `json:"server_message"`
}

type RegisterUserRequest struct {
	Name string `json:"name,omitempty" validate:"required"`
	PhoneNumber string	`json:"phone_number" validate:"required"`
	Email string	`json:"email" validate:"required,email"`
	Password string	`json:"password" validate:"required"`
	Role string `json:"role" validate:"oneof=admin super_admin"`
}

type RegisterUserResponse struct {
	ID *int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	PhoneNumber string	`json:"phone_number"`
	Email string	`json:"email"`
	Role string `json:"role"`
}

type LoginRequest struct {
	Email string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
}

type LoginResponse struct {
	RefreshToken string      `json:"refresh_token"`
	Token        string      `json:"token"`
	User         *UserInfo `json:"user"`
}

type UserInfo struct {
	ID *int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	PhoneNumber string	`json:"phone_number"`
	Email string	`json:"email"`
	Role string `json:"role"`
}
