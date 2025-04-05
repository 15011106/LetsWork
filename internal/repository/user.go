package repository

import "github.com/15011106/LetsWork/internal/domain"

type UserRepository interface {
	Create(user *domain.User) error
	Delete(userID *int64) error
}
