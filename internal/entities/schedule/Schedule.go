package sch

// import "time"
import (
	idate "app/internal/entities/date"
	"errors"
	"fmt"
)

var _ Schedule = &schedule{}

type schedule struct {
	id                 string
	scheduleTemplateId string
	scheduleType       ScheduleType
	branchId           string
	employeeId         string
	workplaceId        string
	serviceId          string
	startDate          idate.Date
	endDate            idate.Date
}

type Schedule interface {
	Id() string
	ScheduleTemplateId() string
	ScheduleType() ScheduleType
	BranchId() string
	EmployeeId() string
	WorkplaceId() string
	ServiceId() string
	StartDate() idate.Date
	EndDate() idate.Date
}

func NewSchedule(
	id string,
	scheduleTemplateId string,
	scheduleType ScheduleType,
	branchId string,
	employeeId string,
	workplaceId string,
	serviceId string,
	startDate idate.Date,
	endDate idate.Date,
) (*schedule, error) {
	sch := &schedule{
		id:                 id,
		scheduleTemplateId: scheduleTemplateId,
		scheduleType:       scheduleType,
		branchId:           branchId,
		employeeId:         employeeId,
		workplaceId:        workplaceId,
		serviceId:          serviceId,
		startDate:          startDate,
		endDate:            endDate,
	}

	if err := sch.validate(); err != nil {
		return nil, err
	}

	return sch, nil
}

func (s *schedule) validate() error {

	var err error

	if s.branchId == "" {
		err = errors.Join(errors.New("Schedule must contain a branch"))
	}

	if s.startDate == nil {
		err = errors.Join(errors.New("Schedule must contain a startDate"))
	}

	if s.startDate != nil && s.endDate != nil && !s.startDate.IsEarlierOrEqual(s.endDate) {
		err = errors.Join(fmt.Errorf("Schedule startDate (%s) must be earlier or equal endDate (%s)", s.startDate, s.endDate))
	}

	switch {
	case s.scheduleTypeHasEmployee() && s.employeeId == "":
		err = errors.Join(errors.New("Schedule must contain a employee"))

	case s.scheduleTypeHasWorkplace() && s.workplaceId == "":
		err = errors.Join(errors.New("Schedule must contain a workplace"))

	case s.scheduleTypeHasService() && s.serviceId == "":
		err = errors.Join(errors.New("Schedule must contain a service"))
	}

	switch {
	case !s.scheduleTypeHasEmployee() && s.employeeId != "":
		err = errors.Join(errors.New("Schedule must not contain an employee"))

	case !s.scheduleTypeHasWorkplace() && s.workplaceId != "":
		err = errors.Join(errors.New("Schedule must not contain a workplace"))

	case !s.scheduleTypeHasService() && s.serviceId != "":
		err = errors.Join(errors.New("Schedule must not contain a service"))
	}

	return err
}

func (s *schedule) scheduleTypeHasEmployee() bool {
	return (s.scheduleType & 1) != 0
}
func (s *schedule) scheduleTypeHasWorkplace() bool {
	return (s.scheduleType & 2) != 0
}
func (s *schedule) scheduleTypeHasService() bool {
	return (s.scheduleType & 4) != 0
}

func (s *schedule) Id() string {
	return s.id
}

func (s *schedule) ScheduleTemplateId() string {
	return s.scheduleTemplateId
}

func (s *schedule) ScheduleType() ScheduleType {
	return s.scheduleType
}

func (s *schedule) BranchId() string {
	return s.branchId
}

func (s *schedule) EmployeeId() string {
	return s.employeeId
}

func (s *schedule) WorkplaceId() string {
	return s.workplaceId
}

func (s *schedule) ServiceId() string {
	return s.serviceId
}

func (s *schedule) StartDate() idate.Date {
	return s.startDate
}

func (s *schedule) EndDate() idate.Date {
	return s.endDate
}
