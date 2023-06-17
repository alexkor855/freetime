package getftime

import (
	"app/internal/usecases"
	getapmnt "app/internal/usecases/appointment"
	getsch "app/internal/usecases/schedule"

	apmnt "app/internal/entities/appointment"
	"app/internal/entities/freetime"
	ivl "app/internal/entities/interval"
	sch "app/internal/entities/schedule"
)

var _ GetFreeTimeService = &getFreeTimeService{}

type getFreeTimeService struct {
	scheduleRepository    getsch.GetScheduleRepository
	appointmentRepository getapmnt.GetAppointmentRepository
}

type GetFreeTimeService interface {
	GetFreeTime(params usecases.GetRequestDTO) ([]ivl.Interval, error)
}

func (s *getFreeTimeService) GetFreeTime(params usecases.GetRequestDTO) ([]ivl.Interval, error) {
	emptyRes := make([]ivl.Interval, 0)

	// get []schedule from DB or cache
	schedules, err := s.scheduleRepository.GetAllByParams(params)
	if err != nil {
		// logger
		return nil, err
	}
	if len(schedules) == 0 {
		return emptyRes, nil
	}

	// make UnionSchedule
	uds, err := sch.NewUnionDaySchedule(schedules)
	if err != nil {
		return nil, err
	}

	// get []appointment from DB or cache
	appointments, err := s.appointmentRepository.GetAllByParams(params)
	if err != nil {
		return nil, err
	}
	// make Appointments
	appointmentsGroup, err := apmnt.NewAppointmentsGroup(appointments)
	if err != nil {
		return nil, err
	}

	// make FreeTimeIntervals
	ft, err := freetime.CreateFromScheduleAndAppointmentGroups(uds, appointmentsGroup)
	if err != nil {
		return nil, err
	}

	return ft.Intervals(), nil
}
