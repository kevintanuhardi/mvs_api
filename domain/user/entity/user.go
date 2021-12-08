package entity

type User struct {
	ID *int `json:"id,omitempty"`
	Name string `json:"name,omitempty" validate:"required"`
	EmployeeId string	`json:"employee_id" validate:"required"`
	CompanyId int `json:"company_id" validate:"required"`
	RoleId int `json:"role_id" validate:"required"`
	Active bool	`json:"active" validate:"required"`
	PhoneNumber string	`json:"phone_number" validate:"required"`
	Email string	`json:"email" validate:"required,email"`
	Password string	`json:"password" validate:"required"`
}

func (User) TableName() string {
	return "user"
}