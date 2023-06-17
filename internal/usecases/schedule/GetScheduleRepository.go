package getsch

import (
	sch "app/internal/entities/schedule"
	"app/internal/usecases"
)

type GetScheduleRepository interface {
	GetAllByParams(params usecases.GetRequestDTO) ([]sch.DaySchedule, error)
}
