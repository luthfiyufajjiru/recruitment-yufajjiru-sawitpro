package repository

import (
	"context"

	"github.com/SawitProRecruitment/UserService/generated"
)

func (r *Repository) GetProfile(ctx context.Context, user_id int) (profile generated.UserProfilePresenter, err error) {
	panic("unimplemented")
}

func (r *Repository) SetProfile(ctx context.Context, inp generated.UserRegistrationRequest) (resp generated.UserRegistrationResponse, err error) {
	panic("unimplemented")
}

func (r *Repository) UpdateProfile(ctx context.Context, user_id int) (err error) {
	panic("unimplemented")
}

func (r *Repository) ComparePassword(ctx context.Context, phone_number string, password string) (err error) {
	panic("unimplemented")
}
