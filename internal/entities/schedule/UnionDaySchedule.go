package sch

import (
	idate "app/internal/entities/date"
	ivl "app/internal/entities/interval"
	"errors"
)

var _ UnionDaySchedule = &unionDaySchedule{}

type unionDaySchedule struct {
	isInit bool

	scheduleTypes map[ScheduleType]bool
	branchId      string
	employeeId    string
	workplaceId   string
	serviceId     string
	date          idate.Date
	DaySchedules  []DaySchedule
	intervals     ivl.Intervals
}

type UnionDaySchedule interface {
	BranchId() string
	EmployeeId() string
	WorkplaceId() string
	ServiceId() string
	Date() idate.Date

	GetUnionSchedule() ([]ivl.Interval, error)
	HasUnionSchedule() (bool, error)
	Intervals() ivl.Intervals
}

/*
** Init
 */
func NewUnionDaySchedule(daySchedules []DaySchedule) (UnionDaySchedule, error) {
	u := &unionDaySchedule{}

	if len(daySchedules) == 0 {
		u.isInit = true
	}

	for _, daySchedule := range daySchedules {
		err := u.addSchedule(daySchedule)

		if err != nil {
			return nil, err
		}
	}
	return u, nil
}

func (u *unionDaySchedule) BranchId() string {
	return u.branchId
}

func (u *unionDaySchedule) EmployeeId() string {
	return u.employeeId
}

func (u *unionDaySchedule) WorkplaceId() string {
	return u.workplaceId
}

func (u *unionDaySchedule) ServiceId() string {
	return u.serviceId
}

func (u *unionDaySchedule) Date() idate.Date {
	return u.date
}

/*
 */
func (u *unionDaySchedule) addSchedule(daySchedule DaySchedule) error {
	var err error

	// Set params from first DaySchedule as main for UnionDaySchedule
	// the others DaySchedules must have the same parameters
	if !u.isInit {
		u.date = daySchedule.Date()
		u.branchId = daySchedule.Schedule().Id()
		u.intervals = daySchedule.Intervals()
	} else {
		err := u.validateDaySchedule(daySchedule)
		if err != nil {
			return err
		}
	}

	u.scheduleTypes[daySchedule.Schedule().ScheduleType()] = true

	if u.isInit && len(u.intervals.Get()) > 0 {

		u.intervals, err = u.intervals.Intersect(daySchedule.Intervals())
		if err != nil {
			return err
		}
	}

	if !u.isInit {
		u.isInit = true
	}

	return nil
}

func (u *unionDaySchedule) validateDaySchedule(daySchedule DaySchedule) error {

	d := daySchedule.Date()
	if !u.date.IsEqual(d) {
		return errors.New("UnionDaySchedule has different date")
	}

	dayScheduleType := daySchedule.Schedule().ScheduleType()

	if u.scheduleTypes[dayScheduleType] {
		return errors.New("DaySchedule has the same schedule type")
	}

	if u.branchId != daySchedule.Schedule().BranchId() {
		return errors.New("UnionDaySchedule has different branchId")
	}

	switch {
	case ScheduleTypeHasService(dayScheduleType) && u.serviceId != "" && u.serviceId != daySchedule.Schedule().ServiceId():
		return errors.New("UnionDaySchedule contains another service")

	case ScheduleTypeHasEmployee(dayScheduleType) && u.employeeId != "" && u.employeeId != daySchedule.Schedule().EmployeeId():
		return errors.New("UnionDaySchedule contains another employee")

	case ScheduleTypeHasWorkplace(dayScheduleType) && u.workplaceId != "" && u.workplaceId != daySchedule.Schedule().WorkplaceId():
		return errors.New("UnionDaySchedule contains another workplace")
	}

	if ScheduleTypeHasService(dayScheduleType) && u.serviceId == "" {
		u.serviceId = daySchedule.Schedule().ServiceId()
	}

	if ScheduleTypeHasEmployee(dayScheduleType) && u.employeeId == "" {
		u.employeeId = daySchedule.Schedule().EmployeeId()
	}

	if ScheduleTypeHasWorkplace(dayScheduleType) && u.workplaceId == "" {
		u.workplaceId = daySchedule.Schedule().WorkplaceId()
	}

	return nil
}

/*
**
 */
func (u *unionDaySchedule) GetUnionSchedule() ([]ivl.Interval, error) {
	if !u.isInit {
		return nil, errors.New("UnionDaySchedule has not been initialised")
	}

	return u.intervals.Get(), nil
}

/*
**
 */
func (u *unionDaySchedule) HasUnionSchedule() (bool, error) {
	if !u.isInit {
		return false, errors.New("UnionDaySchedule has not been initialised")
	}

	return len(u.intervals.Get()) > 0, nil
}

func (u *unionDaySchedule) Intervals() ivl.Intervals {

	return u.intervals
}
