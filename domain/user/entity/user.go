package entity

import "time"

type User struct {
	ID *int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	IsActive bool	`json:"is_active"`
	PhoneNumber string	`json:"phone_number"`
	Email string	`json:"email"`
	Password string	`json:"password"`
	Role string `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "user"
}