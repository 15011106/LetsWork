package repository_impl

import (
	"database/sql"

	"github.com/15011106/LetsWork/internal/domain"
	"github.com/15011106/LetsWork/internal/repository"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (u userRepository) Create(user *domain.User) error {
	return nil
}

func (u userRepository) Delete(userId *int64) error {
	return nil
}
