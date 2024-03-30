package helpers

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"io"

	"github.com/SawitProRecruitment/UserService/helpers/errorIndex"
)

func GenerateSalt(size int) []byte {
	salt := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		panic(err)
	}
	return salt
}

// hashSHA256 hashes the input data using SHA-256 and returns the hash
func HashSHA256(data []byte) []byte {
	hasher := sha256.New()
	hasher.Write(data)
	return hasher.Sum(nil)
}

func HashStringWithEncryptedSalt(msg string, saltSize int, key []byte) (hashedMsg, encryptedSalt []byte, err error) {
	salt := GenerateSalt(saltSize)
	_encryptedSalt := GCM(string(salt)).PlainText()
	err = _encryptedSalt.Encrypt(key)
	if err != nil {
		return
	}
	encryptedSalt = []byte(_encryptedSalt.String())
	hashedMsg = HashSHA256(append([]byte(msg), []byte(salt)...))
	return
}

func ValidatePassword(password string, saltSize int, key, hashedPass, encryptedSalt []byte) (err error) {
	CipherSalt := GCM(string(encryptedSalt)).CipherText()
	err = CipherSalt.Decrypt(key)
	if err != nil {
		err = errorIndex.ErrInvalidPassword
		return
	}
	Salt := CipherSalt.String()
	d := HashSHA256(append([]byte(password), []byte(Salt)...))
	if subtle.ConstantTimeCompare(d, hashedPass) == 0 {
		err = errorIndex.ErrInvalidPassword
		return
	}
	return
}
