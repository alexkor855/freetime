package getsch

import (
	sch "app/internal/entities/schedule"
	"app/internal/usecases"
)

var _ GetScheduleService = &getScheduleService{}

type getScheduleService struct {
	repository GetScheduleRepository
}

type GetScheduleService interface {
	GetAllByParams(params usecases.GetRequestDTO) ([]sch.DaySchedule, error)
}

func (s *getScheduleService) GetAllByParams(params usecases.GetRequestDTO) ([]sch.DaySchedule, error) {
	res, err := s.repository.GetAllByParams(params)

	if err != nil {
		return nil, err
	}

	return res, nil
}
