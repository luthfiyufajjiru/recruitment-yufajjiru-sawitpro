package repository

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
)

type RepositoryInterface interface {
	UpdateProfile(ctx context.Context, user_id int) (err error)
	GetProfile(ctx context.Context, user_id int) (err error)
	ComparePassword(ctx context.Context, phone_number, password string) (err error)
	SignJWTToken(cl jwt.MapClaims, key []byte) (tokenStr string, err error)
	GenJWTTokens(user_id int, name string) (accessTokenStr, refreshTokenStr string, err error)
	GetClaims(tokenStr string, key []byte) (claims jwt.MapClaims, valid bool)
	RefreshToken(token string) (accesstoken string, err error)
}
