package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	test := []byte("hello world")

	res := HashSHA256(test)

	assert.NotEqual(t, test, res)
}

func TestSalt(t *testing.T) {
	assert.Panics(t, func() {
		// this should trigger error io procedure, and into panic
		GenerateSalt(-1)
	})

	salt := GenerateSalt(8)
	assert.Len(t, salt, 8)
}

func TestValidation(t *testing.T) {
	Pass := "HelloWorld"
	key := "AEIJAjrMxlELbwUnVHxAOq5SFYYRYUEP"
	salt := GenerateSalt(16)
	g := GCM(string(salt)).PlainText()
	g.Encrypt([]byte(key))
	encryptedSalt := g.String()

	expectedPass := HashSHA256(append([]byte(Pass), salt...))

	err := ValidatePassword(Pass, 16, []byte(key), expectedPass, []byte(encryptedSalt))
	if err != nil {
		panic(err)
	}
}
