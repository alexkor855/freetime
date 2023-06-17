package freetime

import (
	apmnt "app/internal/entities/appointment"
	idate "app/internal/entities/date"
	ivl "app/internal/entities/interval"
	sch "app/internal/entities/schedule"

	"errors"
)

type FreeTimeIntervals struct {
	branchId    string
	employeeId  string
	workplaceId string
	serviceId   string
	date        idate.Date
	intervals   ivl.Intervals
}

/*
* Initialize from Appointments and UnionDaySchedule
 */
func CreateFromScheduleAndAppointmentGroups(unionSchedule sch.UnionDaySchedule, appointmentsGroups []apmnt.Appointments) (*FreeTimeIntervals, error) {
	adi := &FreeTimeIntervals{
		branchId:    unionSchedule.BranchId(),
		employeeId:  unionSchedule.EmployeeId(),
		workplaceId: unionSchedule.WorkplaceId(),
		serviceId:   unionSchedule.ServiceId(),
		date:        unionSchedule.Date(),
		intervals:   unionSchedule.Intervals(),
	}

	for _, appointments := range appointmentsGroups {
		err := adi.AddAppointments(appointments)
		if err != nil {
			return nil, err
		}
	}

	return adi, nil
}

/*
* Initialize from intervals
 */
func CreateFromIntervals(
	branchId string,
	employeeId string,
	workplaceId string,
	serviceId string,
	date idate.Date,
	intervals ivl.Intervals,
) *FreeTimeIntervals {

	return &FreeTimeIntervals{
		branchId:    branchId,
		employeeId:  employeeId,
		workplaceId: workplaceId,
		serviceId:   serviceId,
		date:        date,
		intervals:   intervals,
	}
}

func (adi *FreeTimeIntervals) AddAppointments(a apmnt.Appointments) error {

	err := adi.validateAppointments(a)
	if err != nil {
		return err
	}

	adi.intervals, err = adi.intervals.Diff(a.Intervals())
	return err
}

func (adi *FreeTimeIntervals) validateAppointments(a apmnt.Appointments) error {

	if !adi.date.IsEqual(a.Date()) {
		return errors.New("FreeTimeIntervals and Appointments has different date")
	}

	if adi.branchId != a.BranchId() {
		return errors.New("FreeTimeIntervals and Appointments has different branchId")
	}

	if adi.employeeId != a.EmployeeId() && adi.workplaceId != a.WorkplaceId() { // can only match employeeId or workplaceId
		return errors.New("FreeTimeIntervals and Appointments has different employeeId and workplaceId")
	}
	// 	err = errors.Join(errors.New("FreeTimeIntervals and Appointments has different employeeId and workplaceId"))

	return nil
}

func (adi *FreeTimeIntervals) Intervals() []ivl.Interval {
	return adi.intervals.Get()
}

func (adi *FreeTimeIntervals) GetMaxInterval() ivl.Interval {

	max := adi.intervals.Get()[0]
	for _, invl := range adi.intervals.Get() {

		if invl.Duration() > max.Duration() {
			max = invl
		}
	}
	return max
}

// func (a *FreeTimeIntervals) GetMaxInterval() *Interval {
// }
