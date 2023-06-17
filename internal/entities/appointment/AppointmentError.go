package apmnt

import (
	idate "app/internal/entities/date"
	"fmt"
)

type AppointmentsDataMismatchError struct {
	BranchId    string
	EmployeeId  string
	WorkplaceId string
	Date        idate.Date
	app         Appointment
}

func (e *AppointmentsDataMismatchError) Error() string {
	switch {
	case e.app.BranchId() != e.BranchId:
		return fmt.Sprintf("appointments has data mismatch in %s, %s, %s", "Appointment.BranchId", e.app.BranchId(), e.BranchId)
	case e.app.EmployeeId() != e.EmployeeId:
		return fmt.Sprintf("appointments has data mismatch in %s, %s, %s", "Appointment.EmployeeId", e.app.EmployeeId(), e.EmployeeId)
	case e.app.WorkplaceId() != e.WorkplaceId:
		return fmt.Sprintf("appointments has data mismatch in %s, %s, %s", "Appointment.WorkplaceId", e.app.WorkplaceId(), e.WorkplaceId)
	case e.app.Date() != e.Date:
		date := e.app.Date()
		return fmt.Sprintf("appointments has data mismatch in %s, %s, %s", "Appointment.Date", date.String(), e.Date.String())
	default:
		return "appointments has unknown data mismatch"
	}
}