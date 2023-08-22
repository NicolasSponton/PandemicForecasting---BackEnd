package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func StartServer() {
	e := echo.New()

	r := e.Group("/app")
	r.GET("/cases", func(c echo.Context) error {
		responseData := CaseData{
			Daily:    dailyCases,
			Forecast: forecast,
		}
		return c.JSON(http.StatusOK, responseData)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
