package repository

import "github.com/golang-jwt/jwt/v5"

// GenJWTTokens implements RepositoryInterface.
func (r *Repository) GenJWTTokens(user_id int, name string) (accessTokenStr string, refreshTokenStr string, err error) {
	panic("unimplemented")
}

// GetClaims implements RepositoryInterface.
func (r *Repository) GetClaims(tokenStr string, key []byte) (claims jwt.MapClaims, valid bool) {
	panic("unimplemented")
}

// RefreshToken implements RepositoryInterface.
func (r *Repository) RefreshToken(token string) (accesstoken string, err error) {
	panic("unimplemented")
}

// SignJWTToken implements RepositoryInterface.
func (r *Repository) SignJWTToken(cl jwt.MapClaims, key []byte) (tokenStr string, err error) {
	panic("unimplemented")
}
