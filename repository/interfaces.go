package repository

import (
	"context"

	"github.com/SawitProRecruitment/UserService/generated"
)

type RepositoryInterface interface {
	GetProfile(ctx context.Context, user_id int) (profile generated.UserProfilePresenter, err error)
	SetProfile(ctx context.Context, inp generated.UserRegistrationRequest) (resp generated.UserRegistrationResponse, err error)
	UpdateProfile(ctx context.Context, user_id int) (err error)
	ComparePassword(ctx context.Context, phone_number, password string) (err error)
}
