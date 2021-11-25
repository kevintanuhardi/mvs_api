package constants

import "errors"

func GetErrDatabaseError() error {
	return errors.New("3001||database Error")
}