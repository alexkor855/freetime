package ftc

type GetFreeTimeRequest struct {
	Date        string `query:"date"`
	BranchId    string `query:"branchId"`
	EmployeeId  string `query:"employeeId"`
	WorkplaceId string `query:"workplaceId"`
	ServiceId   string `query:"serviceId"`
}
