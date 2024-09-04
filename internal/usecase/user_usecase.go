package usecase

import (
	"context"
	"fmt"

	"github.com/Kurichi/go-template/internal/domain/model"
)

type userUsecase struct {
}

func NewUserUsecase() UserUsecase {
	return &userUsecase{}
}

// CreateUser implements UserUsecase.
func (u *userUsecase) CreateUser(ctx context.Context, req *User) (*User, error) {
	user, err := model.NewUser(req.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &User{
		ID:   user.ID.String(),
		Name: user.Name,
	}, nil
}

// GetUser implements UserUsecase.
func (u *userUsecase) GetUser(ctx context.Context, id string) (*User, error) {
	panic("unimplemented")
}
