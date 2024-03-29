package repository

import "context"

// ComparePassword implements RepositoryInterface.
func (r *Repository) ComparePassword(ctx context.Context, phone_number string, password string) (err error) {
	panic("unimplemented")
}

// GetProfile implements RepositoryInterface.
func (r *Repository) GetProfile(ctx context.Context, user_id int) (err error) {
	panic("unimplemented")
}

// UpdateProfile implements RepositoryInterface.
func (r *Repository) UpdateProfile(ctx context.Context, user_id int) (err error) {
	panic("unimplemented")
}
