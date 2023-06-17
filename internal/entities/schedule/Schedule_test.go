package sch

import (
	idate "app/internal/entities/date"
	"testing"
	"time"
)

func TestNewSchedule(t *testing.T) {

	t.Run("positive", func(t *testing.T) {

		for strValue, testCase := range getPositiveTestCases() {

			t.Run("check "+strValue, func(t *testing.T) {
				_, err := NewSchedule(
					testCase.id,
					testCase.scheduleTemplateId,
					testCase.scheduleType,
					testCase.branchId,
					testCase.employeeId,
					testCase.workplaceId,
					testCase.serviceId,
					testCase.startDate,
					testCase.endDate,
				)

				if err != nil {
					t.Errorf("Error not expected for %s", strValue)
				}
			})
		}
	})

	t.Run("negative", func(t *testing.T) {

		for strValue, testCase := range getNegativeTestCases() {

			t.Run("check "+strValue, func(t *testing.T) {
				_, err := NewSchedule(
					testCase.id,
					testCase.scheduleTemplateId,
					testCase.scheduleType,
					testCase.branchId,
					testCase.employeeId,
					testCase.workplaceId,
					testCase.serviceId,
					testCase.startDate,
					testCase.endDate,
				)

				if err == nil {
					t.Errorf("Errors expected for %s", strValue)
				}
			})
		}
	})
}

func TestGetters(t *testing.T) {
	startDate, _ := idate.CreateDate(2023, time.May, 15)
	endDate, _ := idate.CreateDate(2023, time.June, 15)
	testCase := schedule{id: "1", scheduleTemplateId: "1", scheduleType: ServiceEmployeeWorkplaceSchedule,
		branchId: "10", employeeId: "10", workplaceId: "1", serviceId: "1", startDate: startDate, endDate: endDate}

	sch, err := NewSchedule(
		testCase.id,
		testCase.scheduleTemplateId,
		testCase.scheduleType,
		testCase.branchId,
		testCase.employeeId,
		testCase.workplaceId,
		testCase.serviceId,
		testCase.startDate,
		testCase.endDate,
	)

	if err != nil {
		t.Error("Error not expected")
	}

	switch {
	case sch.Id() != testCase.id:
		t.Error("sch.Id() != testCase.id")

	case sch.ScheduleTemplateId() != testCase.scheduleTemplateId:
		t.Error("sch.ScheduleTemplateId() != testCase.scheduleTemplateId")

	case sch.ScheduleType() != testCase.scheduleType:
		t.Error("sch.ScheduleType() != testCase.scheduleType")

	case sch.BranchId() != testCase.branchId:
		t.Error("sch.BranchId() != testCase.branchId")

	case sch.EmployeeId() != testCase.employeeId:
		t.Error("sch.EmployeeId() != testCase.employeeId")

	case sch.WorkplaceId() != testCase.workplaceId:
		t.Error("sch.WorkplaceId() != testCase.workplaceId")

	case sch.ServiceId() != testCase.serviceId:
		t.Error("sch.ServiceId() != testCase.serviceId")

	case sch.StartDate() != testCase.startDate:
		t.Error("sch.StartDate() != testCase.startDate")

	case sch.EndDate() != testCase.endDate:
		t.Error("sch.EndDate() != testCase.endDate")
	}
}

func getPositiveTestCases() map[string]schedule {
	startDate, _ := idate.CreateDate(2023, time.May, 15)
	endDate, _ := idate.CreateDate(2023, time.June, 15)

	return map[string]schedule{
		"BranchSchedule": {id: "1", scheduleTemplateId: "1", scheduleType: BranchSchedule,
			branchId: "10", employeeId: "", workplaceId: "", serviceId: "", startDate: startDate, endDate: endDate},

		"EmployeeSchedule": {id: "1", scheduleTemplateId: "1", scheduleType: EmployeeSchedule,
			branchId: "10", employeeId: "10", workplaceId: "", serviceId: "", startDate: startDate, endDate: endDate},

		"WorkplaceSchedule": {id: "1", scheduleTemplateId: "1", scheduleType: WorkplaceSchedule,
			branchId: "10", employeeId: "", workplaceId: "1", serviceId: "", startDate: startDate, endDate: endDate},

		"EmployeeWorkplaceSchedule": {id: "1", scheduleTemplateId: "1", scheduleType: EmployeeWorkplaceSchedule,
			branchId: "10", employeeId: "10", workplaceId: "1", serviceId: "", startDate: startDate, endDate: endDate},

		"ServiceSchedule": {id: "1", scheduleTemplateId: "1", scheduleType: ServiceSchedule,
			branchId: "10", employeeId: "", workplaceId: "", serviceId: "1", startDate: startDate, endDate: endDate},

		"ServiceEmployeeSchedule": {id: "1", scheduleTemplateId: "1", scheduleType: ServiceEmployeeSchedule,
			branchId: "10", employeeId: "10", workplaceId: "", serviceId: "1", startDate: startDate, endDate: endDate},

		"ServiceWorkplaceSchedule": {id: "1", scheduleTemplateId: "1", scheduleType: ServiceWorkplaceSchedule,
			branchId: "10", employeeId: "", workplaceId: "1", serviceId: "1", startDate: startDate, endDate: endDate},

		"ServiceEmployeeWorkplaceSchedule": {id: "1", scheduleTemplateId: "1", scheduleType: ServiceEmployeeWorkplaceSchedule,
			branchId: "10", employeeId: "10", workplaceId: "1", serviceId: "1", startDate: startDate, endDate: endDate},

		// date nil
		"BranchSchedule and Date nil": {id: "1", scheduleTemplateId: "1", scheduleType: BranchSchedule,
			branchId: "10", employeeId: "", workplaceId: "", serviceId: "", startDate: startDate, endDate: nil},
	}
}

func getNegativeTestCases() map[string]schedule {
	startDate, _ := idate.CreateDate(2023, time.May, 15)
	endDate, _ := idate.CreateDate(2023, time.June, 15)

	return map[string]schedule{
		// empty id
		"BranchSchedule empty id": {id: "1", scheduleTemplateId: "1", scheduleType: BranchSchedule,
			branchId: "", employeeId: "", workplaceId: "", serviceId: "", startDate: startDate, endDate: endDate},

		"EmployeeSchedule empty id": {id: "1", scheduleTemplateId: "1", scheduleType: EmployeeSchedule,
			branchId: "10", employeeId: "", workplaceId: "", serviceId: "", startDate: startDate, endDate: endDate},

		"WorkplaceSchedule empty id": {id: "1", scheduleTemplateId: "1", scheduleType: WorkplaceSchedule,
			branchId: "10", employeeId: "", workplaceId: "", serviceId: "1", startDate: startDate, endDate: endDate},

		"EmployeeWorkplaceSchedule empty id": {id: "1", scheduleTemplateId: "1", scheduleType: EmployeeWorkplaceSchedule,
			branchId: "10", employeeId: "", workplaceId: "", serviceId: "", startDate: startDate, endDate: endDate},

		"ServiceSchedule empty id": {id: "1", scheduleTemplateId: "1", scheduleType: ServiceSchedule,
			branchId: "", employeeId: "", workplaceId: "", serviceId: "", startDate: startDate, endDate: endDate},

		"ServiceEmployeeSchedule empty id": {id: "1", scheduleTemplateId: "1", scheduleType: ServiceEmployeeSchedule,
			branchId: "10", employeeId: "", workplaceId: "", serviceId: "", startDate: startDate, endDate: endDate},

		"ServiceWorkplaceSchedule empty id": {id: "1", scheduleTemplateId: "1", scheduleType: ServiceWorkplaceSchedule,
			branchId: "10", employeeId: "", workplaceId: "", serviceId: "", startDate: startDate, endDate: endDate},

		"ServiceEmployeeWorkplaceSchedule empty id": {id: "1", scheduleTemplateId: "1", scheduleType: ServiceEmployeeWorkplaceSchedule,
			branchId: "10", employeeId: "", workplaceId: "", serviceId: "", startDate: startDate, endDate: endDate},

		// date
		"BranchSchedule and Date replace": {id: "1", scheduleTemplateId: "1", scheduleType: BranchSchedule,
			branchId: "10", employeeId: "", workplaceId: "", serviceId: "", startDate: endDate, endDate: startDate},

		"BranchSchedule and Date nil": {id: "1", scheduleTemplateId: "1", scheduleType: BranchSchedule,
			branchId: "10", employeeId: "", workplaceId: "", serviceId: "", startDate: nil, endDate: nil},

		// extra id
		"EmployeeSchedule extra id": {id: "1", scheduleTemplateId: "1", scheduleType: EmployeeSchedule,
			branchId: "10", employeeId: "10", workplaceId: "1", serviceId: "", startDate: startDate, endDate: endDate},

		"WorkplaceSchedule extra id": {id: "1", scheduleTemplateId: "1", scheduleType: WorkplaceSchedule,
			branchId: "10", employeeId: "1", workplaceId: "1", serviceId: "1", startDate: startDate, endDate: endDate},

		"EmployeeWorkplaceSchedule extra id": {id: "1", scheduleTemplateId: "1", scheduleType: EmployeeWorkplaceSchedule,
			branchId: "10", employeeId: "10", workplaceId: "1", serviceId: "1", startDate: startDate, endDate: endDate},

		"ServiceSchedule extra id": {id: "1", scheduleTemplateId: "1", scheduleType: ServiceSchedule,
			branchId: "10", employeeId: "10", workplaceId: "1", serviceId: "1", startDate: startDate, endDate: endDate},

		"ServiceEmployeeSchedule extra id": {id: "1", scheduleTemplateId: "1", scheduleType: ServiceEmployeeSchedule,
			branchId: "10", employeeId: "10", workplaceId: "1", serviceId: "1", startDate: startDate, endDate: endDate},

		"ServiceWorkplaceSchedule extra id": {id: "1", scheduleTemplateId: "1", scheduleType: ServiceWorkplaceSchedule,
			branchId: "10", employeeId: "10", workplaceId: "1", serviceId: "1", startDate: startDate, endDate: endDate},

		"ServiceEmployeeWorkplaceSchedule": {id: "1", scheduleTemplateId: "1", scheduleType: ServiceEmployeeWorkplaceSchedule,
			branchId: "10", employeeId: "10", workplaceId: "", serviceId: "1", startDate: startDate, endDate: endDate},
	}
}
