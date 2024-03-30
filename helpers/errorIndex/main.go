package errorIndex

import (
	"errors"
	"strconv"
)

const baseDigit = 16

var (
	ErrInvalidToken        = errors.New("token invalid")
	ErrExpiredRefreshToken = errors.New("refresh token expired")
	ErrInvalidRefreshToken = errors.New("refresh token invalid")
	ErrGenerateJWTToken    = errors.New("error when generating token")
	ErrInvalidJWTAlg       = errors.New("invalid JWT algorithm")
	ErrSiginingToken       = errors.New("error when signing token")
	ErrParsingPrivateKey   = errors.New("error when parsing private key")
	ErrParsingPublicKey    = errors.New("error when parsing public key")

	ErrInvalidKeySize = errors.New("invalid key size")

	ErrQueryBuilder     = errors.New("error when building query")
	ErrHashingContent   = errors.New("error when hashing content")
	ErrPhoneNumberExist = errors.New("phone number already registered")
)

var (
	UserRegistrationError = errors.New(strconv.FormatUint(1, baseDigit))
	LoginError            = errors.New(strconv.FormatUint(2, baseDigit))
	UpdateProfileError    = errors.New(strconv.FormatUint(3, baseDigit))
	DevelopmentError      = errors.New(strconv.FormatUint(4, baseDigit))
	JWTModuleError        = errors.New(strconv.FormatUint(5, baseDigit))
	CryptoModuleError     = errors.New(strconv.FormatUint(6, baseDigit))
)
