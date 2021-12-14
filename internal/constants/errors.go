package constants

import (
	"errors"
	"fmt"
)

func GetErrDatabaseError() error {
	return errors.New("3001||database Error")
}

type CustomError struct {
	message string
	// statusCode string
}

func (ce *CustomError) Error() string {
	return ce.message
}

func GetDuplicateUserError() *CustomError {
	return &CustomError{
		message: "user is already registered",
	}
}

func GetUserNotFoundError() *CustomError {
	return &CustomError{
		message: "user tidak ditemukan",
	}
}

func GetPhoneNumberMinError() *CustomError {
	return &CustomError{
		message: "penulisan nomor HP minimal 10 digit. Silakan coba lagi",
	}
}

func GetPhoneNumberMaxError() *CustomError {
	return &CustomError{
		message: "nomor HP melebihi batas maksimum karakter",
	}
}

func GetNotNumericError(field string) *CustomError {
	return &CustomError{
		message: fmt.Sprintf("%s harus dalam karakter numerik 0-9", field),
	}
}

func GetFieldRequiredError(field string) *CustomError {
	return &CustomError{
		message: fmt.Sprintf("%s tidak boleh kosong", field),
	}
}

func GetOtpNotFoundError() *CustomError {
	return &CustomError{
		message: "Kode OTP Salah. Silakan coba lagi",
	}
}

func GetOtpExpiredError() *CustomError {
	return &CustomError{
		message: "Kode OTP Expired. Silakan coba lagi",
	}
}

func GetWrongPassError() *CustomError {
	return &CustomError{
		message: "Kode PIN yang dimasukkan tidak sesuai",
	}
}

func GetCompanyCodeNotFoundError() *CustomError {
	return &CustomError{
		message: "Company Code tidak ditemukani",
	}
}

func GetEmployeeIdNotFoundError() *CustomError {
	return &CustomError{
		message: "BFF ID tidak ditemukan",
	}
}

func GetEmployeeAlreadyActivatedError() *CustomError {
	return &CustomError{
		message: "BFF ID sudah pernah didaftarkan. Silakan Login",
	}
}

func GetCustomError(message string) error {
	return errors.New(message)
}
