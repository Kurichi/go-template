package usecase

import "context"

type UserUsecase interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUser(ctx context.Context, id string) (*User, error)
}
