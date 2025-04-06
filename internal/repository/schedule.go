package repository

import "github.com/15011106/LetsWork/internal/domain"

type ScheduleRepository interface {
	Create(createdScheduleDto domain.CreateScheduleDto) error
	Delete(DeleteScheduleDto domain.DeleteScheduleDto) error
}
