package report

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetReport(c echo.Context) error {
	report, err := GetFullReport()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error getting report")
	}

	return c.JSON(http.StatusOK, report)
}
