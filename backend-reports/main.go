package main

import (
	"github.com/Blackth01/a3-sistemas-distribuidos/backend-reports/report"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Routes
	e.GET("/report", report.GetReport)

	// Start server
	e.Logger.Fatal(e.Start(":9001"))
}
