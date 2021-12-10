package dto

import (
	companyEntity "gitlab.warungpintar.co/sales-platform/brook/domain/company/entity"
	userEntity "gitlab.warungpintar.co/sales-platform/brook/domain/user/entity"
)

type Error struct {
	IsError           bool   `json:"is_error"`
	IsShowPopup       bool   `json:"is_show_popup"`
	ErrorType         string `json:"error_type"`
	HumanErrorTitle   string `json:"human_error_title"`
	HumanErrorMessage string `json:"human_error_message"`
	ServerMessage     string `json:"server_message"`
}

type UserActivateRequest struct {
	EmployeeId  string `json:"employee_id" validate:"required"`
	CompanyCode string `json:"company_code" validate:"required"`
}

type UserInfo struct {
	User    userEntity.User       `json:"user"`
	Company companyEntity.Company `json:"company"`
}

type UserActivateResponse struct {
	User  UserInfo
	Error Error
}
