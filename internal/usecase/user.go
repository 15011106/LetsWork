package usecase

import "github.com/15011106/LetsWork/internal/domain"

type UserUsecase interface {
	Register(user *domain.User) error
	Delete(userID *int64) error
}
