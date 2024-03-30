package helpers

import (
	"errors"
	"testing"

	"github.com/SawitProRecruitment/UserService/helpers/errorIndex"
	"github.com/stretchr/testify/assert"
)

// for testing, we are using same key for access token and refresh token. In staging/production we are using different key.
const keyStr = "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"

func TestParseKey(t *testing.T) {
	_, _, err := ParsePrivateKey(keyStr)
	assert.Nil(t, err)
}

func TestGenerateKey(t *testing.T) {
	Initialize(keyStr, keyStr, "10s", "10s")

	userId := 1
	userName := "foo"

	accToken, refToken, err := GenJWTTokens(userId, userName)
	assert.Nil(t, err)

	claims, err := GetClaims(accToken)
	assert.Nil(t, err)
	resUserId := int(claims["user_id"].(float64))
	assert.Equal(t, userId, resUserId)

	claims, err = GetClaims(refToken)
	assert.Nil(t, err)
	resUserId = int(claims["user_id"].(float64))
	assert.Equal(t, userId, resUserId)
}

func TestClaim(t *testing.T) {
	Initialize(keyStr, keyStr, "10s", "10s")

	userId := 1
	userName := "foo"

	accToken, _, err := GenJWTTokens(userId, userName)
	assert.Nil(t, err)

	claims, err := GetClaims(accToken)
	assert.Nil(t, err)
	resUserId := int(claims["user_id"].(float64))
	assert.Equal(t, userId, resUserId)

	_, err = GetClaims("invalidToken")
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, errorIndex.ErrInvalidToken))
}

func TestRefreshToken(t *testing.T) {
	Initialize(keyStr, keyStr, "10s", "10s")

	userId := 1
	userName := "foo"

	_, refToken, err := GenJWTTokens(userId, userName)
	assert.Nil(t, err)

	_, err = RefreshToken(refToken)
	assert.Nil(t, err)

	_, err = RefreshToken("refToken")
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, errorIndex.ErrInvalidRefreshToken))
}
