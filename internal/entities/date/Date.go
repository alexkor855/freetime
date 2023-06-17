package idate

import (
	"errors"
	"time"
)

var _ Date = &date{}

type date struct {
	year  int
	month time.Month
	day   int
}

type Date interface {
	Year() int
	Month() time.Month
	Day() int
	IsEqual(c Date) bool
	IsEarlierOrEqual(c Date) bool
	String() string
}

func (d *date) Year() int {
	return d.year
}

func (d *date) Month() time.Month {
	return d.month
}

func (d *date) Day() int {
	return d.day
}

func CreateDate(year int, month time.Month, day int) (Date, error) {

	if !IsValidDate(year, month, day) {
		return nil, errors.New("wrong date")
	}

	d := &date{
		year:  year,
		month: month,
		day:   day,
	}

	return d, nil
}

func CreateDateFromYMD(dateStr string) (Date, error) {
	dateTime, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, err
	}

	return CreateDateFromTime(dateTime), nil
}

func CreateDateFromTime(t time.Time) Date {
	year, month, day := t.Date()

	d := &date{
		year:  year,
		month: month,
		day:   day,
	}

	return d
}

func IsValidDate(year int, month time.Month, day int) bool {

	t := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	y, m, d := t.Date()

	return y == year && m == month && d == day
}

func (d *date) IsEqual(c Date) bool {
	return d.Day() == c.Day() && d.Month() == c.Month() && d.Year() == c.Year()
}

func (d *date) IsEarlierOrEqual(c Date) bool {
	day1 := time.Date(d.year, d.month, d.day, 0, 0, 0, 0, time.UTC)
	day2 := time.Date(c.Year(), c.Month(), c.Day(), 0, 0, 0, 0, time.UTC)
	return day1.Compare(day2) <= 0
}

func (d *date) String() string {
	return time.Date(d.year, d.month, d.day, 0, 0, 0, 0, time.UTC).Format(time.DateOnly)
	//return strconv.Itoa(d.year) + "-" + d.month.String() + "-" + strconv.Itoa(d.day)
}
