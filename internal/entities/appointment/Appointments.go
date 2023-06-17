package apmnt

import (
	idate "app/internal/entities/date"
	ivl "app/internal/entities/interval"
	ie "app/internal/errors"
	"strings"
)

var _ Appointments = &appointments{}

type appointments struct {
	list []Appointment

	branchId    string
	employeeId  string
	workplaceId string
	date        idate.Date
	intervals   ivl.Intervals
}

type Appointments interface {
	List() []Appointment
	BranchId() string
	EmployeeId() string
	WorkplaceId() string
	Date() idate.Date
	Intervals() ivl.Intervals
}

func (a *appointments) BranchId() string {
	return a.branchId
}

func (a *appointments) EmployeeId() string {
	return a.employeeId
}

func (a *appointments) WorkplaceId() string {
	return a.workplaceId
}

func (a *appointments) Date() idate.Date {
	return a.date
}

func (a *appointments) Intervals() ivl.Intervals {
	return a.intervals
}

func (a *appointments) List() []Appointment {
	return a.list
}

/*
* Create group of appointment sets
 */
func NewAppointmentsGroup(appointments []Appointment) ([]Appointments, error) {
	result := make([]Appointments, 0)
	groups := make(map[string][]Appointment)

	for _, app := range appointments {

		groupKey := strings.Join([]string{app.Date().String(), app.BranchId(), app.EmployeeId(), app.WorkplaceId()}, "|")

		if _, ok := groups[groupKey]; !ok {
			groups[groupKey] = make([]Appointment, 0)
		}
		groups[groupKey] = append(groups[groupKey], app)
	}

	for _, groupAppointments := range groups {
		apps, err := NewAppointments(groupAppointments)
		if err != nil {
			return nil, err
		}
		result = append(result, apps)
	}
	return result, nil
}

/*
* Create appointment sets
* Appointments must relate to same employeeId, workplaceId or employeeId + workplaceId, so appointments intervals can NOT intersection
 */
func NewAppointments(appointmentList []Appointment) (*appointments, error) {
	if len(appointmentList) == 0 {
		return nil, &ie.EmptyParamError{FuncName: "Appointments.NewAppointments", ParamName: "appointments"}
	} // return nil, nil

	invls := make([]ivl.Interval, 0)
	invls = append(invls, appointmentList[0].Interval())
	branchId := appointmentList[0].BranchId()
	employeeId := appointmentList[0].EmployeeId()
	workplaceId := appointmentList[0].WorkplaceId()
	date := appointmentList[0].Date()

	for i := 1; i < len(appointmentList); i++ {
		app := appointmentList[i]
		if app.BranchId() != branchId || app.EmployeeId() != employeeId || app.WorkplaceId() != workplaceId || app.Date() != date {
			return nil, &AppointmentsDataMismatchError{branchId, employeeId, workplaceId, date, app}
		}
		invls = append(invls, app.Interval())
	}

	invlRes, err := ivl.NewIntervals(invls, ivl.OverlapProhibited)
	if err != nil {
		return nil, err
	}

	apps := &appointments{
		list:        appointmentList,
		branchId:    branchId,
		employeeId:  employeeId,
		workplaceId: workplaceId,
		date:        date,
		intervals:   invlRes,
	}
	return apps, nil
}
