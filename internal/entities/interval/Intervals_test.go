package ivl

import (
	"testing"
	"time"
)

func TestNewIntervals(t *testing.T) {

	hh08 := time.Date(2023, time.May, 1, 8, 0, 0, 0, time.UTC)
	hh10 := time.Date(2023, time.May, 1, 10, 0, 0, 0, time.UTC)
	hh12 := time.Date(2023, time.May, 1, 12, 0, 0, 0, time.UTC)
	hh12mm01 := time.Date(2023, time.May, 1, 12, 1, 0, 0, time.UTC)
	hh14 := time.Date(2023, time.May, 1, 14, 0, 0, 0, time.UTC)
	hh16 := time.Date(2023, time.May, 1, 16, 0, 0, 0, time.UTC)
	hh18 := time.Date(2023, time.May, 1, 18, 0, 0, 0, time.UTC)
	hh20 := time.Date(2023, time.May, 1, 20, 0, 0, 0, time.UTC)
	hh22 := time.Date(2023, time.May, 1, 22, 0, 0, 0, time.UTC)

	type tCase struct {
		caseName string
		overlap CanOverlap
		expected []time.Time
		data []time.Time
	}

	t.Run("positive", func(t *testing.T) {
		testCases := []tCase{
			{
				caseName: "[08 - 12][12:01 - 14]",
				overlap: OverlapProhibited,
				expected: []time.Time{hh08, hh12, hh12mm01, hh14},
				data: []time.Time{hh08, hh12, hh12mm01, hh14},
			},
			{
				caseName: "overlap [08 - 12:01][12 - 14]",
				overlap: OverlapAllowed,
				expected: []time.Time{hh08, hh14},
				data: []time.Time{hh08, hh12mm01, hh12, hh14},
			},
			{
				caseName: "unsorted [12:01 - 14][10 - 12]",
				overlap: OverlapProhibited,
				expected: []time.Time{hh10, hh12, hh12mm01, hh14},
				data: []time.Time{hh12mm01, hh14, hh10, hh12},
			},
			{
				caseName: "unsorted [12:01 - 14][08 - 12]",
				overlap: OverlapProhibited,
				expected: []time.Time{hh08, hh12, hh12mm01, hh14},
				data: []time.Time{hh12mm01, hh14, hh08, hh12},
			},
			{
				caseName: "[08 - 10][12 - 12:01]",
				overlap: OverlapProhibited,
				expected: []time.Time{hh08, hh10, hh12, hh12mm01},
				data: []time.Time{hh08, hh10, hh12, hh12mm01},
			},
			{
				caseName: "unsorted, adjacent [20-22][16-18][14-16][12-12:01][08-10]",
				overlap: OverlapProhibited,
				expected: []time.Time{hh08, hh10, hh12, hh12mm01, hh14, hh18, hh20, hh22},
				data: []time.Time{hh20, hh22, hh16, hh18, hh14, hh16, hh12, hh12mm01, hh08, hh10},
			},
			{
				caseName: "unsorted, adjacent, overlap [20-22][16-18][14-16][12-12:01][08-22]",
				overlap: OverlapAllowed,
				expected: []time.Time{hh08, hh22},
				data: []time.Time{hh20, hh22, hh16, hh18, hh14, hh16, hh12, hh12mm01, hh08, hh22},
			},
		}

		for _, testCase := range testCases {
			t.Run("Valid intervals "+testCase.caseName, func(t *testing.T) {
				invlSet := make([]Interval, 0)

				for i := 0; i < len(testCase.data); i += 2 {
					if invl, err := NewInterval(testCase.data[i], testCase.data[i+1]); err == nil {
						invlSet = append(invlSet, invl)
					} else {
						t.Errorf("Unexpected error for interval data in test case %s, error: %s", testCase.caseName, err)
						return
					}
				}

				invls, err := NewIntervals(invlSet, testCase.overlap)

				if err != nil {
					t.Errorf("Error not expected for %s, error: %s", testCase.caseName, err.Error())
					return
				}

				intervals := invls.Get()
				expectedCount := len(testCase.expected) / 2

				if expectedCount != len(intervals) {
					t.Errorf("not expected count intervals %d != %d", expectedCount, len(intervals))
				}

				for i := 0; i < expectedCount; i++ {
					if ! intervals[i].StartTime().Equal(testCase.expected[i*2]) {
						t.Errorf("incorrect interval StartTime = %s, expected= %s", 
							intervals[i].StartTime().Format(IntervalTimeFormat), testCase.expected[i*2].Format(IntervalTimeFormat))
					}
					if ! intervals[i].EndTime().Equal(testCase.expected[i*2+1]) {
						t.Errorf("incorrect interval EndTime = %s, expected= %s", 
							intervals[i].EndTime().Format(IntervalTimeFormat), testCase.expected[i*2+1].Format(IntervalTimeFormat))
					}
				}
			})
		}
	})

	t.Run("negative", func(t *testing.T) {
		testCases := []tCase{
			{
				caseName: "overlap [08 - 22][12:01 - 14]",
				overlap: OverlapProhibited,
				expected: []time.Time{hh08, hh22},
				data: []time.Time{hh08, hh22, hh12mm01, hh14},
			},
			{
				caseName: "overlap [08 - 12:01][12 - 14]",
				overlap: OverlapProhibited,
				expected: []time.Time{hh08, hh14},
				data: []time.Time{hh08, hh12mm01, hh12, hh14},
			},
			{
				caseName: "[08 - 22][12 - 12:01]",
				overlap: OverlapProhibited,
				expected: []time.Time{hh08, hh22},
				data: []time.Time{hh08, hh22, hh12, hh12mm01},
			},
			{
				caseName: "unsorted, adjacent, overlap [16-22][16-18][14-16][12-12:01][08-10]",
				overlap: OverlapProhibited,
				expected: []time.Time{hh08, hh10, hh12, hh12mm01, hh14, hh22},
				data: []time.Time{hh16, hh22, hh16, hh18, hh14, hh16, hh12, hh12mm01, hh08, hh10},
			},
			{
				caseName: "unsorted, adjacent, overlap [20-22][16-18][14-16][12-12:01][08-22]",
				overlap: OverlapProhibited,
				expected: []time.Time{hh08, hh22},
				data: []time.Time{hh20, hh22, hh16, hh18, hh14, hh16, hh12, hh12mm01, hh08, hh22},
			},
		}

		for _, testCase := range testCases {
			t.Run("Valid intervals "+testCase.caseName, func(t *testing.T) {
				invlSet := make([]Interval, 0)

				for i := 0; i < len(testCase.data); i += 2 {
					if invl, err := NewInterval(testCase.data[i], testCase.data[i+1]); err == nil {
						invlSet = append(invlSet, invl)
					} else {
						t.Errorf("Unexpected error for interval data in test case %s, error: %s", testCase.caseName, err)
						return
					}
				}

				_, err := NewIntervals(invlSet, testCase.overlap)

				if err == nil {
					t.Errorf("Expected error for %s", testCase.caseName)
					return
				}
			})
		}
	})
}

func TestMergeOwnIntervals(t *testing.T) {

	t.Run("positive", func(t *testing.T) {

	})

	t.Run("negative", func(t *testing.T) {

	})

}
