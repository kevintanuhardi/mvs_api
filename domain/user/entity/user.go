package entity

type User struct {
	ID *int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	EmployeeId string	`json:"employee_id"`
	CompanyId int `json:"company_id"`
	RoleId int `json:"role_id"`
	Active bool	`json:"active"`
	PhoneNumber string	`json:"phone_number"`
	Email string	`json:"email"`
	Password string	`json:"password"`
}

func (User) TableName() string {
	return "user"
}