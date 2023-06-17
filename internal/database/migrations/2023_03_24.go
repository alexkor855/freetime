package migrations

import (
	"time"

	//"github.com/jinzhu/gorm"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

type (
	Schedule struct {
		ID           string
		ScheduleType uint8
		BranchId     string
		EmployeeId   string
		WorkplaceId  string
		ServiceId    string
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}
	ScheduleInterval struct {
		ID         string
		ScheduleId string
		Date       time.Time
		IntervalId string
	}
	Interval struct {
		ID        string
		StartTime time.Time
		EndTime   time.Time
	}
	Appointment struct {
		ID         string
		BranchId   string
		Status     uint8
		CustomerId string
		Date       time.Time
		StartTime  time.Time
		EndTime    time.Time
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}
)

func main() {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai", // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})

	if err != nil {
		panic("failed to open database connection")
	}

	// sqlDB, err := db.DB()
	// if err != nil {
	// 	panic("failed to get connection to database")
	// }

	// defer sqlDB.Close()

	db.AutoMigrate(&Schedule{}, &ScheduleInterval{}, &Interval{}, &Appointment{})
}
