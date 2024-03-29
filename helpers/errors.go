package helpers

import "strconv"

type (
	ErrorCodes string
)

const baseDigit = 16

func (ec ErrorCodes) UserRegistrationError() string {
	return strconv.FormatUint(1, baseDigit)
}

func (ec ErrorCodes) LoginError() string {
	return strconv.FormatUint(2, baseDigit)
}

func (ec ErrorCodes) UpdateProfileError() string {
	return strconv.FormatUint(3, baseDigit)
}
