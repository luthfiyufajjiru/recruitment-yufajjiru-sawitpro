package helpers

import (
	"errors"
	"strconv"
)

type (
	ErrorCodes string
)

const baseDigit = 16

func (ec ErrorCodes) UserRegistrationError() error {
	return errors.New(strconv.FormatUint(1, baseDigit))
}

func (ec ErrorCodes) LoginError() error {
	return errors.New(strconv.FormatUint(2, baseDigit))
}

func (ec ErrorCodes) UpdateProfileError() error {
	return errors.New(strconv.FormatUint(3, baseDigit))
}
