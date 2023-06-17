package ivl

import (
	"errors"
	"time"
	//"fmt"
	//"math"
)

var _ Interval = &interval{}
var EmptyInterval Interval = &interval{}

const IntervalTimeFormat string = "15:04"

type interval struct {
	startTime time.Time
	endTime   time.Time
}

type Interval interface {
	StartTime() time.Time
	EndTime() time.Time
	IsBefore(c Interval) bool
	IsAfter(c Interval) bool
	IsOverlap(c Interval) bool
	IsMergeable(c Interval) bool
	Merge(c Interval) Interval
	Duration() float64
}

func NewInterval(startTime, endTime time.Time) (Interval, error) {
	if startTime.IsZero() {
		return &interval{}, errors.New("creating NewInterval: startTime is zero")
	}

	if endTime.IsZero() {
		return &interval{}, errors.New("creating NewInterval: endTime is zero")
	}

	if !endTime.After(startTime) {
		return &interval{}, errors.New("creating NewInterval: endTime must be after startTime")
	}

	return &interval{
		startTime: startTime,
		endTime:   endTime,
	}, nil
}

//
// func NewIntervalFromStr(startTimeStr, endTimeStr string) (Interval, error) {
// 	startTime = startTimeStr
// 	endTime = endTimeStr
	
// 	return NewInterval(startTime, endTime)
// }

/*
*
 */
func (i *interval) StartTime() time.Time {
	return i.startTime
}

/*
*
 */
func (i *interval) EndTime() time.Time {
	return i.endTime
}

/*
*
 */
func (i *interval) IsBefore(c Interval) bool {
	return i.endTime.Before(c.StartTime())
}

/*
*
 */
func (i *interval) IsAfter(c Interval) bool { // not used
	return i.startTime.After(c.EndTime())
}

/*
*
 */
func (i *interval) IsMergeable(c Interval) bool {
	return i.startTime.Compare(c.EndTime()) <= 0 && i.endTime.Compare(c.StartTime()) >= 0
}

/*
*
 */
func (i *interval) Merge(c Interval) Interval {
	if !i.IsMergeable(c) {
		return i
	}

	if i.startTime.After(c.StartTime()) {
		i.startTime = c.StartTime()
	}

	if i.endTime.Before(c.EndTime()) {
		i.endTime = c.EndTime()
	}

	return i
}

/*
*
 */
func (i *interval) Duration() float64 {

	return i.endTime.Sub(i.startTime).Minutes()
}

/*
*
 */
func (i *interval) IsOverlap(c Interval) bool { // used in not actual
	return i.startTime.Before(c.EndTime()) && i.endTime.After(c.StartTime())
}

// ************************************* Not used *************************************
/*
*
 */
func (i *interval) IsBeforeOrAdjacent(c Interval) bool {
	return i.endTime.Compare(c.StartTime()) <= 0
}

/*
*
 */
func (i *interval) IsAfterOrAdjacent(c Interval) bool {
	return i.startTime.Compare(c.EndTime()) >= 0
}

/*
*
 */
func (i *interval) Intersect(c Interval) Interval {

	if i.startTime.Before(c.StartTime()) {
		i.startTime = c.StartTime()
	}

	if i.endTime.After(c.EndTime()) {
		i.endTime = c.EndTime()
	}

	// remove c *interval ?

	return i
}
