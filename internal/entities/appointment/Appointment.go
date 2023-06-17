package apmnt

import (
	idate "app/internal/entities/date"
	ivl "app/internal/entities/interval"
)

var _ Appointment = &appointment{}

type appointment struct {
	id          string
	branchId    string
	employeeId  string
	workplaceId string
	serviceId   string
	status      AppointmentStatus
	customerId  string
	date        idate.Date
	interval	ivl.Interval
	
	//CreatedAt  time.Time
	//UpdatedAt  time.Time
}

type Appointment interface {
	Id() string
	BranchId() string
	EmployeeId() string
	WorkplaceId() string
	ServiceId() string
	Status() AppointmentStatus
	CustomerId() string
	Date() idate.Date
	Interval() ivl.Interval
}

func (a *appointment) Id() string {
	return a.id
}

func (a *appointment) BranchId() string {
	return a.branchId
}

func (a *appointment) EmployeeId() string {
	return a.employeeId
}

func (a *appointment) WorkplaceId() string {
	return a.workplaceId
}

func (a *appointment) ServiceId() string {
	return a.serviceId
}

func (a *appointment) Status() AppointmentStatus {
	return a.status
}

func (a *appointment) CustomerId() string {
	return a.customerId
}

func (a *appointment) Date() idate.Date {
	return a.date
}

func (a *appointment) Interval() ivl.Interval {
	return a.interval
}