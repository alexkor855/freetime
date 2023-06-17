package sch

type ScheduleType uint8

const (
	BranchSchedule ScheduleType = iota
	EmployeeSchedule
	WorkplaceSchedule
	EmployeeWorkplaceSchedule

	ServiceSchedule
	ServiceEmployeeSchedule
	ServiceWorkplaceSchedule
	ServiceEmployeeWorkplaceSchedule
)

// bit mask
func ScheduleTypeHasEmployee(sh ScheduleType) bool {
	return (sh & 1) != 0
}
func ScheduleTypeHasWorkplace(sh ScheduleType) bool {
	return (sh & 2) != 0
}
func ScheduleTypeHasService(sh ScheduleType) bool {
	return (sh & 4) != 0
}