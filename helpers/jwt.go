package helpers

import (
	"github.com/golang-jwt/jwt/v5"
)

func SignJWTToken(cl jwt.MapClaims, key []byte) (tokenStr string, err error) {
	panic("unimplemented")
}

func GenJWTTokens(user_id int, name string) (accessTokenStr, refreshTokenStr string, err error) {
	panic("unimplemented")
}

func GetClaims(tokenStr string, key []byte) (claims jwt.MapClaims, valid bool) {
	panic("unimplemented")
}

func RefreshToken(token string) (accesstoken string, err error) {
	panic("unimplemented")
}
