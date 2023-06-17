package ivl

// var _ ScheduleInterval = &scheduleInterval{}

type ScheduleInterval struct {
	id        string
	interval
}

// type ScheduleInterval interface {
// 	*Interval
// 	ID() string
// }

func (si *ScheduleInterval) ID() string {
	//si.Interval.StartTime()
	si.StartTime()
	return si.id
}
