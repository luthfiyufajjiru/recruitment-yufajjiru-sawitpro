package repository

import (
	"context"
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/helpers/errorIndex"
	"github.com/labstack/gommon/log"
)

func (r *Repository) GetProfile(ctx context.Context, user_id int) (profile generated.UserProfilePresenter, err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.Select(`"core"."name"`, `"core"."phone_number"`).Where("id = ?", user_id)

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
	panic("unimplemented")
}

func (r *Repository) UpdateProfile(ctx context.Context, user_id int, inp generated.UserProfilePresenter) (err error) {
	panic("unimplemented")
}

func (r *Repository) ComparePassword(ctx context.Context, phone_number string, password string) (err error) {
	panic("unimplemented")
}
