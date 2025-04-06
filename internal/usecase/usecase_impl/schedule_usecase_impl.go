package usecase_impl

import (
	"github.com/15011106/LetsWork/internal/domain"
	"github.com/15011106/LetsWork/internal/repository"
)

type scheduleUsecase struct {
	scheduleRepository repository.ScheduleRepository
}

func NewScheduleUsecase(scheduleRepository repository.ScheduleRepository) *scheduleUsecase {
	return &scheduleUsecase{scheduleRepository: scheduleRepository}
}

func (s *scheduleUsecase) Create(createScheduleDto domain.CreateScheduleDto) error {
	return s.scheduleRepository.Create(createScheduleDto)
}

func (s *scheduleUsecase) Delete(deleteSchduleDto domain.DeleteScheduleDto) error {
	return s.scheduleRepository.Delete(deleteSchduleDto)
}
