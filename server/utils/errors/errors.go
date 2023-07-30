package errors

import (
	"errors"
	"fmt"
)

var errBadRequest = errors.New("bad request error")

func NewBadRequestError(msg string) error {
	return WrapBadRequestError(errors.New(msg))
}

func WrapBadRequestError(err error) error {
	return fmt.Errorf("%w: %s", errBadRequest, err.Error())
}

func IsBadRequestError(err error) bool {
	return errors.Is(err, errBadRequest)
}

var errInvalidCredentials = errors.New("invalid credentials")

func NewInvalidCredentialsError(msg string) error {
	return WrapInvalidCredentialsError(errors.New(msg))
}

func WrapInvalidCredentialsError(err error) error {
	return fmt.Errorf("%w: %s", errInvalidCredentials, err.Error())
}

func IsInvalidCredentialsError(err error) bool {
	return errors.Is(err, errInvalidCredentials)
}
