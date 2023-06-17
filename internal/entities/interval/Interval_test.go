package ivl

import (
	"testing"
	"time"
)

func TestNewInterval(t *testing.T) {

	t.Run("positive", func(t *testing.T) {
		hh00 := time.Date(2023, time.May, 1, 0, 0, 0, 0, time.UTC)
		hh10 := time.Date(2023, time.May, 1, 10, 0, 0, 0, time.UTC)
		hh12 := time.Date(2023, time.May, 1, 12, 0, 0, 0, time.UTC)

		testCases := map[string][2]time.Time{
			"[00 - 12]": {hh00, hh12},
			"[10 - 12]": {hh10, hh12},
		}

		for strValue, testCase := range testCases {
			t.Run("Valid interval "+strValue, func(t *testing.T) {
				invl, err := NewInterval(testCase[0], testCase[1])

				if err != nil {
					t.Errorf("Error not expected for %s", strValue)
					return
				}

				switch {
				case invl.StartTime() != testCase[0]:
					t.Errorf("Error not expected StartTime() != %s", testCase[0].Format(time.TimeOnly))

				case invl.EndTime() != testCase[1]:
					t.Errorf("Error not expected EndTime() != %s", testCase[1].Format(time.TimeOnly))
				}
			})
		}
	})

	t.Run("negative", func(t *testing.T) {
		hh10 := time.Date(2023, time.May, 1, 10, 0, 0, 0, time.UTC)
		hh12 := time.Date(2023, time.May, 1, 12, 0, 0, 0, time.UTC)
		empty := time.Time{}

		testCases := map[string][2]time.Time{
			"reverse time":             {hh12, hh10},
			"same time":                {hh12, hh12},
			"empty start time":         {empty, hh12},
			"empty end time":           {hh12, empty},
			"empty start and end time": {empty, empty},
		}

		for strValue, testCase := range testCases {
			t.Run("Invalid interval "+strValue, func(t *testing.T) {
				_, err := NewInterval(testCase[0], testCase[1])

				if err == nil {
					t.Error("Expected error")
				}
			})
		}
	})
}

func TestIsBefore(t *testing.T) {
	hh08 := time.Date(2023, time.May, 1, 8, 0, 0, 0, time.UTC)
	hh10 := time.Date(2023, time.May, 1, 10, 0, 0, 0, time.UTC)
	hh12 := time.Date(2023, time.May, 1, 12, 0, 0, 0, time.UTC)
	hh12mm01 := time.Date(2023, time.May, 1, 12, 1, 0, 0, time.UTC)
	hh14 := time.Date(2023, time.May, 1, 14, 0, 0, 0, time.UTC)

	t.Run("positive", func(t *testing.T) {

		testCases := map[string][4]time.Time{
			"[08:00 - 10:00][12:00 - 14:00]": {hh08, hh10, hh12, hh14},
			"[10:00 - 12:00][12:01 - 14:00]": {hh10, hh12, hh12mm01, hh14},
		}

		for caseName, testCase := range testCases {
			t.Run("IsBefore "+caseName, func(t *testing.T) {

				invlFirst, err1 := NewInterval(testCase[0], testCase[1])
				invlSecond, err2 := NewInterval(testCase[2], testCase[3])

				if err1 != nil || err2 != nil {
					t.Errorf("Error of creating is not expected for %s", caseName)
					return
				}

				if !invlFirst.IsBefore(invlSecond) {
					t.Errorf("Expect interval [%s - %s] before [%s, %s]",
						invlFirst.StartTime().Format(time.TimeOnly), invlFirst.EndTime().Format(time.TimeOnly),
						invlSecond.StartTime().Format(time.TimeOnly), invlSecond.EndTime().Format(time.TimeOnly))
				}

			})
		}
	})

	t.Run("negative", func(t *testing.T) {

		testCases := map[string][4]time.Time{
			"has overlaped end [08:00 - 12:00][10:00 - 14:00]":       {hh08, hh12, hh10, hh14},
			"after [12:01 - 14:00][10:00 - 12:00]":                   {hh12mm01, hh14, hh10, hh12},
			"overlap [08:00 - 14:00][10:00 - 12:00]":                 {hh08, hh14, hh10, hh12},
			"inside [10:00 - 12:00][08:00 - 14:00]":                  {hh10, hh12, hh08, hh14},
			"overlap end and touch[10:00 - 12:00][08:00 - 12:00]":    {hh10, hh12, hh08, hh12},
			"overlap end [10:00 - 14:00][08:00 - 12:00]":             {hh10, hh14, hh08, hh12},
			"overlap start and touch [08:00 - 12:00][08:00 - 14:00]": {hh08, hh12, hh08, hh14},
			"touch [10:00 - 12:00][12:00 - 14:00]":                   {hh10, hh12, hh12, hh14},
		}

		for caseName, testCase := range testCases {
			t.Run("IsBefore "+caseName, func(t *testing.T) {

				invlFirst, err1 := NewInterval(testCase[0], testCase[1])
				invlSecond, err2 := NewInterval(testCase[2], testCase[3])

				if err1 != nil || err2 != nil {
					t.Errorf("Error of creating is not expected for %s", caseName)
					return
				}

				if invlFirst.IsBefore(invlSecond) {
					t.Errorf("Expected error for %s", caseName)
				}
			})
		}
	})
}

func TestIsAfter(t *testing.T) {
	hh08 := time.Date(2023, time.May, 1, 8, 0, 0, 0, time.UTC)
	hh10 := time.Date(2023, time.May, 1, 10, 0, 0, 0, time.UTC)
	hh12 := time.Date(2023, time.May, 1, 12, 0, 0, 0, time.UTC)
	hh12mm01 := time.Date(2023, time.May, 1, 12, 1, 0, 0, time.UTC)
	hh14 := time.Date(2023, time.May, 1, 14, 0, 0, 0, time.UTC)

	t.Run("positive", func(t *testing.T) {

		testCases := map[string][4]time.Time{
			"[12:00 - 14:00][08:00 - 10:00]": {hh12, hh14, hh08, hh10},
			"[12:01 - 14:00][10:00 - 12:00]": {hh12mm01, hh14, hh10, hh12},
		}

		for caseName, testCase := range testCases {
			t.Run("IsAfter "+caseName, func(t *testing.T) {

				invlFirst, err1 := NewInterval(testCase[0], testCase[1])
				invlSecond, err2 := NewInterval(testCase[2], testCase[3])

				if err1 != nil || err2 != nil {
					t.Errorf("Error of creating is not expected for %s", caseName)
					return
				}

				if !invlFirst.IsAfter(invlSecond) {
					t.Errorf("Error not expected for %s", invlFirst)
				}

			})
		}
	})

	t.Run("negative", func(t *testing.T) {

		testCases := map[string][4]time.Time{
			"has overlaped end [08:00 - 12:00][10:00 - 14:00]":       {hh08, hh12, hh10, hh14},
			"before [10:00 - 12:00][12:01 - 14:00]":                  {hh10, hh12, hh12mm01, hh14},
			"overlap [08:00 - 14:00][10:00 - 12:00]":                 {hh08, hh14, hh10, hh12},
			"inside [10:00 - 12:00][08:00 - 14:00]":                  {hh10, hh12, hh08, hh14},
			"overlap end and touch[10:00 - 12:00][08:00 - 12:00]":    {hh10, hh12, hh08, hh12},
			"overlap end [10:00 - 14:00][08:00 - 12:00]":             {hh10, hh14, hh08, hh12},
			"overlap start and touch [08:00 - 12:00][08:00 - 14:00]": {hh08, hh12, hh08, hh14},
			"touch [10:00 - 12:00][12:00 - 14:00]":                   {hh10, hh12, hh12, hh14},
		}

		for caseName, testCase := range testCases {
			t.Run("IsAfter "+caseName, func(t *testing.T) {

				invlFirst, err1 := NewInterval(testCase[0], testCase[1])
				invlSecond, err2 := NewInterval(testCase[2], testCase[3])

				if err1 != nil || err2 != nil {
					t.Errorf("Error of creating is not expected for %s", caseName)
					return
				}

				if invlFirst.IsAfter(invlSecond) {
					t.Errorf("Expected error for %s", invlFirst)
				}
			})
		}
	})
}

func TestIsOverlap(t *testing.T) {
	hh08 := time.Date(2023, time.May, 1, 8, 0, 0, 0, time.UTC)
	hh10 := time.Date(2023, time.May, 1, 10, 0, 0, 0, time.UTC)
	hh12 := time.Date(2023, time.May, 1, 12, 0, 0, 0, time.UTC)
	hh12mm01 := time.Date(2023, time.May, 1, 12, 1, 0, 0, time.UTC)
	hh14 := time.Date(2023, time.May, 1, 14, 0, 0, 0, time.UTC)

	t.Run("positive", func(t *testing.T) {

		testCases := map[string][4]time.Time{
			"has overlaped end [08:00 - 12:00][10:00 - 14:00]":       {hh08, hh12, hh10, hh14},
			"overlap [08:00 - 14:00][10:00 - 12:00]":                 {hh08, hh14, hh10, hh12},
			"inside [10:00 - 12:00][08:00 - 14:00]":                  {hh10, hh12, hh08, hh14},
			"overlap end and touch[10:00 - 12:00][08:00 - 12:00]":    {hh10, hh12, hh08, hh12},
			"overlap end [10:00 - 14:00][08:00 - 12:00]":             {hh10, hh14, hh08, hh12},
			"overlap start and touch [08:00 - 12:00][08:00 - 14:00]": {hh08, hh12, hh08, hh14},
		}

		for caseName, testCase := range testCases {
			t.Run("IsOverlap "+caseName, func(t *testing.T) {

				invlFirst, err1 := NewInterval(testCase[0], testCase[1])
				invlSecond, err2 := NewInterval(testCase[2], testCase[3])

				if err1 != nil || err2 != nil {
					t.Errorf("Error of creating is not expected for %s", caseName)
					return
				}

				if !invlFirst.IsOverlap(invlSecond) {
					t.Errorf("Expected it is overlap for %s", caseName)
				}

			})
		}
	})

	t.Run("negative", func(t *testing.T) {

		testCases := map[string][4]time.Time{
			"[12:00 - 14:00][08:00 - 10:00]":       {hh12, hh14, hh08, hh10},
			"[12:01 - 14:00][10:00 - 12:00]":       {hh12mm01, hh14, hh10, hh12},
			"[08:00 - 10:00][12:00 - 14:00]":       {hh08, hh10, hh12, hh14},
			"[10:00 - 12:00][12:01 - 14:00]":       {hh10, hh12, hh12mm01, hh14},
			"touch [10:00 - 12:00][12:00 - 14:00]": {hh10, hh12, hh12, hh14},
		}

		for caseName, testCase := range testCases {
			t.Run("IsOverlap "+caseName, func(t *testing.T) {

				invlFirst, err1 := NewInterval(testCase[0], testCase[1])
				invlSecond, err2 := NewInterval(testCase[2], testCase[3])

				if err1 != nil || err2 != nil {
					t.Errorf("Error of creating is not expected for %s", caseName)
					return
				}

				if invlFirst.IsOverlap(invlSecond) {
					t.Errorf("Expected it is not overlap for %s", caseName)
				}
			})
		}
	})
}

func TestIsMergeable(t *testing.T) {
	hh08 := time.Date(2023, time.May, 1, 8, 0, 0, 0, time.UTC)
	hh10 := time.Date(2023, time.May, 1, 10, 0, 0, 0, time.UTC)
	hh12 := time.Date(2023, time.May, 1, 12, 0, 0, 0, time.UTC)
	hh12mm01 := time.Date(2023, time.May, 1, 12, 1, 0, 0, time.UTC)
	hh14 := time.Date(2023, time.May, 1, 14, 0, 0, 0, time.UTC)

	t.Run("positive", func(t *testing.T) {

		testCases := map[string][4]time.Time{
			"has overlaped end [08:00 - 12:00][10:00 - 14:00]":       {hh08, hh12, hh10, hh14},
			"overlap [08:00 - 14:00][10:00 - 12:00]":                 {hh08, hh14, hh10, hh12},
			"inside [10:00 - 12:00][08:00 - 14:00]":                  {hh10, hh12, hh08, hh14},
			"overlap end and touch[10:00 - 12:00][08:00 - 12:00]":    {hh10, hh12, hh08, hh12},
			"overlap end [10:00 - 14:00][08:00 - 12:00]":             {hh10, hh14, hh08, hh12},
			"overlap start and touch [08:00 - 12:00][08:00 - 14:00]": {hh08, hh12, hh08, hh14},
			"touch [10:00 - 12:00][12:00 - 14:00]":                   {hh10, hh12, hh12, hh14},
		}

		for caseName, testCase := range testCases {
			t.Run("IsMergeable "+caseName, func(t *testing.T) {

				invlFirst, err1 := NewInterval(testCase[0], testCase[1])
				invlSecond, err2 := NewInterval(testCase[2], testCase[3])

				if err1 != nil || err2 != nil {
					t.Errorf("Error of creating is not expected for %s", caseName)
					return
				}

				if !invlFirst.IsMergeable(invlSecond) {
					t.Errorf("Expected it is mergeable for %s, (%s, %s)", invlFirst, err1, err2)
				}

			})
		}
	})

	t.Run("negative", func(t *testing.T) {

		testCases := map[string][4]time.Time{
			"[12:00 - 14:00][08:00 - 10:00]": {hh12, hh14, hh08, hh10},
			"[12:01 - 14:00][10:00 - 12:00]": {hh12mm01, hh14, hh10, hh12},
			"[08:00 - 10:00][12:00 - 14:00]": {hh08, hh10, hh12, hh14},
			"[10:00 - 12:00][12:01 - 14:00]": {hh10, hh12, hh12mm01, hh14},
		}

		for caseName, testCase := range testCases {
			t.Run("IsMergeable "+caseName, func(t *testing.T) {

				invlFirst, err1 := NewInterval(testCase[0], testCase[1])
				invlSecond, err2 := NewInterval(testCase[2], testCase[3])

				if err1 != nil || err2 != nil {
					t.Errorf("Error of creating is not expected for %s", caseName)
					return
				}

				if invlFirst.IsMergeable(invlSecond) {
					t.Errorf("Expected it is not mergeable for %s", caseName)
				}
			})
		}
	})
}

func TestMerge(t *testing.T) {
	hh08 := time.Date(2023, time.May, 1, 8, 0, 0, 0, time.UTC)
	hh10 := time.Date(2023, time.May, 1, 10, 0, 0, 0, time.UTC)
	hh12 := time.Date(2023, time.May, 1, 12, 0, 0, 0, time.UTC)
	hh12mm01 := time.Date(2023, time.May, 1, 12, 1, 0, 0, time.UTC)
	hh14 := time.Date(2023, time.May, 1, 14, 0, 0, 0, time.UTC)

	t.Run("positive", func(t *testing.T) {

		testCases := map[string][6]time.Time{
			"has overlaped end [08:00 - 12:00][10:00 - 14:00]":       {hh08, hh12, hh10, hh14, hh08, hh14},
			"overlap [08:00 - 14:00][10:00 - 12:00]":                 {hh08, hh14, hh10, hh12, hh08, hh14},
			"inside [10:00 - 12:00][08:00 - 14:00]":                  {hh10, hh12, hh08, hh14, hh08, hh14},
			"overlap end and touch[10:00 - 12:00][08:00 - 12:00]":    {hh10, hh12, hh08, hh12, hh08, hh12},
			"overlap end [10:00 - 14:00][08:00 - 12:00]":             {hh10, hh14, hh08, hh12, hh08, hh14},
			"overlap start and touch [08:00 - 12:00][08:00 - 14:00]": {hh08, hh12, hh08, hh14, hh08, hh14},
			"touch [10:00 - 12:00][12:00 - 14:00]":                   {hh10, hh12, hh12, hh14, hh10, hh14},
		}

		for caseName, testCase := range testCases {
			t.Run("Merge "+caseName, func(t *testing.T) {

				invlFirst, err1 := NewInterval(testCase[0], testCase[1])
				invlSecond, err2 := NewInterval(testCase[2], testCase[3])

				if err1 != nil || err2 != nil {
					t.Errorf("Error of creating is not expected for %s", caseName)
					return
				}

				invlFirst.Merge(invlSecond)

				if !invlFirst.StartTime().Equal(testCase[4]) {
					t.Errorf("Expected it is equal %s and %s)", invlFirst.StartTime(), testCase[4])
				}

				if !invlFirst.EndTime().Equal(testCase[5]) {
					t.Errorf("Expected it is equal %s and %s)", invlFirst.EndTime(), testCase[5])
				}
			})
		}
	})

	t.Run("negative", func(t *testing.T) {

		testCases := map[string][4]time.Time{
			"[12:00 - 14:00][08:00 - 10:00]": {hh12, hh14, hh08, hh10},
			"[12:01 - 14:00][10:00 - 12:00]": {hh12mm01, hh14, hh10, hh12},
			"[08:00 - 10:00][12:00 - 14:00]": {hh08, hh10, hh12, hh14},
			"[10:00 - 12:00][12:01 - 14:00]": {hh10, hh12, hh12mm01, hh14},
		}

		for caseName, testCase := range testCases {
			t.Run("Merge "+caseName, func(t *testing.T) {

				invlFirst, err1 := NewInterval(testCase[0], testCase[1])
				invlSecond, err2 := NewInterval(testCase[2], testCase[3])

				if err1 != nil || err2 != nil {
					t.Errorf("Error of creating is not expected for %s", caseName)
					return
				}

				invlFirst.Merge(invlSecond)

				if !invlFirst.StartTime().Equal(testCase[0]) || !invlFirst.EndTime().Equal(testCase[1]) {
					t.Error("Expected that the interval should not change")
				}
			})
		}
	})
}

func TestDuration(t *testing.T) {

	hh00 := time.Date(2023, time.May, 1, 0, 0, 0, 0, time.UTC)
	hh10 := time.Date(2023, time.May, 1, 10, 0, 0, 0, time.UTC)
	hh12 := time.Date(2023, time.May, 1, 12, 0, 0, 0, time.UTC)
	hh12mm01 := time.Date(2023, time.May, 1, 12, 1, 0, 0, time.UTC)

	testCases := map[string][2]time.Time{
		"[00:00 - 12:00]": {hh00, hh12},     // 12*60=720
		"[10:00 - 12:00]": {hh10, hh12},     // 2*60=120
		"[12:00 - 12:01]": {hh12, hh12mm01}, // 1
	}

	testCasesRes := map[string]int{
		"[00:00 - 12:00]": 720,
		"[10:00 - 12:00]": 120,
		"[12:00 - 12:01]": 1,
	}

	for caseName, testCase := range testCases {
		t.Run("Duration "+caseName, func(t *testing.T) {
			invl, err := NewInterval(testCase[0], testCase[1])

			if err != nil {
				t.Errorf("Error not expected for %s", caseName)
				return
			}

			if int(invl.Duration()) != testCasesRes[caseName] {
				t.Errorf("Duration must be equal == %d", testCasesRes[caseName])
			}
		})
	}
}
