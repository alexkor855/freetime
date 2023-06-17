package ftc

import (
	idate "app/internal/entities/date"
	"app/internal/usecases"
	getftime "app/internal/usecases/freetime"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service getftime.GetFreeTimeService
}

func (c *Controller) GetFreeTime(ctx echo.Context) (err error) {
	var freeTimeParam GetFreeTimeRequest
	//freeTimeParam := new(GetFreeTime)

	if err := ctx.Bind(freeTimeParam); err != nil {
		return ctx.String(http.StatusBadRequest, "bad request 1")
	}

	day, err := idate.CreateDateFromYMD(freeTimeParam.Date)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request 2")
	}

	request := usecases.GetRequestDTO{
		Date:        day,
		BranchId:    freeTimeParam.BranchId,
		EmployeeId:  freeTimeParam.EmployeeId,
		WorkplaceId: freeTimeParam.WorkplaceId,
		ServiceId:   freeTimeParam.ServiceId,
	}

	inlvs, err := c.service.GetFreeTime(request)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request 3")
	}

	return ctx.JSONPretty(http.StatusOK, inlvs, "  ")
}
