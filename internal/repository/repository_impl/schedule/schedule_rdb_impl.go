package repository_impl

import (
	"database/sql"

	"github.com/15011106/LetsWork/internal/domain"
	"github.com/15011106/LetsWork/internal/repository"
)

type scheduleRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.ScheduleRepository {
	return &scheduleRepository{db: db}
}

func (s scheduleRepository) Create(createScheduleDto domain.CreateScheduleDto) error {
	return nil
}

func (s scheduleRepository) Delete(deleteScheduleDto domain.DeleteScheduleDto) error {
	return nil
}
