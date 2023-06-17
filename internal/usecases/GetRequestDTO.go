package usecases

import (
	idate "app/internal/entities/date"
)

type GetRequestDTO struct {
	Date        idate.Date
	BranchId    string
	EmployeeId  string
	WorkplaceId string
	ServiceId   string
}
