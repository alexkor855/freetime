package getapmnt

import (
	apmnt "app/internal/entities/appointment"
	"app/internal/usecases"
)

type GetAppointmentRepository interface {
	GetAllByParams(params usecases.GetRequestDTO) ([]apmnt.Appointment, error)
}
