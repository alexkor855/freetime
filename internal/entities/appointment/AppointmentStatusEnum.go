package apmnt

type AppointmentStatus uint8

const (
	Reserved AppointmentStatus = iota + 1
	Booked
	Canceled
)