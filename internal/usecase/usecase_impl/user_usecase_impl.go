package usecase_impl

import (
	"github.com/15011106/LetsWork/internal/domain"
	"github.com/15011106/LetsWork/internal/repository"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepository}
}

func (u *userUsecase) Register(user *domain.User) error {
	return u.userRepository.Create(user)
}

func (u *userUsecase) Delete(userID *int64) error {
	return u.userRepository.Delete(userID)
}
