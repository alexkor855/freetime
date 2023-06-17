package sch

import (
	idate "app/internal/entities/date"
	ivl "app/internal/entities/interval"
)

var _ DaySchedule = &daySchedule{}

type daySchedule struct {
	id        string
	schedule  Schedule
	date      idate.Date
	intervals ivl.Intervals
}

type DaySchedule interface {
	Schedule() Schedule
	Date() idate.Date
	Intervals() ivl.Intervals
}

func NewDaySchedule(
	id string,
	schedule Schedule,
	date idate.Date,
	intervals ivl.Intervals,
) (DaySchedule, error) {

	d := &daySchedule{}
	d.id = id
	d.schedule = schedule
	d.date = date
	d.intervals = intervals
	return d, nil
}

func (d *daySchedule) Schedule() Schedule {
	return d.schedule
}

func (d *daySchedule) Date() idate.Date {
	return d.date
}

func (d *daySchedule) Intervals() ivl.Intervals {
	return d.intervals
}
