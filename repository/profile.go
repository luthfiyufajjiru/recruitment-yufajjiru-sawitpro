package repository

import (
	"context"
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/helpers"
	"github.com/SawitProRecruitment/UserService/helpers/errorIndex"
	"github.com/labstack/gommon/log"
	"github.com/lib/pq"
)

func (r *Repository) GetProfile(ctx context.Context, user_id int) (profile generated.UserProfilePresenter, err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.Select(`"core"."name"`, `"core"."phone_number"`).
		From(`"core"."users"`).Where("id = ?", user_id)

	queryStr, args, err := query.ToSql()
	if err != nil {
		log.Errorf("(%w) - %w", errorIndex.DevelopmentError, err)
		err = errorIndex.ErrQueryBuilder
		return
	}

	var user UserModel
	err = r.Db.GetContext(ctx, &user, queryStr, args...)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		err = nil
		return
	} else if err != nil {
		return
	}

	profile.Name = &user.Name
	profile.PhoneNumber = &user.PhoneNumber

	return
}

func (r *Repository) SetProfile(ctx context.Context, inp generated.UserRegistrationRequest) (resp generated.UserRegistrationResponse, err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	hashedPass, encryptedSalt, err := helpers.HashStringWithEncryptedSalt(inp.Password, r.SaltSize, r.SecretKey)
	if err != nil {
		log.Errorf("(%w) - %w", errorIndex.UserRegistrationError, err)
		err = errorIndex.ErrHashingContent
		return
	}

	query := psql.Insert(`"core"."users"`).
		Columns("name", "phone_number", "password_hash", "password_salt").
		Values(inp.Name, inp.PhoneNumber, hashedPass, encryptedSalt)

	queryStr, args, err := query.ToSql()
	if err != nil {
		log.Errorf("(%w) - %w", errorIndex.DevelopmentError, err)
		err = errorIndex.ErrQueryBuilder
		return
	}

	var pqErr *pq.Error
	res, err := r.Db.ExecContext(ctx, queryStr, args...)
	errors.As(err, &pqErr)

	if err != nil && pqErr != nil && pqErr.Code.Class() == IntegrityViolationClassCode {
		err = errorIndex.ErrPhoneNumberExist
		return
	} else if err != nil {
		return
	}

	_id, _ := res.LastInsertId()
	id := int(_id)

	resp.UserId = &id

	return
}

func (r *Repository) UpdateProfile(ctx context.Context, user_id int, inp generated.UserProfilePresenter) (err error) {
	if inp.PhoneNumber == nil && inp.Name == nil {
		err = errorIndex.ErrInvalidRequest
		return
	}

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query := psql.Update(`"core"."users"`)

	if inp.PhoneNumber != nil {
		query = query.Set("phone_number", inp.PhoneNumber)
	}

	if inp.Name != nil {
		query = query.Set("name", inp.Name)
	}

	query = query.Where(`"id" = ?`, user_id)

	queryStr, args, err := query.ToSql()
	if err != nil {
		log.Errorf("(%w) - %w", errorIndex.DevelopmentError, err)
		err = errorIndex.ErrQueryBuilder
		return
	}

	var pqErr *pq.Error
	_, err = r.Db.ExecContext(ctx, queryStr, args...)
	errors.As(err, &pqErr)
	if err != nil && pqErr != nil && pqErr.Code.Class() == IntegrityViolationClassCode {
		err = errorIndex.ErrPhoneNumberExist
		return
	} else if err != nil {
		log.Errorf(`(%w) - %w`, errorIndex.UpdateProfileError, err)
		return
	}

	return
}

func (r *Repository) ComparePassword(ctx context.Context, phone_number string, password string) (err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query := psql.Select("password_hash", "password_salt").From(`"core".users`).Where("phone_number = ?", phone_number)

	queryStr, args, err := query.ToSql()
	if err != nil {
		log.Errorf("(%w) - %w", errorIndex.DevelopmentError, err)
		err = errorIndex.ErrQueryBuilder
		return
	}

	var creds UserCredentialModel
	err = r.Db.GetContext(ctx, &creds, queryStr, args...)
	if err != nil {
		return
	}

	err = helpers.ValidatePassword(password, r.SaltSize, r.SecretKey, creds.PasswordHash, creds.PasswordSalt)

	return
}
