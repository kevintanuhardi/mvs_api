package constants

import "errors"

func GetErrDatabaseError() error {
	return errors.New("3001||database Error")
}

func GetDuplicateUserError() error {
	return errors.New("user is already registered")
}

func GetUserNotFoundError() error {
	return errors.New("nomor HP tidak ditemukan")
}

func GetPhoneNumberMinError() error {
	return errors.New("penulisan nomor HP minimal 10 digit. Silakan coba lagi")
}

func GetPhoneNumberMaxError() error {
	return errors.New("nomor HP melebihi batas maksimum karakter")
}

func GetPhoneNumberNotNumericError() error {
	return errors.New("nomor HP harus dalam karakter numerik 0-9" )
}

func GetCustomError(message string) error {
	return errors.New(message)
}
