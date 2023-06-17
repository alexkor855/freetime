package idate

import (
	"testing"
	"time"
)

func getPositiveTestCases() map[string]date {
	return map[string]date{
		"2023-05-29": {2023, time.May, 29},
		"2023-01-01": {2023, time.January, 1},
		"2020-02-29": {2020, time.February, 29},
	}
}

func getNegativeTestCases() map[string]date {
	return map[string]date{
		"2023-05-32": {2023, time.May, 32},
		"2023-01-01": {2023, time.January, -1},
		"2023-02-29": {2023, time.February, 29},
		"2023-04-31": {2023, time.April, 31},
	}
}

func TestCreateDate(t *testing.T) {

	t.Run("positive", func(t *testing.T) {

		for strValue, testCase := range getPositiveTestCases() {
			t.Run("Correct date "+strValue, func(t *testing.T) {
				year := testCase.Year()
				month := testCase.month
				day := testCase.day

				date, err := CreateDate(year, month, day)
				if err != nil {
					t.Errorf("Error not expected for %s", strValue)
				}

				if err == nil && date.String() != strValue {
					t.Errorf("Date string not equal %s != %s", date.String(), strValue)
				}
			})
		}
	})

	t.Run("negative", func(t *testing.T) {

		for strValue, testCase := range getNegativeTestCases() {
			t.Run("Incorrect date "+strValue, func(t *testing.T) {
				year := testCase.Year()
				month := testCase.month
				day := testCase.day

				_, err := CreateDate(year, month, day)
				if err == nil {
					t.Error("Expected error")
				}
			})
		}
	})
}

func TestIsEqual(t *testing.T) {

	t.Run("positive", func(t *testing.T) {
		date1, _ := CreateDate(2023, time.January, 1)
		t.Run("Expected equal", func(t *testing.T) {
			date2, _ := CreateDate(2023, time.January, 1)

			if !date1.IsEqual(date2) {
				t.Errorf("Expected equal result: %d, %d, %d != %d, %d, %d", date1.Year(), date1.Month(), date1.Day(), date2.Year(), date2.Month(), date2.Day())
			}
		})
	})

	t.Run("negative", func(t *testing.T) {
		date1, _ := CreateDate(2023, time.January, 2)
		t.Run("Expected not equal", func(t *testing.T) {
			date2, _ := CreateDate(2023, time.January, 1)

			if date1.IsEqual(date2) {
				t.Errorf("Expected not equal result: %d, %d, %d == %d, %d, %d", date1.Year(), date1.Month(), date1.Day(), date2.Year(), date2.Month(), date2.Day())
			}
		})
	})
}

func TestIsValidDate(t *testing.T) {

	t.Run("positive", func(t *testing.T) {

		for strValue, testCase := range getPositiveTestCases() {
			t.Run("Valid date "+strValue, func(t *testing.T) {
				year := testCase.Year()
				month := testCase.month
				day := testCase.day

				if !IsValidDate(year, month, day) {
					t.Errorf("Expected valid result: %d, %d, %d", year, month, day)
				}
			})
		}
	})

	t.Run("negative", func(t *testing.T) {

		for strValue, testCase := range getNegativeTestCases() {
			t.Run("Invalid date "+strValue, func(t *testing.T) {
				year := testCase.Year()
				month := testCase.month
				day := testCase.day

				if IsValidDate(year, month, day) {
					t.Errorf("Expected invalid result: %d, %d, %d", year, month, day)
				}
			})
		}
	})
}
