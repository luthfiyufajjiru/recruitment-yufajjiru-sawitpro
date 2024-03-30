package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/SawitProRecruitment/UserService/helpers/errorIndex"
)

type (
	Encryption interface {
		PlainText() PlainText
		CipherText() CipherText
	}

	PlainText interface {
		Encrypt([]byte) error
		AsCipherText() CipherText
		String() string
	}

	CipherText interface {
		Decrypt([]byte) error
		AsPlainText() PlainText
		String() string
	}

	gcmString string

	gcmPlainText string

	gcmCipherText string
)

func GCM(inp string) Encryption {
	c := gcmString(inp)
	return &c
}

func (k *gcmString) PlainText() PlainText {
	c := string(*k)
	p := gcmPlainText(c)
	return &p
}

func (k *gcmString) CipherText() CipherText {
	c := string(*k)
	p := gcmCipherText(c)
	return &p
}

func (e *gcmPlainText) Encrypt(key []byte) (err error) {
	if len(key) != 32 {
		err = errorIndex.ErrInvalidKeySize
		return
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		return
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return
	}

	*e = gcmPlainText(gcm.Seal(nonce, nonce, []byte(string(*e)), nil))
	return
}

func (e *gcmPlainText) String() string {
	return string(*e)
}

func (e *gcmPlainText) AsCipherText() CipherText {
	s := e.String()
	v := gcmCipherText(s)
	return &v
}

func (e *gcmCipherText) Decrypt(key []byte) (err error) {
	if len(key) != 32 {
		err = errorIndex.ErrInvalidKeySize
		return
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		return
	}

	ciphertext := []byte(string(*e))
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return
	}

	nonce := ciphertext[:nonceSize]
	ciphertext = ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return
	}

	*e = gcmCipherText(plaintext)

	return err
}

func (e *gcmCipherText) String() string {
	return string(*e)
}

func (e *gcmCipherText) AsPlainText() PlainText {
	s := e.String()
	v := gcmPlainText(s)
	return &v
}
