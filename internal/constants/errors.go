package constants

import "errors"

func GetErrDatabaseError() error {
	return errors.New("3001||database Error")
}

func GetDuplicateUserError() error {
	return errors.New("user is already registered")
}

func GetCustomError(message string) error {
	return errors.New(message)
}
