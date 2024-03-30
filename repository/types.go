// This file contains types that are used in the repository layer.
package repository

type (
	UserModel struct {
		Id          int    `db:"id"`
		Name        string `db:"name"`
		PhoneNumber string `db:"phone_number"`
	}

	UserCredentialModel struct {
		Id           int    `db:"id"`
		PasswordHash []byte `db:"password_hash"`
		PasswordSalt []byte `db:"password_salt"`
	}
)
