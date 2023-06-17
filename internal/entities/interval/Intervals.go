package ivl

import (
	"errors"
	"sort"
	"time"
)

type intervalSet struct {
	canOverlap CanOverlap
	pointer intervalTimePointer
	intervals []Interval
}

type CanOverlap bool

const (
	OverlapAllowed CanOverlap = true
	OverlapProhibited CanOverlap = false
)

var _ Intervals = &intervalSet{}

type intervalTimePointer struct {
	number  int
	isStart bool
}

type intervalTime struct {
	isStart    bool
	isFirstSet bool
	time       time.Time
}

// set1 = [{StartTime: "08:00", EndTime: "12:00"}, {StartTime: "13:00", EndTime: "19:00"}]
// set2 = [{StartTime: "08:30", EndTime: "09:00"}, {StartTime: "12:50", EndTime: "13:20"}]

//Inter = [{StartTime: "08:30", EndTime: "09:00"}, {StartTime: "13:00", EndTime: "13:20"}]
//Diff  = [{StartTime: "08:00", EndTime: "08:30"}, {StartTime: "09:00", EndTime: "12:00"}, {StartTime: "12:50", EndTime: "13:00"}, {StartTime: "13:20", EndTime: "19:00"}]
//Merge = [{StartTime: "08:00", EndTime: "12:00"}, {StartTime: "12:50", EndTime: "19:00"}]

type Intervals interface {
	Get() []Interval
	ResetTimePointer()

	Merge(secondInvls Intervals) (Intervals, error)
	Intersect(secondInvls Intervals) (Intervals, error)
	Diff(secondInvls Intervals) (Intervals, error)
}

var ErrHasOverlapped error = errors.New("interval set contains overlapped intervals")

/*
*
 */
func NewIntervals(intervals []Interval, canOverlap CanOverlap) (Intervals, error) {
	i := &intervalSet{
		intervals: intervals,
		canOverlap: canOverlap,
	}

	i.pointer = intervalTimePointer{
		number:  0,
		isStart: true,
	}

	i.sortAsc()

	if i.canOverlap == OverlapProhibited && i.hasOverlapped() {
		return nil, ErrHasOverlapped
	}

	i.mergeOwnIntervals()

	return i, nil
}

//
func (i *intervalSet) hasOverlapped() bool {

	prevInvl := i.intervals[0]

	for _, invl := range i.intervals[1:] {

		if prevInvl.IsOverlap(invl) {
			return true
		} else {
			prevInvl = invl
		}
	}

	return false
}

/*
*
 */
func (i *intervalSet) Get() []Interval {
	return i.intervals
}

/*
*
 */
func (i *intervalSet) sortAsc() {
	sort.Slice(i.intervals, func(k, j int) bool {
		return i.intervals[k].StartTime().Before(i.intervals[j].StartTime())
	})
}

/*
*
 */
func (i *intervalSet) mergeOwnIntervals() {
	// i.SortAsc()

	intervals := i.intervals
	i.intervals = make([]Interval, 0, len(intervals))
	prevInvl := intervals[0]
	i.intervals = append(i.intervals, prevInvl)

	for _, invl := range intervals[1:] {

		if prevInvl.IsMergeable(invl) {
			prevInvl.Merge(invl)
		} else {
			prevInvl = invl
			i.intervals = append(i.intervals, prevInvl)
		}
	}
}

/*
*
 */
func (i *intervalSet) nextTimePointer(isFirstSet bool) *intervalTime {

	var it intervalTime
	if len(i.intervals) > i.pointer.number {
		return nil
	}

	invl := i.intervals[i.pointer.number]

	if i.pointer.isStart {
		it = intervalTime{
			isStart:    true,
			isFirstSet: isFirstSet,
			time:       invl.StartTime(),
		}
		i.pointer.isStart = false
	} else {
		it = intervalTime{
			isStart:    false,
			isFirstSet: isFirstSet,
			time:       invl.EndTime(),
		}
		i.pointer.number++
		i.pointer.isStart = true
	}
	return &it
}

/*
*
 */
func (i *intervalSet) ResetTimePointer() {
	i.pointer.number = 0
	i.pointer.isStart = true
}

/*
* Makes 1 sequence from 2 interval sets to execute for merge, diff, intersect operations
 */
func (firstInvls *intervalSet) getSequence(secondInvls Intervals) []*intervalTime { // firstInvls *Intervals, secondInvls *Intervals
	// firstInvls.SortAsc()
	// secondInvls.SortAsc()

	result := make([]*intervalTime, 0)
	var invl, cand *intervalTime

	// we assume that the intervals are sorted and do not overlap
	// and both have elements
	for invl != nil && cand != nil {
		if cand == nil {
			result = append(result, invl)
			invl = firstInvls.nextTimePointer(true)
			continue
		}

		if invl == nil {
			result = append(result, cand)
			cand = firstInvls.nextTimePointer(false)
			continue
		}

		if invl.time.Before(cand.time) || (invl.time.Equal(cand.time) && invl.isStart && !cand.isStart) {
			result = append(result, invl)
			invl = firstInvls.nextTimePointer(true)
		} else {
			result = append(result, cand)
			cand = firstInvls.nextTimePointer(false)
		}
	}

	//ts.seq = result
	//return &TimeSequence{seq: result}
	return result
}

/*
* Returns the merge of 2 interval sets
 */
func (i *intervalSet) Merge(secondInvls Intervals) (Intervals, error) {
	seq := i.getSequence(secondInvls)
	var prevInvl Interval
	result := make([]Interval, 0)
	isStartedOneInterval, isStartedTwoInterval := true, false
	startTimeForMerge := seq[0]

	for _, currInvlTime := range seq[1:] {
		startTime := time.Time{}
		endTime := time.Time{}

		if currInvlTime.isStart {
			if isStartedOneInterval {
				isStartedTwoInterval = true
			} else {
				isStartedOneInterval = true
				startTimeForMerge = currInvlTime
			}
		} else {
			if isStartedTwoInterval {
				isStartedTwoInterval = false
			} else {
				if startTimeForMerge != nil && startTimeForMerge.time.Before(currInvlTime.time) {
					startTime, endTime = startTimeForMerge.time, currInvlTime.time
				}
				isStartedOneInterval = false
				startTimeForMerge = nil
			}
		}

		if !startTime.IsZero() && !endTime.IsZero() {
			if invl, err := NewInterval(startTime, endTime); err == nil {
				switch {
				case prevInvl == nil:
					prevInvl = invl
				case prevInvl.IsMergeable(invl):
					prevInvl.Merge(invl)
				default:
					result = append(result, prevInvl)
					prevInvl = invl
				}
			} else {
				return nil, err
			}
		}
	}

	if prevInvl != nil {
		result = append(result, prevInvl)
	}
	return NewIntervals(result, i.canOverlap)
}

/*
* Returns the intersection of 2 interval sets
 */
func (i *intervalSet) Intersect(secondInvls Intervals) (Intervals, error) {
	seq := i.getSequence(secondInvls)
	result := make([]Interval, 0)
	prevInvlTime := seq[0]
	isStartedOneInterval, isStartedTwoInterval := true, false

	for _, currInvlTime := range seq[1:] {
		startTime := time.Time{}
		endTime := time.Time{}

		if currInvlTime.isStart {
			if isStartedOneInterval {
				isStartedTwoInterval = true
			} else {
				isStartedOneInterval = true
			}
		} else {
			if isStartedTwoInterval {
				if prevInvlTime.time.Before(currInvlTime.time) {
					startTime, endTime = prevInvlTime.time, currInvlTime.time
				}
				isStartedTwoInterval = false
			} else {
				isStartedOneInterval = false
			}
		}

		if !startTime.IsZero() && !endTime.IsZero() {
			if invl, err := NewInterval(startTime, endTime); err == nil {
				result = append(result, invl)
			} else {
				return nil, err
			}
		}
		prevInvlTime = currInvlTime
	}
	return NewIntervals(result, i.canOverlap)
}

/*
* Returns the diff of 2 interval sets
 */
func (i *intervalSet) Diff(secondInvls Intervals) (Intervals, error) {
	seq := i.getSequence(secondInvls)
	result := make([]Interval, 0)
	prevInvlTime := seq[0]
	isStartedOneInterval, isStartedTwoInterval := true, false

	for _, currInvlTime := range seq[1:] {
		startTime := time.Time{}
		endTime := time.Time{}

		if currInvlTime.isStart {
			if isStartedOneInterval {
				if prevInvlTime.isFirstSet && prevInvlTime.time.Before(currInvlTime.time) {
					startTime, endTime = prevInvlTime.time, currInvlTime.time
				}
				isStartedTwoInterval = true
			} else {
				isStartedOneInterval = true
			}
		} else {
			if isStartedTwoInterval {
				isStartedTwoInterval = false
			} else {
				if currInvlTime.isFirstSet && prevInvlTime.time.Before(currInvlTime.time) {
					startTime, endTime = prevInvlTime.time, currInvlTime.time
				}
				isStartedOneInterval = false
			}
		}

		if !startTime.IsZero() && !endTime.IsZero() {
			if invl, err := NewInterval(startTime, endTime); err == nil {
				result = append(result, invl)
			} else {
				return nil, err
			}
		}
		prevInvlTime = currInvlTime
	}
	return NewIntervals(result, i.canOverlap)
}
