package helpers

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"time"

	"github.com/SawitProRecruitment/UserService/helpers/errorIndex"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/gommon/log"
)

var (
	accessTimeDuration, refreshTimeDuration       time.Duration
	privateKeyAccessToken, privateKeyRefreshToken *rsa.PrivateKey
	publicKeyAccessToken                          *rsa.PublicKey
)

func ParsePrivateKey(storedPrivKeyStr string) (privKey *rsa.PrivateKey, pubKey *rsa.PublicKey, err error) {
	pemString := fmt.Sprintf("-----BEGIN RSA PRIVATE KEY-----\n%s\n-----END RSA PRIVATE KEY-----", storedPrivKeyStr)
	block, _ := pem.Decode([]byte(pemString))
	privKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Error("(%w) - %w. Error parsing private key.", errorIndex.DevelopmentError, err)
		return nil, nil, errorIndex.ErrParsingPrivateKey
	}
	pubKey = &privKey.PublicKey
	return
}

func Initialize(privateKeyAccessTokenStr, privateKeyRefreshTokenStr, accessTimeoutStr, refreshDurationStr string) {
	if val, err := time.ParseDuration(accessTimeoutStr); err == nil {
		accessTimeDuration = val
	} else if err != nil {
		panic("invalid jwt access time duration env")
	}

	if val, err := time.ParseDuration(refreshDurationStr); err == nil {
		refreshTimeDuration = val
	} else if err != nil {
		panic("invalid jwt refresh time duration env")
	}

	var err error
	privateKeyAccessToken, publicKeyAccessToken, err = ParsePrivateKey(privateKeyAccessTokenStr)
	if err != nil {
		panic(err)
	}

	privateKeyRefreshToken, _, err = ParsePrivateKey(privateKeyRefreshTokenStr)
	if err != nil {
		panic(err)
	}
}

func SignJWTToken(cl jwt.MapClaims, key *rsa.PrivateKey) (tokenStr string, err error) {
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS256, cl)
	tokenStr, err = tkn.SignedString(key)
	return
}

func GenJWTTokens(user_id int, name string) (accessTokenStr, refreshTokenStr string, err error) {
	accessClaims := jwt.MapClaims{
		"name":       name,
		"user_id":    user_id,
		"expired_at": time.Now().Add(accessTimeDuration).UnixMilli(),
	}

	refreshClaims := jwt.MapClaims{
		"name":       name,
		"user_id":    user_id,
		"expired_at": time.Now().Add(refreshTimeDuration).UnixMilli(),
	}

	accessTokenStr, err = SignJWTToken(accessClaims, privateKeyAccessToken)
	if err != nil {
		log.Errorf("(%w) - %w. Error when signing access token", errorIndex.JWTModuleError, err)
		err = errorIndex.ErrSiginingToken
		return
	}

	refreshTokenStr, err = SignJWTToken(refreshClaims, privateKeyRefreshToken)
	if err != nil {
		log.Errorf("(%w) - %w. Error when signing refresh token", errorIndex.JWTModuleError, err)
		err = errorIndex.ErrSiginingToken
	}

	return
}

func getClaims(tokenStr string, key *rsa.PublicKey) (claims jwt.MapClaims, err error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errorIndex.ErrInvalidJWTAlg
		}
		return key, nil
	})
	if algErr := errors.Is(err, errorIndex.ErrInvalidJWTAlg); err != nil && algErr {
		return
	} else if err != nil && !algErr {
		log.Error("(%w) - %w", errorIndex.JWTModuleError, err)
		err = errorIndex.ErrInvalidToken
		return
	}
	claims, _ = token.Claims.(jwt.MapClaims)
	return
}

func GetClaims(tokenStr string) (claims jwt.MapClaims, err error) {
	claims, err = getClaims(tokenStr, publicKeyAccessToken)
	return
}

func RefreshToken(token string) (accesstoken string, err error) {
	claims, err := getClaims(token, publicKeyAccessToken)
	if errors.Is(err, errorIndex.ErrInvalidJWTAlg) {
		err = errorIndex.ErrInvalidRefreshToken
		return
	}

	expired, ok := claims["expired_at"].(float64)
	if !ok {
		log.Errorf("(%w) - refresh token expired_at payload was invalid")
		err = errorIndex.ErrInvalidRefreshToken
		return
	}

	exp := int64(expired)

	if exp < time.Now().UnixMilli() {
		err = errorIndex.ErrExpiredRefreshToken
		return
	}

	accesstoken, err = SignJWTToken(jwt.MapClaims{
		"name":       claims["name"],
		"user_id":    claims["user_id"],
		"expired_at": time.Now().Add(accessTimeDuration).UnixMilli(),
	}, privateKeyAccessToken)

	if err != nil {
		log.Errorf("(%w) - %w", errorIndex.JWTModuleError, err)
		err = errorIndex.ErrGenerateJWTToken
		return
	}

	return
}
