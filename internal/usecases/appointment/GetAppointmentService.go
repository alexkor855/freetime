package getapmnt

import (
	apmnt "app/internal/entities/appointment"
	"app/internal/usecases"
)

var _ GetAppointmentService = &getAppointmentService{}

type getAppointmentService struct {
	repository GetAppointmentRepository
}

type GetAppointmentService interface {
	GetAllByParams(params usecases.GetRequestDTO) ([]apmnt.Appointment, error)
}

func (s *getAppointmentService) GetAllByParams(params usecases.GetRequestDTO) ([]apmnt.Appointment, error) {
	res, err := s.repository.GetAllByParams(params)

	if err != nil {
		return nil, err
	}

	return res, nil
}
