package main

import (
	ftc "app/internal/controllers/freetime"
	//"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(ctx echo.Context) error {
		controller := ftc.Controller{} 
		return controller.GetFreeTime(ctx)
		//return ctx.HTML(http.StatusOK, "<html><h1>Hello, World!</h1></html>")
	})

	//http.HandleFunc()
	//http.DefaultServeMux()

	e.Logger.Fatal(e.Start(":1324"))
}
