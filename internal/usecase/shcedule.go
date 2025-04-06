package usecase

import "github.com/15011106/LetsWork/internal/domain"

type ScheduleUsecase interface {
	Create(schedule *domain.Schedule) error
	Delete(userID *int64) error
}
